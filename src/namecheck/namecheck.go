package namecheck

import (
	"context"
	"fmt"
	"net/http"
)

type Validator interface {
	IsValid(username string) bool
}

type Availabler interface {
	IsAvailable(ctx context.Context, username string) (bool, error)
}

type Checker interface {
	Validator
	Availabler
	fmt.Stringer
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type ErrUnknownAvailability struct {
	Username string
	Platform string
	Cause    error
}

func (err *ErrUnknownAvailability) Error() string {
	const tmpl = "unknown availability of %q on %s: %v"
	return fmt.Sprintf(tmpl, err.Username, err.Platform, err.Cause)
}

func (err *ErrUnknownAvailability) Unwrap() error {
	return err.Cause
}
