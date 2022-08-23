package blogrenderer_test

import (
	blogposts "17-reading-files/lib"
	blogrenderer "18-templating"
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		post = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post to HTML", func(t *testing.T) {
		w := &bytes.Buffer{}
		err := blogrenderer.Render(w, post)

		if err != nil {
			t.Fatal(err)
		}

		got := w.String()
		want := `<h1>hello world</h1>`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
