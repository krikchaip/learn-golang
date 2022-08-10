package dependency_injection

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ?? you are allow to pass a pointer of a type
// ?? that implements the same interface as an interface parameter
// ?? eg. *bytes.Buffer, *os.File, ...
func Greet(writer io.Writer, name string) {
	// ** this is actually what fmt.Printf does under the hood!
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetStdout() {
	Greet(os.Stdout, "Elodie\n")
}

func GreetHttp() {
	// ** convert a normal function to HTTPHandler type
	// ** this is basically like int(2.3), float(3), ...
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "world!") // ** w implementa Writer interface
	})

	// go to http://localhost:5001
	http.ListenAndServe(":5001", handler)
}
