package main

import (
	"testing"
	"fmt"
)

func TestGreetingsByLocale(t *testing.T) {
	testCases := []struct {
		locale   string
		excepted string
	}{
		{"", "Sorry, we didn't identified you"},
		{"en", "Hello"},
		{"pt", "Olá"},
	}
	for _, tc := range testCases {
		//Rename sub-test | https://go.dev/blog/subtestshttps://gobyexample.com/testing
		testname := fmt.Sprintf("%s", tc.locale)
		t.Run(testname, func(t *testing.T) {
			greetings := GreetingsByLocale(tc.locale)
			if greetings != tc.excepted {
				t.Errorf("greeting(%s) = %v; want %s", tc.locale, greetings, tc.excepted)
			}
		})
	}
}
