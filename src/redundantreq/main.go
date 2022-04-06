package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	data string
	err  error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	dcs := []string{"Strasbourg", "Roubaix", "Paris"}
	results := make(chan Result, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for _, dc := range dcs {
		go fetchFrom(ctx, results, dc)
		select {
		case r := <-results:
			cancel()
			fmt.Println(r)
			return
		case <-time.After(250 * time.Millisecond):
		}
	}
}

func fetchFrom(ctx context.Context, results chan<- Result, dc string) {
	timer := time.NewTimer(time.Duration(rand.Intn(1000)) * time.Millisecond)
	var res Result
	select {
	case <-ctx.Done():
		fmt.Println("cancelling request to", dc)
		timer.Stop()
		res.err = ctx.Err()
	case <-timer.C:
		res.data = fmt.Sprintf("data from %s", dc)
	}
	results <- res

}
