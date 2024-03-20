package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

type Renderer struct {
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
}

func NewRenderer(codeStyle string) *Renderer {
	formatter := html.New(html.WithClasses(true), html.TabWidth(2))
	if formatter == nil {
		panic("could not create html formatter")
	}

	style := styles.Get(codeStyle)
	if style == nil {
		log.Printf("could not find the style %s, defaulting to dracula", codeStyle)
	}
	style = styles.Get("dracula")

	return &Renderer{
		htmlFormatter:  formatter,
		highlightStyle: style,
	}
}

// TODO: implement this to actually parse metadata in the project. Will probably reimplement the yaml header that most blog sites do
func parseProjectHeader(doc ast.Node) Project {
	headerParsed := false
	var p Project
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if text, ok := node.(*ast.Text); ok && entering && !headerParsed {
			buf := bytes.NewBuffer(text.Leaf.Literal)
			err := json.NewDecoder(buf).Decode(&p)
			if err != nil {
				log.Fatal("decoding project header: ", err)
			}
			headerParsed = true
			ast.RemoveFromTree(node)
		}
		return ast.GoToNext
	})
	return p
}

func (r *Renderer) RenderProject(md []byte) Project {
	parse := parser.NewWithExtensions(parser.CommonExtensions)
	doc := parse.Parse(md)
	proj := parseProjectHeader(doc)
	renderer := newCustomizedRender(r)

	source := markdown.Render(doc, renderer)
	proj.Source = source
	return proj
}

// based on https://github.com/alecthomas/chroma/blob/master/quick/quick.go
func (r *Renderer) htmlHighlight(w io.Writer, source, lexer, style string) error {
	// Determine lexer.
	l := lexers.Get(lexer)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return r.htmlFormatter.Format(w, r.highlightStyle, it)
}

// an actual rendering of Paragraph is more complicated
func (r *Renderer) renderCode(w io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	defaultLang := ""
	lang := string(codeBlock.Info)
	r.htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
}

// where we apply any of the logic that is not processed in default rendering
func (r *Renderer) myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*ast.Document); ok && entering {
		w.Write([]byte("<style>"))
		r.htmlFormatter.WriteCSS(w, r.highlightStyle)
		w.Write([]byte("</style>"))
		return ast.GoToNext, false
	}

	if code, ok := node.(*ast.CodeBlock); ok {
		r.renderCode(w, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func newCustomizedRender(r *Renderer) *mdhtml.Renderer {
	opts := mdhtml.RendererOptions{
		Flags:          mdhtml.CommonFlags,
		RenderNodeHook: r.myRenderHook,
	}
	return mdhtml.NewRenderer(opts)
}
