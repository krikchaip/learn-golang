package main_test

import (
	"testing"

	"21-acceptance-testing/lib/testutil"
	"21-acceptance-testing/lib/testutil/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, _ := testutil.LaunchTestProgram(port)
	t.Cleanup(cleanup) // ?? afterThis()

	// just check the server works before we shut things down
	assert.CanGet(t, url)
}
