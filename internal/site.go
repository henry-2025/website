package internal

import (
	"bytes"
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type cacheIndex uint8

const (
	CSS cacheIndex = iota
	PROJ
	ABOUT
	CONTACT
	TEMPLATE
	LANDING
	INDEX
)

//TODO: must implement full-page and partial page rendering, which is done by checking the HX-Request header being set
//TODO: test performance of caching verus non-caching

type Site struct {
	renderer     *Renderer
	handleFuncs  map[string]func(w http.ResponseWriter, r *http.Request)
	doCache      bool
	pageCache    map[cacheIndex][]byte
	projectCache map[string][]byte
}

func NewSite(cache bool) *Site {
	return &Site{
		renderer:     NewRenderer("dracula"),
		handleFuncs:  nil,
		pageCache:    make(map[cacheIndex][]byte),
		projectCache: make(map[string][]byte),
		doCache:      cache,
	}
}

func (s *Site) setupHandlers() {
	s.handleFuncs = map[string]func(w http.ResponseWriter, r *http.Request){
		"/":                    s.handleRoot,
		"/style.css":           s.handleCss,
		"/projects":            s.handleProjects,
		"/projects/{p}":        s.handleProjectWriteup,
		"/about":               s.handleAbout,
		"/contact":             s.handleContact,
		"POST /contact/submit": s.handleContactSubmit,
	}
	for handler, function := range s.handleFuncs {
		http.HandleFunc(handler, function)
	}
}

func (s *Site) renderProjects() ([]Project, []byte) {
	entries, err := os.ReadDir("static/projects")
	if err != nil {
		log.Fatal("problem reading entries: ", err)
	}

	var projects []Project

	for _, e := range entries {
		if !e.IsDir() {
			b, err := os.ReadFile(filepath.Join("static/projects", e.Name()))
			if err != nil {
				log.Fatalf("problem reading file %s: %s", e.Name(), err)
			}

			proj := s.renderer.RenderProject(b)
			proj.Path = strings.Split(filepath.Base(e.Name()), ".")[0]
			projects = append(projects, proj)
		}
	}

	tmpl, err := template.ParseFiles("static/projects.html")
	if err != nil {
		log.Fatal("could not parse project template: ", err)
	}

	payload := bytes.NewBuffer(nil)
	err = tmpl.Execute(payload, projects)
	if err != nil {
		log.Fatal("executing projects template: ", err)
	}

	return projects, payload.Bytes()
}

func (s *Site) SetupAndServe() {
	s.setupHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Site) handleRoot(w http.ResponseWriter, r *http.Request) {
	if !s.doCache {
		if r.URL.Path != "/" {
			//TODO: see if we can cache other data as well
			http.ServeFile(w, r, filepath.Join("static", r.URL.String()))
		} else if r.Header.Get("HX-Request") == "true" {
			http.ServeFile(w, r, "static/landing.html")
		} else {
			tmpl, err := os.ReadFile("static/index.html")
			if err != nil {
				log.Fatal("could not load index: ", err)
			}
			landing, err := os.ReadFile("static/landing.html")
			if err != nil {
				log.Fatal("could not load landing: ", err)
			}
			t, err := template.New("index").Parse(string(tmpl))
			if err != nil {
				log.Fatal("could not parse index template: ", err)
			}
			buf := bytes.NewBuffer(nil)
			t.Execute(buf, template.HTML(landing))
			w.Write(buf.Bytes())
		}
	} else {
		if r.URL.Path != "/" {
			//TODO: see if we can cache other data as well
			http.ServeFile(w, r, filepath.Join("static", r.URL.String()))
		} else if r.Header.Get("HX-Request") == "true" {
			payload, found := s.pageCache[LANDING]
			if !found {
				var err error
				payload, err = os.ReadFile("static/landing.html")
				if err != nil {
					log.Fatal("unable to load landing file: ", err)
				}
			}
			w.Write(payload)
		} else {
			payload, found := s.pageCache[INDEX]
			if !found {
				tmpl, found := s.pageCache[TEMPLATE]
				if !found {
					var err error
					tmpl, err = os.ReadFile("static/index.html")
					if err != nil {
						log.Fatal("unable to load template file: ", err)
					}
					s.pageCache[TEMPLATE] = tmpl
				}
				landing, found := s.pageCache[TEMPLATE]
				if !found {
					var err error
					landing, err = os.ReadFile("static/landing.html")
					if err != nil {
						log.Fatal("unable to load template file: ", err)
					}
					s.pageCache[LANDING] = landing
				}
				t, err := template.New("landing").Parse(string(tmpl))
				if err != nil {
					log.Fatal("error creating template", err)
				}
				p := bytes.NewBuffer(nil)
				t.Execute(p, landing)
				payload = p.Bytes()
			}
			w.Write(payload)
		}
	}
}

func (s *Site) handleCss(w http.ResponseWriter, r *http.Request) {
	css, found := s.pageCache[CSS]
	if !found || !s.doCache {
		var err error
		css, err = os.ReadFile("static/style.css")
		if err != nil {
			log.Printf("unable to load css file due to error: %s", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		buf := bytes.NewBuffer(nil)
		s.renderer.htmlFormatter.WriteCSS(buf, s.renderer.highlightStyle)
		css = append(css, buf.Bytes()...)
		s.pageCache[CSS] = css
	}
	w.Header().Set("Content-Type", mime.TypeByExtension(".css"))
	w.Write(css)
}

func (s *Site) handleProjects(w http.ResponseWriter, r *http.Request) {
	// load the project html. If cached, return that
	if r.Header.Get("HX-Request") == "true" {
		var payload []byte
		if s.doCache {
			var cached bool
			payload, cached = s.pageCache[PROJ]
			if !cached {
				var projectSources []Project
				projectSources, payload = s.renderProjects()
				if s.doCache {
					s.pageCache[PROJ] = payload
					for _, p := range projectSources {
						s.projectCache[p.Path] = p.Source
					}
				}
			}
		} else {
			_, payload = s.renderProjects()
		}
		w.Write(payload)

	} else {
		tmpl, err := os.ReadFile("static/index.html")
		if err != nil {
			log.Fatal("could not load index: ", err)
		}
		_, payload := s.renderProjects()
		t, err := template.New("index").Parse(string(tmpl))
		if err != nil {
			log.Fatal("could not parse index template: ", err)
		}
		buf := bytes.NewBuffer(nil)
		t.Execute(buf, template.HTML(payload))
		w.Write(buf.Bytes())
	}
}
func (s *Site) handleAbout(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		payload, cached := s.pageCache[ABOUT]
		if !cached || s.doCache {
			http.ServeFile(w, r, "static/about.html")
			return
		}
		w.Write(payload)
	} else {
		tmpl, err := os.ReadFile("static/index.html")
		if err != nil {
			log.Fatal("could not load index: ", err)
		}
		about, err := os.ReadFile("static/about.html")
		if err != nil {
			log.Fatal("could not load index: ", err)
		}
		t, err := template.New("index").Parse(string(tmpl))
		if err != nil {
			log.Fatal("could not parse index template: ", err)
		}
		buf := bytes.NewBuffer(nil)
		t.Execute(buf, template.HTML(about))
		w.Write(buf.Bytes())
	}
}

func (s *Site) handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		payload, cached := s.pageCache[CONTACT]
		if !cached || s.doCache {
			http.ServeFile(w, r, "static/contact.html")
			return
		}
		w.Write(payload)
	} else {
		tmpl, err := os.ReadFile("static/index.html")
		if err != nil {
			log.Fatal("could not load index: ", err)
		}
		about, err := os.ReadFile("static/contact.html")
		if err != nil {
			log.Fatal("could not load index: ", err)
		}
		t, err := template.New("index").Parse(string(tmpl))
		if err != nil {
			log.Fatal("could not parse index template: ", err)
		}
		buf := bytes.NewBuffer(nil)
		t.Execute(buf, template.HTML(about))
		w.Write(buf.Bytes())
	}
}

func (s *Site) handleProjectWriteup(w http.ResponseWriter, r *http.Request) {
	projectName := r.PathValue("p")
	if source, found := s.projectCache[projectName]; found {
		w.Write(source)
		return
	}
	var (
		projectSources []Project
		found          bool
		source         []byte
	)

	projectSources, _ = s.renderProjects()
	for _, p := range projectSources {
		if s.doCache {
			s.projectCache[p.Path] = p.Source
		}
		if p.Path == projectName {
			found = true
			source = p.Source
		}
	}

	if found {
		w.Write(source)
	} else {
		http.NotFound(w, r)
	}
}

func (s *Site) handleContactSubmit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
