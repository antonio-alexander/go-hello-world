package hello

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	internal "github.com/antonio-alexander/go-hello-world/internal"
)

func Main(pwd string, args []string, envs map[string]string, osSignal chan os.Signal) (err error) {
	var wg sync.WaitGroup

	//generate and create handle func, when connecting, it will use this port
	//indicate via console that the webserver is starting
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, World!\nVersion: \"%s\"\nGit Commit: \"%s\"\nGit Branch: \"%s\"\n", internal.Version, internal.GitCommit, internal.GitBranch)
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	fmt.Printf("starting web server on :8080")
	stopped := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(stopped)

		if err = server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
	select {
	case <-stopped:
	case <-osSignal:
		err = server.Close()
	}
	wg.Wait()

	return
}
