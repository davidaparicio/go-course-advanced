package magicnumbers

import (
	"fmt"
	"net/http"
	"net/url"
)

// START OMIT
func IsAvailableOnGitHub(username string) (bool, error) {
	endpoint := fmt.Sprintf("https://github.com/%s", url.PathEscape(username))
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == 404, nil
}

// END OMIT
