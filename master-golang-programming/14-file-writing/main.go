package main

import (
	"bufio"
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
	log.Printf("Bytes written: %d\n", nWritten) // 39

	// ?? WRITE bytes to file - the convenient way
	err = os.WriteFile(filename, bytes, 0644)
	handleErr(err)

	// ?? WRITE bytes to file - using buffer
	buffer := bufio.NewWriter(file) // ** first, create a buffer writer to file

	// ** writing the bytes slice to a buffer IN MEMORY
	nWritten, err = buffer.Write([]byte{'a', 'b', 'c'})
	handleErr(err)
	log.Printf("Bytes written to buffer (not file): %d\n", nWritten) // 3

	// ?? checking available buffer (default max = 4096bytes)
	log.Printf("Bytes available in buffer: %d\n", buffer.Available()) // 4093

	// ** writing a string to a buffer IN MEMORY
	nWritten, err = buffer.WriteString("\nJust a random string")
	handleErr(err)
	log.Printf("Bytes written to buffer (not file): %d\n", nWritten) // 21

	// ?? checking how much data is stored in buffer,
	// ?? just waiting to be written to disk.
	log.Printf("Bytes buffered: %d\n", buffer.Buffered()) // 24

	// ** FLUSH bytes written in buffer to the file
	buffer.Flush()
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
