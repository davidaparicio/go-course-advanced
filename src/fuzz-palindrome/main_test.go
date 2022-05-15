package main

import "testing"

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
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
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
		t1 := isPalindrome(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}

func FuzzIsPalindrom_WithRune(f *testing.F) {
	f.Add("kayak")
	f.Fuzz(func(t *testing.T, str string) {
		t1 := isPalindrom_WithRune(str)
		t2 := reverse(str) == str
		if t1 != t2 {
			t.Fail()
		}
	})
}
