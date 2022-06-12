package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	//demoURL, err := url.Parse("http://neverssl.com")
	demoURL, err := url.Parse("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(demoURL)
	http.ListenAndServe(":8080", proxy)
}
