package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	cli "github.com/antonio-alexander/go-hello-world/internal/cli"
)

func main() {
	//get the present working directory, get
	// any supplied arguments (trimming the
	// command), transform environmental
	// variables to map and then execute the
	// main business logic and if error
	// output and provide an exit status of 1
	pwd, _ := os.Getwd()
	args := os.Args[1:]
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		if s := strings.Split(env, "="); len(s) > 1 {
			switch {
			case len(s) == 2:
				envs[s[0]] = s[1]
			case len(s) > 2:
				envs[s[0]] = strings.Join(s[1:], "=")
			}
		}
	}
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	if err := cli.Main(pwd, args, envs, osSignal); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
