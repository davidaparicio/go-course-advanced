package twitter_test

import (
	"fmt"
	"testing"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/twitter"
)

var _ namecheck.Checker = (*twitter.Twitter)(nil)

func ExampleTwitter_IsValid() {
	var tw twitter.Twitter
	fmt.Println(tw.IsValid("jub0bs"))
	// Output: true
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		label    string
		username string
		want     bool
	}{
		{label: "too long", username: "obviously_longer_than_15_chars", want: false},
		{label: "too short", username: "foo", want: false},
		{label: "illegal patterns", username: "FtWittEroo", want: false},
		{label: "illegal chars", username: "jub0bs-", want: false},
		{label: "valid", username: "jub0bs", want: true},
	}
	for _, c := range cases {
		f := func(t *testing.T) {
			var tw twitter.Twitter
			got := tw.IsValid(c.username)
			if got != c.want {
				t.Errorf("twitter.IsValid(%s): got %t; want %t", c.username, got, c.want)
			}
		}
		t.Run(c.label, f)
	}
}
