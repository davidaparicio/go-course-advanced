package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var once sync.Once // HL
	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			f := func() { fmt.Printf("Only once: %d\n", i) }
			once.Do(f) // HL
		}()
	}
	wg.Wait()
	// END OMIT
}
