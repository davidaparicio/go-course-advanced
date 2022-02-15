package github_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/stub"
)

var _ namecheck.Checker = (*github.GitHub)(nil)

func TestUsernameTooLong(t *testing.T) {
	var gh = github.GitHub{}
	username := "obviously-longer-than-39-chars-skjdhsdkhfkshkfshdkjfhksdjhf"
	want := false
	got := gh.IsValid(username)
	if got != want {
		t.Errorf(
			"IsValid(%s) = %t; want %t",
			username,
			got,
			want,
		)
	}
}

func TestIsAvailable200(t *testing.T) {
	gh := github.GitHub{
		Client: stub.ClientWithStatusCode(http.StatusOK),
	}
	username := "whatever"
	avail, _ := gh.IsAvailable(context.Background(), username)
	if avail {
		t.Error("unexpected availability")
	}
}
