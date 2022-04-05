package main

import (
	"fmt"
	"regexp"

	"mymodulepath/bimap"
)

var re regexp.Regexp

func init() {
	//sdfgsdfg
	///sdfg
	//sdfg
	re = //...
}

func main() {
	init() 
	var frToEn bimap.Bimap
	frToEn.Store("un", "one")
	frToEn.Store("deux", "two")
	fmt.Println(frToEn)
}
