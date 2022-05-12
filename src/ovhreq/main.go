package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Result string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	dc := "Strasbourg"
	results := make(chan Result)
	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()
	go func() {
		res, err := fetchFrom(ctx, dc)
		if err != nil {
			return
		}
		results <- res
	}()
	select {
	case <-ctx.Done():
		fmt.Fprintln(os.Stderr, ctx.Err())
		return
	case res := <-results:
		fmt.Println(res)
	}
}

func fetchFrom(ctx context.Context, dc string) (Result, error) {
	timer := time.NewTimer(time.Duration(rand.Intn(1000)) * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("cancelling request to", dc)
		timer.Stop()
		return "", ctx.Err()
	case <-timer.C:
		break
	}
	res := Result(fmt.Sprintf("data from %s", dc))
	return res, nil
}
