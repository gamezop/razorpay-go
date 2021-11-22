package requestor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	ErrFailedToCloseHttpBody   = fmt.Errorf("failed to read http body")
	ErrCannotReadHttpBody      = fmt.Errorf("cannot read HTTP response body")
	ErrFailedToMarshalHttpBody = fmt.Errorf("failed to unmarshal http body")
)

// ReadResponse checks response status code and response body, and tries to unmarshal.
// v is a pointer
func ReadResponse(resp *http.Response, v interface{}) error {
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error().
				Err(err).
				Err(ErrFailedToCloseHttpBody).
				Send()
		}
	}()
	// Try read body
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrCannotReadHttpBody, err.Error())
	}
	if err := json.Unmarshal(buff, v); err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToMarshalHttpBody, err.Error())
	}
	return nil
}
