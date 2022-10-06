package osexec

import (
	"encoding/xml"
	"os/exec"
	"strings"
)

type Payload struct {
	Message string `xml:"message"`
}

func GetData() string {
	cmd := exec.Command("cat", "msg.xml")
	out, _ := cmd.StdoutPipe()

	var payload Payload
	decoder := xml.NewDecoder(out)

	cmd.Start()
	decoder.Decode(&payload)
	cmd.Wait()

	return strings.ToUpper(payload.Message)
}
