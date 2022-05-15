package main

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{
			str:  "eye",
			want: true,
		},
		{
			str:  "1221",
			want: true,
		},
		{
			str:  "//--//",
			want: true,
		},
		{
			str:  "foo",
			want: false,
		},
	}
	var desc strings.Builder
	for _, tt := range tests {
		desc.WriteString(tt.str)
		// https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
		desc.WriteString(" isPalindrome")
		// https://go.dev/blog/subtests
		t.Run(desc.String(), func(t *testing.T) {
			if got := isPalindrome(tt.str); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
		desc.WriteString(" 2")
		t.Run(desc.String(), func(t *testing.T) {
			if got := isPalindrome(tt.str); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	strings := []string{"", "bar", "eye", "foo"}
	for i := 0; i < b.N; i++ {
		for _, str := range strings {
			isPalindrome(str)
		}
	}
}

func BenchmarkIsPalindrome_BigInputs(b *testing.B) {
	strings := []string{"tattarrattat", "saippuakivikauppias"}
	for i := 0; i < b.N; i++ {
		for _, str := range strings {
			isPalindrome(str)
		}
	}
}

func FuzzIsPalindrome(f *testing.F) {
	f.Add("kayak")
	f.Fuzz(func(t *testing.T, str string) {
		t1 := IsPalindrome_WithRune_Final(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}

/*
func FuzzIsPalindrome_1(f *testing.F) {
	f.Add("kayak")
	f.Fuzz(func(t *testing.T, str string) {
		t1 := isPalindrome(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}

func FuzzIsPalindrome_2(f *testing.F) {
	f.Add("kayak")
	f.Fuzz(func(t *testing.T, str string) {
		t1 := isPalindrome_Fixed(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}

func FuzzIsPalindrome_3(f *testing.F) {
	f.Add("kayak")
	f.Fuzz(func(t *testing.T, str string) {
		t1 := isPalindrom_WithRune(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}

func FuzzIsPalindrome_4(f *testing.F) {
	f.Add("kayak")
	f.Fuzz(func(t *testing.T, str string) {
		t1 := IsPalindrome_WithRune_Final(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}
*/
