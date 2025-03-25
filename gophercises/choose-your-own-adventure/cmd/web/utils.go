package main

import (
	"encoding/json"
	"krikchaip/choose-your-own-adventure/internal/model"
	"log"
	"os"
)

func load(path string) (stories model.Stories) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &stories)

	return stories
}
