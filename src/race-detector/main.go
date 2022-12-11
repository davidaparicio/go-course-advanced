package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	counter int64
	mux     *sync.RWMutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	counter++
	mux.Unlock()
	mux.RLock()
	fmt.Fprintln(w, counter)
	mux.RUnlock()
}

func main() {
	mux = &sync.RWMutex{}
	http.HandleFunc("/", handler)
	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
