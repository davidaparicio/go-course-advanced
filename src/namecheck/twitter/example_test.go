package twitter_test

import (
	"fmt"
	"net/http"

	"github.com/davidaparicio/go-course-advanced/src/namecheck/twitter"
)

func ExampleTwitter_IsValid() {
	var tw twitter.Twitter
	fmt.Println(tw.IsValid("eczxaw"))
	// Output: true
}

func ExampleTwitter_IsAvailable() {
	t := twitter.Twitter{
		Client: http.DefaultClient,
	}
	fmt.Println(t.IsAvailable("dadideo"))
	fmt.Println(t.IsAvailable("eczxaw"))
	// Output:
	// false <nil>
	// true <nil>
}

// More information https://go.dev/blog/examples | https://go.dev/blog/godoc
