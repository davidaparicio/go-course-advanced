package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Result string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func get(ctx context.Context, results chan Result, dc string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := fetchFrom(ctx, dc)
	if err != nil {
		return
	}
	results <- res
}

func main() {
	//dc := "Strasbourg"
	dcs := []string{"Strasbourg", "Roubaix", "Paris", "Sydney", "Beauharnois"}
	results := make(chan Result)

	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(dcs), func(i, j int) { dcs[i], dcs[j] = dcs[j], dcs[i] })

	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	defer wg.Wait()
	//wg.Add(len(dcs))

	for _, dc := range dcs {
		dc := dc
		wg.Add(1)
		go get(ctx, results, dc, &wg)

		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stderr, ctx.Err())
			return
		case res := <-results:
			cancel()
			fmt.Println(res)
			return
		case <-time.After(200 * time.Millisecond):
			continue
		}
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
