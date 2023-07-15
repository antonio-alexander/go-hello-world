package main

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	_ "embed"
)

// These variables are populated at build time
// REFERENCE: https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
// to find where the variables are...
//
//	go tool nm ./app | grep app
var (
	Version   string
	GitCommit string
	GitBranch string
)

//go:embed data/sample_embed.txt
var embedFolder embed.FS

func main() {
	var wg sync.WaitGroup
	var err error

	defer func() {
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}
	}()

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

	//get the command
	command, ok := envs["COMMAND"]
	if !ok && len(args) > 0 {
		command = args[0]
	}
	//switch over the command argument/environmental variable and
	// execute the appropriate case
	switch strings.TrimSpace(strings.ToLower(strings.ReplaceAll(command, "-", ""))) {
	default:
		//print some common information, including the working
		// directory, arguments and environmental variables,
		// then try to find the command if provied
		fmt.Println("Hello, World!")
		fmt.Printf("Version: \"%s\"\n", Version)
		fmt.Printf("Git Commit: \"%s\"\n", GitCommit)
		fmt.Printf("Git Branch: \"%s\"\n", GitBranch)
		fmt.Printf(" Present Working Directory: %s\n", pwd)
		fmt.Printf(" Arguments: %v\n", args)
		fmt.Printf(" Environmental Variables:\n")
		for key, value := range envs {
			fmt.Printf("  %s: %s\n", key, value)
		}
	case "error":
		//output an error on purpose
		err = errors.New("error on purpose")
	case "panic":
		//panic on purpose
		panic("Panicking on purpose")
	case "cpu":
		var goRoutines = 1

		//start a go routine that runs indefinitely
		// with a default case
		fmt.Println("Starting cpu routine")
		wg.Add(1)
		for i := 0; i < goRoutines; i++ {
			go func(i int) {
				defer wg.Done()

				<-osSignal
				fmt.Printf("Stopping cpu routine %d\n", i+1)
			}(i)
		}
		wg.Wait()
	case "memory":
		var memoryLeak []byte
		var megaByte = make([]byte, 1024*1024)

		//create a memory leak that increases a megabyte per second
		fmt.Println("Starting memory leak routine")
		tLeak := time.NewTicker(time.Second)
		defer tLeak.Stop()
		for {
			select {
			case <-tLeak.C:
				memoryLeak = append(memoryLeak, megaByte...) //nolint:staticcheck
			case <-osSignal:
				fmt.Println("Stopping memory leak routine")
				return
			}
		}
	case "race":
		//create a reace condition by starting to go routines
		// attempting to read from a race condition simultaneously
		// i'm also 99% sure I pulled this code from stack overflow
		// but can't find the comment to give appropriate credit
		// i'll try to update if I ever find it again
		fmt.Println("Racing...")
		m := make(map[string]int, 1)
		m[`foo`] = 1
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				m[`foo`]++
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				m[`foo`]++
			}
		}()
		wg.Wait()
		fmt.Println("Racing complete.")
	case "signal":
		//generate a signal that's triggered when a sigint or
		// term is provided (ctrl+c)
		wg.Add(1)
		go func() {
			defer wg.Done()
			sig := <-osSignal
			fmt.Println()
			fmt.Println(sig)
		}()
		fmt.Println("awaiting signal")
		wg.Wait()
		fmt.Println("exiting")
	case "file":
		var bytes []byte

		sampleReadOnlyFile := filepath.Join(pwd, "./data/sample_read_only.txt")
		if bytes, err = os.ReadFile(sampleReadOnlyFile); err != nil {
			return
		}
		fmt.Printf("Contents of file (%s):\n", sampleReadOnlyFile)
		fmt.Println(string(bytes))
	case "embed":
		var bytes []byte

		fileEmbed := "data/sample_embed.txt"
		bytes, err = embedFolder.ReadFile(fileEmbed)
		if err != nil {
			return
		}
		fmt.Printf("Contents of embed (%s):\n", fileEmbed)
		fmt.Println(string(bytes))
	}
}
