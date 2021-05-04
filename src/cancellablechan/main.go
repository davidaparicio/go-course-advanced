package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	const timeout = 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printInts(ctx)
	}()
	// do more work (omitted)...
	wg.Wait()
}

// START OMIT
// main is unchanged

func printInts(ctx context.Context) {
	ticker := time.NewTicker(100 * time.Millisecond)
	var i int
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println(i)
			i++
		}
	}
}

// END OMIT
