package main

import (
	"fmt"
)

// START OMIT
type Location struct {
	Long float64
	Lat  float64
	_    struct{} // HL
}

func main() {
	strasbourg := Location{7.7521, 48.5734} // doesn't compile // HL
	fmt.Println(strasbourg)
}

// END OMIT
