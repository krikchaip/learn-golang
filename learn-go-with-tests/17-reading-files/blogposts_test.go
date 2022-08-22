package blogposts_test

import (
	blogposts "17-reading-files"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md":  {Data: []byte("Title: Post 1")},
			"hello_world2.md": {Data: []byte("Title: Post 2")},
		}
		posts, err := blogposts.NewPostFromFS(fs)

		got := posts
		want := []blogposts.Post{
			{Title: "Post 1"},
			{Title: "Post 2"},
		}

		if err != nil {
			t.Fatal(err)
		}

		assertPosts(t, got, want)
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

func assertPosts(t testing.TB, got, want []blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
