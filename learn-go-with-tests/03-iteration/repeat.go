package iteration

import "strings"

func Repeat(char string, times int) (repeated string) {
	// ?? old version
	// for i := 0; i < times; i++ {
	// 	repeated += char
	// }

	// ?? refactored
	repeated = strings.Repeat(char, times)
	return
}
