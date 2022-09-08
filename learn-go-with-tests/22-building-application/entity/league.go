package entity

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(source io.Reader) (league League, err error) {
	// ** reads and parse from source directly
	err = json.NewDecoder(source).Decode(&league)

	// // ?? alternative to json.Decoder (using json.Unmarshal)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(source)
	// err = json.Unmarshal(buf.Bytes(), &league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return
}
