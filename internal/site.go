package internal

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

//TODO: test performance of caching verus non-caching

type Site struct {
	renderer     *Renderer
	handleFuncs  map[string]func(w http.ResponseWriter, r *http.Request)
	projectCache map[string][]byte
	index        *template.Template
	css          []byte
	projects     []byte
	about        []byte
	contact      []byte
	landing      []byte
	watch        bool
}

func NewSite(watch bool) *Site {
	return &Site{
		renderer:     NewRenderer("dracula"),
		watch:        watch,
		projectCache: make(map[string][]byte),
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

func (s *Site) cachePages() {
	//css
	var err error
	s.css, err = os.ReadFile("static/style.css")
	if err != nil {
		log.Fatalf("unable to load css file due to error: %s", err)
	}
	buf := bytes.NewBuffer(nil)
	s.renderer.htmlFormatter.WriteCSS(buf, s.renderer.highlightStyle)
	s.css = append(s.css, buf.Bytes()...)

	// template
	s.index, err = template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal("could not parse index template: ", err)
	}

	// projects and respective pages
	var proj []Project
	proj, s.projects = s.renderProjects()
	for _, p := range proj {
		s.projectCache[p.Path] = p.Source
	}

	// about
	s.about, err = os.ReadFile("static/about.html")
	if err != nil {
		log.Fatal("could not load about.html: ", err)
	}

	// landing
	s.landing, err = os.ReadFile("static/landing.html")
	if err != nil {
		log.Fatal("could not load landing.html: ", err)
	}

	// contact
	s.contact, err = os.ReadFile("static/contact.html")
	if err != nil {
		log.Fatal("could not load contact.html: ", err)
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

func (s *Site) watchStatic() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	defer watcher.Close()

	go func() {
		log.Println("starting server")
		http.ListenAndServe(":8080", nil)
		log.Println("shutting down server")
	}()

	if err := watcher.Add("static"); err != nil {
		return fmt.Errorf("unable to add static to watcher: %s", err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return errors.New("watcher closed")
			}
			if event.Has(fsnotify.Write) {
				log.Println("recaching pages")
				s.cachePages()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return errors.New("watcher closed")
			}
			log.Println("error:", err)
		}
	}
}

func (s *Site) SetupAndServe() {
	s.setupHandlers()
	s.cachePages()

	if s.watch {
		log.Fatal(s.watchStatic())
	} else {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

func (s *Site) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//TODO: see if we can cache other data as well
		http.ServeFile(w, r, filepath.Join("static", r.URL.String()))
	} else if r.Header.Get("HX-Request") == "true" {
		w.Write(s.landing)
	} else {
		s.handleFill(w, r, s.landing)
	}
}

func (s *Site) handleCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", mime.TypeByExtension(".css"))
	w.Write(s.css)
}

func (s *Site) handleFill(w http.ResponseWriter, r *http.Request, page []byte) {
	if r.Header.Get("HX-Request") == "true" {
		w.Write(page)
	} else {
		err := s.index.Execute(w, template.HTML(page))
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			log.Print("unable to render template when handling fill: ", err)
		}
	}
}

func (s *Site) handleProjects(w http.ResponseWriter, r *http.Request) {
	// load the project html. If cached, return that
	s.handleFill(w, r, s.projects)
}
func (s *Site) handleAbout(w http.ResponseWriter, r *http.Request) {
	s.handleFill(w, r, s.about)
}

func (s *Site) handleContact(w http.ResponseWriter, r *http.Request) {
	s.handleFill(w, r, s.contact)
}

func (s *Site) handleProjectWriteup(w http.ResponseWriter, r *http.Request) {
	projectName := r.PathValue("p")
	if source, found := s.projectCache[projectName]; found {
		s.handleFill(w, r, source)
		return
	} else {
		http.Error(w, "resource not found", http.StatusNotFound)
		log.Printf("requested project %s which was not found in the page cache", projectName)
	}
}

func (s *Site) handleContactSubmit(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
