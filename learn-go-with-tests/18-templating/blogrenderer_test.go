package blogrenderer_test

import (
	blogposts "17-reading-files/lib"
	blogrenderer "18-templating"
	"bytes"
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func init() {
	approvals.UseFolder("snapshots")
}

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

		if err := blogrenderer.RenderHTML(w, post); err != nil {
			t.Fatal(err)
		}

		// ?? will replace this
		// got := w.String()
		// want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		// ?? with snapshot testing
		approvals.VerifyString(t, w.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		w    = &strings.Builder{}
		post = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogrenderer.RenderHTML(w, post)
	}
}
