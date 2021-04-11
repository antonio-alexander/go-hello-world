package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	cli "github.com/antonio-alexander/go-hello-world/cli"
)

func main() {
	var envs = make(map[string]string)

	//get the present working directory, get
	// any supplied arguments (trimming the
	// command), transform environmental
	// variables to map and then execute the
	// main business logic and if error
	// output and provide an exit status of 1
	pwd, _ := os.Getwd()
	args := os.Args[1:]
	for _, env := range os.Environ() {
		splitEnv := strings.Split(env, "=")
		envs[splitEnv[0]] = splitEnv[1]
	}
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	if err := cli.Main(pwd, args, envs, osSignal); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
