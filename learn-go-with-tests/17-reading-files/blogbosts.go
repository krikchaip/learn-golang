package blogposts

import "io/fs"

type Post struct {
}

func NewPostFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	posts := make([]Post, len(dir))

	return posts, nil
}
