package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// Passed by "go build -ldflags" for the show version
	commit string
	date   string
)

type CounterHandler struct {
    counter int
    filename string
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("ServeHTTP")
    ct.counter++
    log.Println("Counter:", ct.counter)
    http.ServeFile(w, r, ct.filename)
}

func main() {

    version := flag.Bool("version", false, "version")
    port := flag.String("port", ":8888", "TCP port for http server")
    path := flag.String("path","/openapi/timesync","http path handle")
    fn := flag.String("fn","/home/das/samsung_time_response_bytes_zero","filename containing the response")

	flag.Parse()

    if *version {
		fmt.Println("commit:", commit, "\tdate(UTC):", date)
		os.Exit(0)
	}

    fmt.Println("fn:",*fn)

    th := &CounterHandler{
        counter: 0,
        filename: *fn,
    }
    http.Handle(*path, th)
    http.ListenAndServe(*port, nil)
}