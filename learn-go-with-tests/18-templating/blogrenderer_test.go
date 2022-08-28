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

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post to HTML", func(t *testing.T) {
		w := &bytes.Buffer{}

		if err := pr.RenderHTML(w, post); err != nil {
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

	t.Run("it renders an index of posts", func(t *testing.T) {
		w := &bytes.Buffer{}
		posts := []blogposts.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}

		if err := pr.RenderIndexHTML(w, posts); err != nil {
			t.Fatal(err)
		}

		// ?? will replace this
		// got := w.String()
		// got := w.String()
		// want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`

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

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pr.RenderHTML(w, post)
	}
}
