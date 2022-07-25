package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEqual := func(t testing.TB, expected, actual string) {
		t.Helper()
		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	}

	t.Run("no repeat", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertEqual(t, expected, repeated)
	})

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertEqual(t, expected, repeated)
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("Winner", 2))
	fmt.Printf("%q", Repeat("Bullshit", 0))
	// Output:
	// WinnerWinner
	// ""
}

// ?? run `go test -bench=.` to see the result
func BenchmarkRepeat(b *testing.B) {
	// the framework will determine what is a good value for "b.N"
	// so you don't have to worry.
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
