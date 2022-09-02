package util

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Catch() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}

func FindRoot(start string) (packageRoot string, err error) {
	var entries []fs.DirEntry

	if !filepath.IsAbs(start) {
		start, _ = filepath.Abs(start)
	}

	if entries, err = os.ReadDir(start); err != nil {
		return "", err
	}

	for _, ent := range entries {
		if ent.Name() == "go.mod" && ent.Type().IsRegular() {
			return start, nil
		}
	}

	if start == filepath.FromSlash("/") {
		return "", fmt.Errorf("go.mod was not found along the path\n")
	}

	return FindRoot(filepath.Dir(start))
}
