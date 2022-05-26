package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, err := Reverse(tc.in)
		if err != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

//dummy test that executes main()
//https://stackoverflow.com/a/51277293
func TestRunMain(t *testing.T) {
	main()
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			//return
			// Rather than returning, you can also call t.Skip()
			// to stop the execution of that fuzz input.
			t.Skip()
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			t.Skip()
		}
		//Log added to understand better
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func benchmarkReverse(b *testing.B, s []string) {
	// b is lowcase to avoid the error: wrong signature for BenchmarkReverse,
	// must be: func BenchmarkReverse(b *testing.B)
	for i := 0; i < b.N; i++ {
		for _, str := range s {
			Reverse(str)
		}
	}
}

func BenchmarkReverse_Simple(b *testing.B) {
	s := []string{"", "bar", "eye", "foo"}
	benchmarkReverse(b, s)
}

func BenchmarkReverse_BigInputs(b *testing.B) {
	s := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent lacinia quam in elit laoreet rutrum ac id ante. Aliquam at mollis felis. Nam congue orci non vestibulum blandit. Fusce efficitur ligula lorem, eu cursus enim varius eget. Mauris hendrerit auctor mattis. Pellentesque et luctus nibh. Integer eget mi sed dolor.",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. In aliquam purus volutpat mattis bibendum. Donec eget feugiat diam. Phasellus id diam dolor. Proin quis lobortis.",
	}
	benchmarkReverse(b, s)
}
