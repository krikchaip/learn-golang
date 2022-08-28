package blogrenderer

import (
	blogposts "17-reading-files/lib"
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"
)

// ?? replace this string literal with actual template files
// const (
// 	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
// )

//go:embed templates/*.gohtml
var postTemplates embed.FS

type PostRenderer struct {
	templ *template.Template
}

// ?? to parse templates only once
func NewPostRenderer() (*PostRenderer, error) {
	// ?? parsing tempalte from a string literal
	// templ, err := template.New("blog").Parse(postTemplate)

	templ, err := template.New("").
		// ?? providing template functions
		// ** NOT RECOMMENED - no separation of concern (hard to test)
		Funcs(template.FuncMap{
			"snakecase": func(s string) string {
				return strings.ToLower(strings.Replace(s, " ", "-", -1))
			},
		}).

		// ?? use io.FS to parse multiple files 👍🏻
		ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (pr *PostRenderer) RenderHTML(w io.Writer, post blogposts.Post) error {
	// ** be careful with template names when calling templ.Execute with ParseFS.
	// ** only the first file will be the render output.
	// ** eg. layout.gohtml, blog.gohtml -> will render blog.gohtml
	// **     _.gohtml, blog.gohtml      -> will render _.gohtml
	// if err := templ.Execute(w, post); err != nil {
	// 	return err
	// }

	// ?? the safer alternative
	if err := pr.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

func (pr *PostRenderer) RenderIndexHTML(w io.Writer, posts []blogposts.Post) error {
	if err := pr.templ.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}

	return nil
}

// ?? basic version
func Render(w io.Writer, post blogposts.Post) error {
	_, err := fmt.Fprint(w,
		tag("h1", post.Title),
		tag("p", post.Description),
		fmt.Sprintf(
			"Tags: %s",
			tag("ul",
				strings.Join(transform(post.Tags, func(t string) string {
					return tag("li", t)
				}), ""),
			),
		),
	)

	return err
}

func tag(name, content string) string {
	return fmt.Sprintf("<%[1]s>%[2]s</%[1]s>", name, content)
}

func transform[T, U any](xs []T, f func(x T) U) []U {
	results := make([]U, 0, len(xs))

	for _, x := range xs {
		results = append(results, f(x))
	}

	return results
}
