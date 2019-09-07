package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/c-bata/measure"

	_ "expvar"
	_ "net/http/pprof"
)

func main() {
	// measure csv
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6070", &measure.Handler{}))
	}()
	// default servemux
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	err := http.ListenAndServe("0.0.0.0:8080", &hello{})
	if err != nil {
		panic(err)
	}
}

type hello struct{}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer measure.Start(r.URL.Path).Stop()
	time.Sleep(100 * time.Millisecond)
	_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
}
