package osexec

import (
	"bytes"
	"encoding/xml"
	"io"
	"os/exec"
	"strings"
)

type Payload struct {
	Message string `xml:"message"`
}

// ?? We are trying to Separate side-effects from business logic
// ?? in order to make it testable.
func GetData(out io.Reader) string {
	var payload Payload
	xml.NewDecoder(out).Decode(&payload)
	return strings.ToUpper(payload.Message)
}

// ?? For this case, getting XML from os.exec is the side-effect.
func getXMLFromCommand() io.Reader {
	cmd := exec.Command("cat", "msg.xml")
	out, _ := cmd.StdoutPipe()

	cmd.Start()
	data, _ := io.ReadAll(out)
	cmd.Wait()

	return bytes.NewReader(data)
}
