package countdown

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "3")
}
