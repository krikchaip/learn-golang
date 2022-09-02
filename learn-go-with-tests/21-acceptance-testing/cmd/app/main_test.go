package main_test

import (
	"testing"
	"time"

	"21-acceptance-testing/lib/testutil"
	"21-acceptance-testing/lib/testutil/assert"
)

const (
	PORT string = "8080"
	URL  string = "http://localhost:" + PORT
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, interupt := testutil.LaunchTestProgram(PORT)
	t.Cleanup(cleanup) // ?? afterThis()

	// just check the server works before we shut things down
	assert.CanGet(t, URL)

	// fire off a request, and before it has a chance to respond send SIGTERM.
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoPanic(t, interupt)
	})

	// Without graceful shutdown, this would fail
	assert.CanGet(t, URL)

	// // after interrupt, the server should be shutdown, and no more requests will work
	// assert.CantGet(t, URL)
}
