package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct{}

func (s *Server) Foo() http.HandlerFunc {
	var msg string
	var once sync.Once
	f := func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() { msg = s.initializeFoo() })
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, msg)
	}
	return f
}

func (*Server) initializeFoo() string {
	time.Sleep(30 * time.Second) // simulate work
	return "Foo!"
}

func (*Server) Bar() http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Bar!")
	}
	return f
}

func main() {
	var srv Server
	http.Handle("/bar", srv.Bar())
	http.Handle("/foo", srv.Foo())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
