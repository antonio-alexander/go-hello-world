package hello

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func Main(pwd string, args []string, envs map[string]string, osSignal chan os.Signal) (err error) {
	var wg sync.WaitGroup

	//generate and create handle func, when connecting, it will use this port
	//indicate via console that the webserver is starting
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, World!")
	})
	server := &http.Server{
		Addr:    "",
		Handler: nil,
	}
	fmt.Println("Starting web server on port 8080")
	stopped := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(stopped)

		if err := server.ListenAndServe(); err != nil {
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
