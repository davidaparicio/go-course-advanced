// twitter allows users to check the validity and availabity...
package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/internal"
)

// Twitter is an abstraction ...
type Twitter struct {
	Client namecheck.Client
}

const (
	minLen         = 4
	maxLen         = 15
	illegalPattern = "twitter"
)

var legalPattern = regexp.MustCompile("^[0-9A-Z_a-z]*$")

// String returns a string repr of Twitter
func (*Twitter) String() string {
	return "Twitter"
}

// IsValid check whether a username is valid
func (*Twitter) IsValid(username string) bool {
	return internal.IsLongEnough(username, minLen) &&
		isShortEnough(username) &&
		containsNoIllegalPattern(username) &&
		containsOnlyLegalChars(username)
}

// IsAvailable checks whethr username is available on Twitter
func (tw *Twitter) IsAvailable(username string) (bool, error) {
	const tmpl = "https://europe-west6-namechecker-api.cloudfunctions.net/userlookup?username=%s&simulateLatency=1"
	endpoint := fmt.Sprintf(tmpl, url.QueryEscape(username))
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, err
	}
	resp, err := tw.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("unexpected response from API")
	}
	var dto struct {
		Data interface{} `json:"data"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&dto); err != nil {
		return false, err
	}
	// the absence of a data field in the response body indicates the username's availability
	return dto.Data == nil, nil
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(strings.ToLower(username), illegalPattern)
}

func containsOnlyLegalChars(username string) bool {
	return legalPattern.MatchString(username)
}
