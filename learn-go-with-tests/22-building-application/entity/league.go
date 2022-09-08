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

// Returns reference to a player in the league
//
// NOTE: if we choose not to return a pointer to Player then we can't mutate.
// Because of "pass-by-value" mechanism in Go.
//
// Example:
//
//	for i, p := range league {
//		if name == p.Name {
//			// ** this will not work because when you `range` over a slice
//			// ** you are returned a COPY OF AN ELEMENT at the current index.
//			// p.Wins++
//
//			// ** For that reason, we need to get the reference of the actual value
//			// ** and then changing that value instead.
//			league[i].Wins++
//		}
//	}
func (l League) Find(name string) *Player {
	for i, p := range l {
		if name == p.Name {
			return &l[i]
		}
	}

	return nil
}
