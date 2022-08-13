package lib

func Public(path string) string {
	if rune(path[0]) != '/' {
		path = "/" + path
	}

	return "./public" + path
}
