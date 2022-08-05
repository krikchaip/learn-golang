package main

import (
	"log"
	"os"
)

func main() {
	filename := public("writing.txt")
	bytes := []byte("I lean Golang! วินเนอร์") // ** converting a string to a bytes slice

	// ?? WRITE bytes to file - the "low level" way

	// ** opening the file in write-only mode if the file exists
	// ** and then it truncates the file.
	// ** if the file doesn't exist it creates the file with 0644 permissions
	file, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0644,
	)

	handleErr(err)
	defer file.Close() // ** will call this statement to the end of the function

	nWritten, err := file.Write(bytes)
	handleErr(err)
	log.Printf("Bytes written: %d\n", nWritten) // 2022/08/06 01:21:19 Bytes written: 39

	// ?? WRITE bytes to file - the convenient way
	err = os.WriteFile(filename, bytes, 0644)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Printf("err type: %T\n", err)

		// ** will add date and time prefix to every log message
		// ** eg. YYYY/MM/DD HH:MM:SS {message}
		log.Fatal(err)
	}
}

func public(path string) string {
	if rune(path[0]) != '/' {
		path = "/" + path
	}

	return "./public" + path
}
