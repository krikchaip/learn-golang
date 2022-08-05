package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open(public(".gitkeep"))
	handleErr(err)
	defer file.Close()

	// ?? READ EXACTLY N bytes from a file
	bytes := make([]byte, 2)
	nRead, err := io.ReadFull(file, bytes)
	handleErr(err)

	log.Printf("Number of bytes read: %d\n", nRead) // 2
	log.Printf("Data read: %s\n", bytes)            // EI

	// ?? READ whole file into memory
	content, _ := os.ReadFile(public(".gitkeep"))
	log.Printf("%s\n", content) // EIEI ... YOU ARE PRANKED!

	// ?? READ til end, start from current file pointer
	// ** continue reading from byte 3 and so on (because we read 2bytes before)
	content, _ = io.ReadAll(file)
	log.Printf("%s\n", content) // EI ... YOU ARE PRANKED!
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
