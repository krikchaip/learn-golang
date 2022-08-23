package blogrenderer

import (
	blogposts "17-reading-files/lib"
	"fmt"
	"io"
)

func Render(w io.Writer, post blogposts.Post) error {
	_, err := fmt.Fprint(w, tag("h1", post.Title))
	return err
}

func tag(name, content string) string {
	return fmt.Sprintf("<%[1]s>%[2]s</%[1]s>", name, content)
}
