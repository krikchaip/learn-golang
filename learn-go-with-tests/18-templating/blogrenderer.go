package blogrenderer

import (
	blogposts "17-reading-files/lib"
	"fmt"
	"io"
	"strings"
)

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
