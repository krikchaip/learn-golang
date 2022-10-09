package ctx_reader_test

import (
	"context"
	"io"
	"strings"
	"testing"

	ctx_reader "25-context-reader"
)

// ?? An easy way to start this kind of work is to wrap your delegate
// ?? and write a test that asserts it behaves how the delegate normally does
// ?? before you start composing other parts to change behaviour.
// ?? This will help you to keep things working correctly as you code toward your goal

func TestContextAwareReader(t *testing.T) {
	t.Run("read with buffer length of 3", func(t *testing.T) {
		ctx := context.Background()
		reader := ctx_reader.NewContextReader(ctx, strings.NewReader("123456"))

		got := make([]byte, 3)

		readOnce(t, reader, got)
		assertBufferHas(t, got, "123")

		readOnce(t, reader, got)
		assertBufferHas(t, got, "456")
	})

	t.Run("stops reading when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		reader := ctx_reader.NewContextReader(ctx, strings.NewReader("123456"))

		got := make([]byte, 3)

		readOnce(t, reader, got)
		assertBufferHas(t, got, "123")

		cancel()
		assertCannotRead(t, reader, got)
	})
}

func readOnce(t *testing.T, reader io.Reader, res []byte) int {
	t.Helper()

	n, err := reader.Read(res)
	if err != nil {
		t.Fatal(err)
	}

	return n
}

func assertCannotRead(t *testing.T, reader io.Reader, res []byte) {
	t.Helper()

	n, err := reader.Read(res)
	if err == nil {
		t.Error("expected an error after cancellation but didn't get one")
	}

	if n > 0 {
		t.Errorf("expected 0 bytes to be read after cancellation but %d were read", n)
	}
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
