package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	//r := []byte(s) // bad version
	if !utf8.ValidString(s) { // good version V2
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s) // good version V1
	//FOR DEBUGGING
	/*fmt.Printf("input: %q\n", s)
	fmt.Printf("runes: %q\n", r)*/
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
