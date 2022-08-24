package blogrenderer

import (
	blogposts "17-reading-files/lib"
	"fmt"
	"html/template"
	"io"
	"strings"
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func RenderHTML(w io.Writer, post blogposts.Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}

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