package main

import (
	blogposts "17-reading-files/lib"
	"embed"
	"fmt"
	"os"
)

//go:embed blogs/*.md
var folder embed.FS

// go run 17-reading-files
func main() {
	// ?? using os.DirFS
	posts, _ := blogposts.NewPostFromFS(os.DirFS("17-reading-files/blogs"))
	fmt.Printf("%+v\n", posts)

	// ?? using go:embed
	dir, _ := folder.ReadDir(".")
	for _, f := range dir {
		// blogs (will show only folder)
		fmt.Printf("name: %q, folder: %v\n", f.Name(), f.IsDir())
	}

	// posts, _ = blogposts.NewPostFromFS(folder)
	// fmt.Printf("%+v", posts)
}
