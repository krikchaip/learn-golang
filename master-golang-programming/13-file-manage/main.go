package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ?? CREATE a new file
	filePtr, err := os.Create(public("test.txt"))
	// filePtr, err := os.Create(public("/err/should/not/nil"))
	handleErr(err)

	// ** don't forget to close the file after finished working
	filePtr.Close()

	// ?? REDUCE file content to N bytes
	err = os.Truncate(public("test.txt"), 0)
	handleErr(err)

	// ?? OPEN a file to do something (read/write/etc.)
	filePtr, err = os.OpenFile(public("test.txt"), os.O_APPEND|os.O_CREATE, 0644)
	handleErr(err)
	filePtr.Close()

	// ?? get a FILE STAT (file info)
	fileInfo, err := os.Stat(public("test.txt"))
	handleErr(err)
	fmt.Printf(
		"File Name: %q\n"+
			"Size in Bytes: %d\n"+
			"Last Modified: %v\n"+
			"Is Dir? %t\n"+
			"Permissions: %v\n"+
			"\n",
		fileInfo.Name(),
		fileInfo.Size(),
		fileInfo.ModTime(),
		fileInfo.IsDir(),
		fileInfo.Mode(),
	)

	// ?? validating an error value returned by os package
	if _, err := os.Stat("foo"); os.IsNotExist(err) {
		log.Print("File does not exist!")
	}

	// ?? RENAMING a file
	err = os.Rename(public("test.txt"), public("eiei.txt"))
	handleErr(err)

	// ?? REMOVE a file
	err = os.Remove(public("eiei.txt"))
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
