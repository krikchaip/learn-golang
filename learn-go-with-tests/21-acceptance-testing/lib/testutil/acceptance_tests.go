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
	modulePath  = "21-acceptance-testing/cmd/app"
	baseBinName = "21-acceptance-testing/bin/actest"
)

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.Write([]byte("Hello World!"))
}

func LaunchTestProgram(port string) (
	cleanup func(),
	interupt func(),
) {
	defer util.Catch()

	binName := buildBinary()
	// fmt.Printf("binName: %v\n", binName)
	cleanup, interupt = runServer(binName, port)

	return
}

func buildBinary() (name string) {
	name = absolutePath(baseBinName)
	path := absolutePath(modulePath)

	build := exec.Command("go", "build", "-o", name, path)
	if err := build.Run(); err != nil {
		panic(fmt.Errorf("cannot build tool %s: %s", name, err))
	}

	return
}

func runServer(name, port string) (cleanup, interupt func()) {
	binCmd := exec.Command(name)

	binCmd.Stdout = os.Stdout
	binCmd.Stderr = os.Stderr

	if err := binCmd.Start(); err != nil {
		panic(fmt.Errorf("cannot run %s: %s", name, err))
	}

	cleanup = func() {
		if err := binCmd.Process.Kill(); err != nil {
			panic(err)
		}
		os.Remove(name)
	}

	interupt = func() {
		if err := binCmd.Process.Signal(os.Interrupt); err != nil {
			panic(err)
		}
	}

	if err := waitForNetworkResponse(port); err != nil {
		cleanup()
		panic(err)
	}

	return
}

func absolutePath(name string) string {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	path := filepath.Join(dir, name)
	return path
}

func waitForNetworkResponse(port string) error {
	const (
		RETRIES = 30
		DELAY   = 100 * time.Millisecond
	)

	for i := 0; i < RETRIES; i++ {
		// connecting to localhost:port
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", port))

		// close on successful
		if conn != nil {
			return conn.Close()
		}

		time.Sleep(DELAY)
	}

	return fmt.Errorf("nothing seems to be listening on localhost:%s", port)
}
