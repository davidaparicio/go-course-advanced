package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/http2"
)

func init() {
	// https://youtu.be/tWSmUsYLiE4
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	//demoURL, err := url.Parse("http://neverssl.com")
	demoURL, err := url.Parse("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	//proxy := httputil.NewSingleHostReverseProxy(demoURL)
	proxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.Host = demoURL.Host
		req.URL.Host = demoURL.Host
		req.URL.Scheme = demoURL.Scheme
		req.RequestURI = ""
		s, _, _ := net.SplitHostPort(req.RemoteAddr) //8m
		req.Header.Add("X-Forward-For", s)
	
		http2.ConfigureTransport(http.DefaultTransport.(*http.Transport))
		
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			return
		}
		// To maintain the correct Content-Type/Length etc..
		for key, values := range resp.Header {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}

		// To handle data streams
		ticker := time.NewTicker(10 * time.Millisecond)
		done := make(chan bool)
		go func(){
			for {
				select {
				case <- ticker.C: // select case must be send or receive
					rw.(http.Flusher).Flush()
				case <- done:
					return
				}
			}
		}()

		// To handle trailer header
		trailerKeys := []string{}
		for key := range resp.Trailer {
			trailerKeys = append(trailerKeys, key)
		}

		rw.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		rw.WriteHeader(resp.StatusCode)
		io.Copy(rw, resp.Body)

		for key, values := range resp.Trailer {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}
		ticker.Stop()
		close(done) //end the go routine
	})
	// HTTP 1
	//http.ListenAndServe(":8080", proxy)
	// HTTP 2
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", proxy)
}
