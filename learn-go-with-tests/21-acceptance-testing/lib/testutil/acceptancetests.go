package testutil

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"21-acceptance-testing/lib/util"
)

const (
	APP_PATH      = "cmd/app"
	BIN_PATH      = "bin/actest"
	HANDLER_DELAY = 2 * time.Second
	PING_RETRIES  = 30
	PING_DELAY    = 100 * time.Millisecond
)

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(HANDLER_DELAY)
	w.Write([]byte("Hello World!"))
}

func LaunchTestProgram(port string) (
	cleanup func(),
	interupt func(),
) {
	defer util.Catch()

	output := build()
	cleanup, interupt = runServer(port, output)

	return
}

func build() (output string) {
	root, _ := util.FindRoot(".")
	pkg := filepath.Join(root, APP_PATH)
	output = filepath.Join(root, BIN_PATH)

	build := exec.Command("go", "build", "-o", output, pkg)
	if err := build.Run(); err != nil {
		panic(fmt.Errorf("cannot build tool %s: %s", output, err))
	}

	return output
}

func runServer(port, bin string) (cleanup, interupt func()) {
	cmd := exec.Command(bin)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		panic(fmt.Errorf("cannot run %s: %s", bin, err))
	}

	cleanup = func() {
		if err := cmd.Process.Kill(); err != nil {
			panic(err)
		}
		os.Remove(bin)
	}

	interupt = func() {
		if err := cmd.Process.Signal(os.Interrupt); err != nil {
			panic(err)
		}
	}

	if err := ping(port); err != nil {
		cleanup()
		panic(err)
	}

	return
}

func ping(port string) error {
	for i := 0; i < PING_RETRIES; i++ {
		// connecting to localhost:port
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", port))

		// close on successful
		if conn != nil {
			return conn.Close()
		}

		time.Sleep(PING_DELAY)
	}

	return fmt.Errorf("nothing seems to be listening on localhost:%s", port)
}
