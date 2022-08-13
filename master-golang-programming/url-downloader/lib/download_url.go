package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadURL(url string) error {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s -> [code %d]\n", url, res.StatusCode)
	}

	fmt.Printf("%s -> code %d\n", url, res.StatusCode)

	content, _ := io.ReadAll(res.Body)
	return os.WriteFile(formatURL(url), content, 0644)
}

func formatURL(url string) string {
	// extract domain name from URL
	domain := strings.Split(url, "//")[1]

	return Public(domain) + ".html"
}
