package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/?msg=script>window.alert("sometext");</script>
	// Without => Content-Type: text/html; charset=utf-8
	// The patch
	w.Header().Set("Content-Type", "text/plain")
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		http.Error(w, "'msg' query param is required", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, msg)
	return
}
