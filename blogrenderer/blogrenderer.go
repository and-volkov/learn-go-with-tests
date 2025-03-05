package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: p}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(post, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type PostViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(post Post, r *PostRenderer) PostViewModel {
	vm := PostViewModel{Post: post}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(post.Body), r.mdParser, nil))
	return vm
}
