package main

import (
	"fmt"
	"regexp"
	"strings"
)

//https://regex101.com/library/RzBwPX
var emailValidPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_ \x60{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsValidEmail checks if it's a valid e-mail
// It returns a boolean.
func IsValidEmail(email string) bool {
	return emailValidPattern.MatchString(email)
}

// IsValidEmail performs only @ check and returns a boolean.
func IsBasicEmail(email string) bool {
	return strings.Contains(strings.ToLower(email), "@")
}

func main() {
	myWeirdEmail := "#!$%&'*+-/=?^_`{}|~@example.org"
	rfcs := `The formal definitions of e-mail addresses are in:
	RFC 5322 (sections 3.2.3 and 3.4.1, obsoletes RFC 2822), RFC 5321, RFC 3696,
	RFC 6531 (permitted characters).
	`
	fmt.Printf("Is the email %q valid? The result is: %t\n", myWeirdEmail, IsValidEmail(myWeirdEmail))
	fmt.Println(rfcs)
}
