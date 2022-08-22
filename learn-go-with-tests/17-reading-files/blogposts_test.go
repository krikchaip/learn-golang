package blogposts_test

import (
	blogposts "17-reading-files"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md":  {Data: []byte("hi")},
			"hello_world2.md": {Data: []byte("hola")},
		}
		posts, err := blogposts.NewPostFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("on error", func(t *testing.T) {
		fs := StubFailingFS{}
		_, err := blogposts.NewPostFromFS(fs)

		if err == nil {
			t.Error("expect an error but didn't get one")
		}
	})
}

// implements: fs.Fs
type StubFailingFS struct{}

// ?? fs.ReadDir will call this method inside
func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
