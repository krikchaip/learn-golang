package generics_test

import (
	generics "19-generics"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})

	// AssertEqual(t, 1, "1") // uncomment to see the error
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		st := new(generics.Stack[int])

		// check stack is empty
		AssertTrue(t, st.IsEmpty())

		// add a thing, then check it's not empty
		st.Push(123)
		AssertFalse(t, st.IsEmpty())

		// add another thing, pop it back again
		st.Push(456)
		value, _ := st.Pop()
		AssertEqual(t, value, 456)

		value, _ = st.Pop()
		AssertEqual(t, value, 123)

		value, ok := st.Pop()
		AssertEqual(t, struct {
			int
			bool
		}{value, ok}, struct {
			int
			bool
		}{0, false})

		AssertTrue(t, st.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		st.Push(1)
		st.Push(2)
		firstNum, _ := st.Pop()
		secondNum, _ := st.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})

	t.Run("string stack", func(t *testing.T) {
		st := new(generics.Stack[string])

		// check stack is empty
		AssertTrue(t, st.IsEmpty())

		// add a thing, then check it's not empty
		st.Push("123")
		AssertFalse(t, st.IsEmpty())

		// add another thing, pop it back again
		st.Push("456")
		value, _ := st.Pop()
		AssertEqual(t, value, "456")

		value, _ = st.Pop()
		AssertEqual(t, value, "123")

		AssertTrue(t, st.IsEmpty())
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
