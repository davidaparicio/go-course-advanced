package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// START OMIT
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

func printInts(ctx context.Context) {
	for i := 0; ; i++ {
		if ctx.Err() != nil { // HL
			return
		}
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

// END OMIT
