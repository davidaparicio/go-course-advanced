package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/davidaparicio/go-course-advanced/src/namecheck"
	"github.com/davidaparicio/go-course-advanced/src/namecheck/github"
	"github.com/davidaparicio/go-course-advanced/src/namecheck/twitter"
)

type Status int

type Result struct {
	Username  string
	Platform  string
	Valid     bool
	Available bool
}

const (
	Unknown Status = iota
	Active
	Suspended
	Available
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("username args is required")
	}
	username := os.Args[1]

	var checkers []namecheck.Checker
	for i := 0; i < 3; i++ {
		t := &twitter.Twitter{
			Client: http.DefaultClient,
		}
		g := &github.GitHub{
			Client: http.DefaultClient,
		}
		checkers = append(checkers, t, g)
	}
	results := make(chan Result, len(checkers))
	errc := make(chan error, len(checkers))
	var wg sync.WaitGroup
	wg.Add(len(checkers))
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for _, checker := range checkers {
		go check(ctx, checker, username, &wg, results, errc)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for {
		select {
		case err := <-errc:
			//OLD
			//const tmpl = "namecheck: some error occurred: %s\n"
			//fmt.Fprintf(os.Stderr, tmpl, err)
			//NEW since the Go 1.13
			/*type wrapper interface{ Unwrap() error }
			if err, ok := err.(wrapper); ok { // err has a cause
				// call err.Unwrap to access the error that caused err
				fmt.Println(err.Unwrap())
			}*/
			var uae *namecheck.UnknownAvailabilityError
			// Errors.Is for default errrors
			if errors.As(err, &uae) {
				fmt.Println(uae.Platform, uae.Username)
			}
			fmt.Println(err)
			return
		case res, ok := <-results:
			if !ok {
				return
			}
			fmt.Println(res)
		}
	}
}

func check(
	ctx context.Context,
	checker namecheck.Checker,
	username string,
	wg *sync.WaitGroup,
	results chan<- Result,
	errc chan<- error,
) {
	defer wg.Done()
	res := Result{
		Username: username,
		Platform: checker.String(),
	}
	res.Valid = checker.IsValid(username)
	if !res.Valid {
		results <- res
		return
	}
	avail, err := checker.IsAvailable(ctx, username)
	if err != nil {
		errc <- err
		return
	}
	res.Available = avail
	results <- res
}
