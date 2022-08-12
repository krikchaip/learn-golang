package reflection_test

import (
	reflection "12-reflection"
	"reflect"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name   string
		X      any
		Result []string
	}{
		{
			Name:   "struct with one string field",
			X:      struct{ Name string }{"Chris"},
			Result: []string{"Chris"},
		},
		{
			Name:   "struct with multiple string fields",
			X:      struct{ Name, City string }{"Chris", "London"},
			Result: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			X: struct {
				Name string
				Age  int
			}{"Chris", 33},
			Result: []string{"Chris"},
		},
		{
			Name: "nested struct",
			X: struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Chris",
				struct {
					Age  int
					City string
				}{33, "London"},
			},
			Result: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			X: &struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Chris",
				struct {
					Age  int
					City string
				}{33, "London"},
			},
			Result: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			X: []struct {
				Age  int
				City string
			}{
				{33, "London"},
				{34, "Reykjavík"},
			},
			Result: []string{"London", "Reykjavík"},
		},
		{
			Name: "arrays",
			X: [2]struct {
				Age  int
				City string
			}{
				{33, "London"},
				{34, "Reykjavík"},
			},
			Result: []string{"London", "Reykjavík"},
		},
		{
			Name: "maps",
			X: map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			// ** [WARNING] this will not always guarantee order
			Result: []string{"Bar", "Boz"},
			// Result: []string{"Boz", "Bar"},
		},
		{
			Name: "channels",
			X: func() chan struct {
				Age  int
				City string
			} {
				ch := make(chan struct {
					Age  int
					City string
				})

				go func() {
					ch <- struct {
						Age  int
						City string
					}{33, "Berlin"}
					ch <- struct {
						Age  int
						City string
					}{34, "Katowice"}

					// ** don't forget to close the channel after finished !
					close(ch)
				}()

				return ch
			}(),
			// ** [WARNING] this will not always guarantee order
			Result: []string{"Berlin", "Katowice"},
			// Result: []string{"Katowice", "Berlin"},
		},
		{
			Name: "with function",
			X: func() (
				*struct {
					Name    string
					Profile struct {
						Age  int
						City string
					}
				},
				[]string,
			) {
				return &struct {
						Name    string
						Profile struct {
							Age  int
							City string
						}
					}{
						"Winner",
						struct {
							Age  int
							City string
						}{
							26, "Bangkok",
						},
					},
					[]string{"Mom", "Dad"}
			},
			Result: []string{"Winner", "Bangkok", "Mom", "Dad"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var comp func(a, b any) (equal bool)

			var got []string
			want := test.Result

			reflection.Walk(test.X, func(v string) {
				got = append(got, v)
			})

			switch {
			case
				strings.Contains(test.Name, "map"),
				strings.Contains(test.Name, "chan"):
				comp = unorderedEqual
			default:
				comp = reflect.DeepEqual
			}

			if equal := comp(got, want); !equal {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

// Supported types: []string
//
// Complexity: O(n)
func unorderedEqual(a, b any) bool {
	as, aOk := a.([]string)
	bs, bOk := b.([]string)

	if !aOk || !bOk || len(as) != len(bs) {
		return false
	}

	set := make(map[string]struct{})
	for _, s := range as {
		set[s] = struct{}{}
	}

	for _, s := range bs {
		if _, ok := set[s]; !ok {
			return false
		}
	}

	return true
}
