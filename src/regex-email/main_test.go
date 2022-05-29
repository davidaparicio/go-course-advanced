package main

import (
	"net/mail"
	"testing"
	"unicode/utf8"
)

//Don’t Test Main: dummy test that executes main()
//https://fossa.com/blog/golang-best-practices-testing-go/#Don't
//https://stackoverflow.com/a/51277293
func TestRunMain(t *testing.T) {
	main()
}

func TestIsValidEmail(t *testing.T) {
	testcases := []struct {
		in    string
		want  bool
		basic bool
	}{
		{"Hello, world", false, false},
		{" ", false, false},
		{"!12345", false, false},
		{"a@a.fr", true, true},
		{"hello@world.com", true, true},
		{"test@abc.net", true, true},
		//The following emails have been taken from https://stackoverflow.com/a/67686133
		{"test44$@gmail.com", true, true}, //for the current regex
		{"test-email.com", false, false},
		{"test+email@test.com", true, true},
		//The following emails have been taken from https://stackoverflow.com/a/38787343
		{"hei@やる.ca", false, true}, //for the current regex
		{"\"very.(),:;<>[]\".VERY.\"very@\"very\".unusual\"@strange.example.com", false, true}, //for the current regex
		{"#!$%&'*+-/=?^_`{}|~@example.org", true, true},
		{"\"()<>[]:,;@\\\"!#$%&'-/=?^_`{}| ~.a\"@example.org", false, true}, //for the current regex
		{"\" \"@example.org", false, true},                                  //for the current regex
		{"example@localhost", true, true},
		{"user@[IPv6:2001:db8::1]", false, true}, //for the current regex
		{"email@123.123.123.123", true, true},
		//Got from the fuzzing test
		{"examzle@localhos ", false, true},
	}
	// Regex validation
	for _, tc := range testcases {
		got := IsValidEmail(tc.in)
		if got != tc.want {
			t.Errorf("IsValidEmail: %q got %t, want %t", tc.in, got, tc.want)
		}
	}
	// @ contains check
	for _, tc := range testcases {
		got := IsBasicEmail(tc.in)
		if got != tc.basic {
			t.Errorf("IsBasicEmail: %q got %t, want %t", tc.in, got, tc.basic)
		}
	}
}

func FuzzIsValidEmail(f *testing.F) {
	testcases := []string{"hello@world.com", "test@abc.net", "example@localhost"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, email string) {
		if !utf8.ValidString(email) {
			t.Errorf("Fuzzer produced invalid UTF-8 string %q", email)
		}
		res1 := IsValidEmail(email)
		_, err := mail.ParseAddress(email) //https://stackoverflow.com/a/66624104
		if err != nil {
			t.Skip() //Error found, stop the execution of that fuzz input.
		}
		if !res1 {
			t.Errorf("@: %q IsValidEmail: %t, mail.ParseAddress: %t", email, res1, true)
		}
	})
}

func benchmarkIsValidEmail(b *testing.B, s []string) {
	// b is lowcase to avoid the error: wrong signature for BenchmarkReverse,
	// must be: func BenchmarkReverse(b *testing.B)
	for i := 0; i < b.N; i++ {
		for _, str := range s {
			IsValidEmail(str)
		}
	}
}

func BenchmarkIsValidEmail_Simple(b *testing.B) {
	s := []string{"Hello, world", " ", "!12345", "a@a.fr", "hello@world.com", "test@abc.net"}
	benchmarkIsValidEmail(b, s)
}

func BenchmarkIsValidEmail_BigInputs(b *testing.B) {
	s := []string{
		"contact-admin-hello-webmaster-info-services-peter-crazy-but-oh-so-ubber-cool-english-alphabet-loverer-abcdefghijklmnopqrstuvwxyz@please-try-to.send-me-an-email-if-you-can-possibly-begin-to-remember-this-coz.this-is-the-longest-email-address-known-to-man-but-to-be-honest.this-is-such-a-stupidly-long-sub-domain-it-could-go-on-forever.pacraig.com",
		"#!$%&'*+-/=?^_`{}|~@example.org",
		"\"very.(),:;<>[]\".VERY.\"very@\"very\".unusual\"@strange.example.com",
		"\"()<>[]:,;@\\\"!#$%&'-/=?^_`{}| ~.a\"@example.org",
	}
	benchmarkIsValidEmail(b, s)
}

func benchmarkIsBasicEmail(b *testing.B, s []string) {
	// b is lowcase to avoid the error: wrong signature for BenchmarkReverse,
	// must be: func BenchmarkReverse(b *testing.B)
	for i := 0; i < b.N; i++ {
		for _, str := range s {
			IsBasicEmail(str)
		}
	}
}

func BenchmarkIsBasicEmail_Simple(b *testing.B) {
	s := []string{"Hello, world", " ", "!12345", "a@a.fr", "hello@world.com", "test@abc.net"}
	benchmarkIsBasicEmail(b, s)
}

func BenchmarkIsBasicEmail_BigInputs(b *testing.B) {
	s := []string{
		"contact-admin-hello-webmaster-info-services-peter-crazy-but-oh-so-ubber-cool-english-alphabet-loverer-abcdefghijklmnopqrstuvwxyz@please-try-to.send-me-an-email-if-you-can-possibly-begin-to-remember-this-coz.this-is-the-longest-email-address-known-to-man-but-to-be-honest.this-is-such-a-stupidly-long-sub-domain-it-could-go-on-forever.pacraig.com",
		"#!$%&'*+-/=?^_`{}|~@example.org",
		"\"very.(),:;<>[]\".VERY.\"very@\"very\".unusual\"@strange.example.com",
		"\"()<>[]:,;@\\\"!#$%&'-/=?^_`{}| ~.a\"@example.org",
	}
	benchmarkIsBasicEmail(b, s)
}
