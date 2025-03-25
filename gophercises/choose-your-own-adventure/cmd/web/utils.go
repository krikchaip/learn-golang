package main

import (
	"encoding/json"
	"krikchaip/choose-your-own-adventure/internal/model"
	"log"
	"os"
)

func loadJSON(path string) (stories model.Stories) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &stories); err != nil {
		log.Fatal(err)
	}

	return stories
}

func toJSON(v any) string {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println(err)
	}

	return string(bs)
}
