package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int

	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process id")
	}

	// simulate killing process id
	fmt.Printf("killing process id %v\n", pid)

	if err := os.Remove(pidFile); err != nil {
		// we can go on if we fail here
		log.Printf("warning, can't remove pid file: %s", err)
	}

	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error %s\n", err)
		os.Exit(1)
	}
}
