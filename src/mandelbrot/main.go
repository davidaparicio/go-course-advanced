package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

const (
	output = "out.png"
	width  = 2048
	height = 2048
)

func main() {
	profile, _ := os.Create("trace.out")
	trace.Start(profile)
	defer trace.Stop()
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	img := create(width, height)

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func create(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	sem := make(chan struct{}, 16)
	for i := 0; i < width; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			for j := 0; j < height; j++ {
				m.Set(i, j, pixel(i, j, width, height))
			}
		}(i)
	}
	wg.Wait()
	/*
		//1st optimization
		var wg sync.WaitGroup
		wg.Add(width)
		for i := 0; i < width; i++ {
			go func(i int) {
				defer wg.Done()
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}(i)
		}
		wg.Wait()
	*/
	return m
}

func pixel(i, j, width, height int) color.Color {
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
