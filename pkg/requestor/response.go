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
	ErrNon2XXStatusCode        = fmt.Errorf("got non 2xx status code")
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
	log.Trace().
		Str("api", "rzHttpClient.Do.ReadResponse").
		Str("responseBody", string(buff)).
		Send()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("%w: status:%d body: %s", ErrNon2XXStatusCode, resp.StatusCode, string(buff))
	}
	if err := json.Unmarshal(buff, v); err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToMarshalHttpBody, err.Error())
	}
	return nil
}
