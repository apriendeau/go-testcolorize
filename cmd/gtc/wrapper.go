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

	cmd := exec.Command("go", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go io.Copy(stdin, os.Stdin)

	var wg sync.WaitGroup
	wg.Add(2)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		colortest(os.Stdout, stdout)
		wg.Done()
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		colortest(os.Stderr, stderr)
		wg.Done()
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c)
		for sig := range c {
			cmd.Process.Signal(sig)
		}
	}()

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		os.Exit(1)
	}
}

func colortest(dst io.Writer, src io.Reader) {
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
