package lib

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	modulePath  = "21-acceptance-testing/cmd/app"
	baseBinName = "21-acceptance-testing/bin/actest"
)

func LaunchTestProgram(port string) (
	cleanup func(),
	interupt func(),
) {
	defer catch()

	binName := buildBinary()
	cleanup, interupt = runServer(binName, port)

	return
}

func buildBinary() (name string) {
	name = baseBinName

	build := exec.Command("go", "build", "-o", name, modulePath)
	if err := build.Run(); err != nil {
		panic(fmt.Errorf("cannot build tool %s: %s", name, err))
	}

	return
}

func runServer(name, port string) (cleanup, interupt func()) {
	binCmd := absoluteCmd(name)

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

func absoluteCmd(name string) *exec.Cmd {
	dir, _ := os.Getwd()
	cmdpath := filepath.Join(dir, name)
	return exec.Command(cmdpath)
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
