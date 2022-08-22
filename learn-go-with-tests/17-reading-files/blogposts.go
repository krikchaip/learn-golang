package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		p, err := getPost(filesystem, f.Name())

		// TODO: needs clarification, should we totally fail
		// if one file fails? or just ignore?
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error) {
	file, err := filesystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	return parsePost(file)
}

func parsePost(f io.Reader) (Post, error) {
	content, err := io.ReadAll(f)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(content)[7:]}
	return post, nil
}
