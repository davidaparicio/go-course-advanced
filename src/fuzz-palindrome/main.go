package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	strs := []string{"foo", "madam", "kayak", "bar", "//--//", "eye"}
	for _, str := range strs {
		fmt.Printf("Is '%s' a palindrome ? %t\n", str, isPalindrome(str))
		//https://pkg.go.dev/fmt#hdr-Printing
	}
}

// https://julien.ponge.org/blog/playing-with-test-fuzzing-in-go/
func isPalindrome(str string) bool {
	first := 0
	last := len(str) - 1
	for first <= last {
		if str[first] != str[last] {
			return false
		}
		first++
		last--
	}
	return true
}

func reverse(str string) string {
	r := []rune(str)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}

func isPalindrom_WithRune(str string) bool {
	if !utf8.ValidString(str) {
		return false
	}
	r := []rune(str)
	first := 0
	last := len(r) - 1
	for first <= last {
		if r[first] != r[last] {
			return false
		}
		first++
		last--
	}
	return true
}
