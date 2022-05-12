package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	defer signal.Reset(os.Interrupt)
	foo(ctx)
}

// START OMIT
func foo(ctx context.Context) {
	// proposal: Go 2: time: deprecate time.Tick(d Duration) #37144 
	// https://github.com/golang/go/issues/37144
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop() // HL
			return
		case v := <-ticker.C:
			fmt.Println(v)
		}
	}
}

// END OMIT
