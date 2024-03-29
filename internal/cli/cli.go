package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	internal "github.com/antonio-alexander/go-hello-world/internal"
)

//Main
func Main(pwd string, args []string, envs map[string]string, osSignal chan os.Signal) (err error) {
	var wg sync.WaitGroup
	var command string
	var ok bool

	//print hello world, then attempt to get the command
	fmt.Println("Hello, World!")
	if command, ok = envs["Command"]; !ok {
		if ok = (len(args) > 0); ok {
			command = args[0]
		}
	}
	//switch over the command argument/environmental variable and
	// execute the appropriate case
	switch strings.TrimSpace(strings.ToLower(strings.ReplaceAll(command, "-", ""))) {
	default:
		//print some common information, including the working
		// directory, arguments and environmental variables,
		// then try to find the command if provied
		fmt.Printf("Version: \"%s\"\n", internal.Version)
		fmt.Printf("Git Commit: \"%s\"\n", internal.GitCommit)
		fmt.Printf("Git Branch: \"%s\"\n", internal.GitBranch)
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
				memoryLeak = append(memoryLeak, megaByte...)
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
	}

	return
}
