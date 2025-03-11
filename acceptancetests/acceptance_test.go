package main

import (
	"testing"
	"time"

	"github.com/quii/go-graceful-shutdown/acceptancetests"
	"github.com/quii/go-graceful-shutdown/assert"
)

const (
	url = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	assert.CanGet(t, url)

	time.AfterFunc(time.Millisecond*50, func() {
		assert.NoError(t, sendInterrupt())
	})

	assert.CanGet(t, url)

	assert.CantGet(t, url)
}
