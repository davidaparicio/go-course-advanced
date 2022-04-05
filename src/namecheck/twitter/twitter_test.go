package twitter_test

import (
	"testing"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/twitter"
)

var _ namecheck.Checker = (*twitter.Twitter)(nil)

func TestIsValid(t *testing.T) {
	cases := []struct {
		label    string
		username string
		want     bool
	}{
		{label: "too long", username: "obviously_longer_than_15_chars", want: false},
		{label: "too short", username: "foo", want: false},
		{label: "valid", username: "jub0bs", want: true},
		{label: "contains twitter", username: "FtWittEroo", want: false},
		{label: "contains illegal chars", username: "jub0bs-", want: false},
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
