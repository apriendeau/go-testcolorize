package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"

	"github.com/apriendeau/go-testcolorize"
)

var mu sync.Mutex

func Test(args []string) {
	log.SetFlags(0)
	if !Contains("-v", args) {
		tmp := []string{args[0], "-v"}
		args = append(tmp, args[1:]...)
	}

	// Execute "go" command with the same arguments.
	cmd := exec.Command("go", args...)

	// Pass through standard input.
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go io.Copy(stdin, os.Stdin)

	// Create a wait group for stdout/stderr.
	var wg sync.WaitGroup
	wg.Add(2)

	// Pass through standard out.
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		processPipe(os.Stdout, stdout)
		wg.Done()
	}()

	// Read through stderr and decorate.
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		processPipe(os.Stderr, stderr)
		wg.Done()
	}()

	// Execute command.
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Forward signals to command.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c)
		for sig := range c {
			cmd.Process.Signal(sig)
		}
	}()

	// Wait for pipes to finish reading and then wait for command to exit.
	wg.Wait()

	if err := cmd.Wait(); err != nil {
		os.Exit(1)
	}
}

// processPipe scans the src by line and attempts to match the first FILE:LINE.
func processPipe(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		func() {
			mu.Lock()
			defer mu.Unlock()
			txt, _ := testcolorize.Color(line)
			fmt.Fprintln(dst, txt)
		}()
	}
}

func Contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
