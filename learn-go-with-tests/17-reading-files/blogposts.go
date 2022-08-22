package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
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

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func parsePost(f io.Reader) (Post, error) {
	// ?? scan line by line
	scanner := bufio.NewScanner(f)

	// ?? TrimPrefix(prefix) == s[len(prefix):]
	readMetaLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	Title := readMetaLine(titleSeparator)
	Description := readMetaLine(descriptionSeparator)

	post := Post{Title, Description}
	return post, nil
}
