package util

import (
	"log"
)

func Catch() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}
