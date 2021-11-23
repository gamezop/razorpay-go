package requestor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type API string

var (
	API_CONTACT_CREATE API = "API_CONTACT_CREATE"
)

type IRazorPayHttpClientHelper interface {
	GetMethod(api API) string
	GetPath(api API, urlParams []string) string

	AddAuth(req *http.Request)
	Do(httpClient *http.Client, api API, reqBody interface{}, resp interface{}) error
}

type razorPayHttpClientHelper struct {
	api    string
	secret string
}

func New(api, secret string) IRazorPayHttpClientHelper {
	return &razorPayHttpClientHelper{
		api:    api,
		secret: secret,
	}
}

// resp is expecting a pointer
func (r *razorPayHttpClientHelper) Do(httpClient *http.Client, api API, reqBody interface{}, resp interface{}) error {
	bArr, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	log.Trace().
		Str("api", "rzHttpClient.Do").
		Str("requestBody", string(bArr)).
		Msg("send message")

	httpReq, err := http.NewRequest(
		r.GetMethod(api),
		r.GetPath(api, nil),
		bytes.NewReader(bArr),
	)
	httpReq.Header.Set("User-Agent", "gzp-rzpay")
	httpReq.Header.Set("Content-Type", "application/json")
	r.AddAuth(httpReq)
	if err != nil {
		return err
	}
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	err = ReadResponse(res, resp)
	if err != nil {
		return err
	}
	return nil
}

func (r *razorPayHttpClientHelper) GetMethod(api API) string {
	switch api {
	case API_CONTACT_CREATE:
		return "POST"
	default:
		panic(fmt.Sprintf("unknown path %s", api))
	}
}

const (
	BASE_URL = "https://api.razorpay.com/v1"
)

func (r *razorPayHttpClientHelper) GetPath(api API, urlParams []string) string {
	return fmt.Sprintf("%s%s", BASE_URL, r.getPath(api, urlParams))
}

func (r *razorPayHttpClientHelper) getPath(api API, urlParams []string) string {
	switch api {
	case API_CONTACT_CREATE:
		return "/contacts"
	default:
		panic(fmt.Sprintf("unknown path %s", api))
	}
}

func (r *razorPayHttpClientHelper) AddAuth(req *http.Request) {
	req.SetBasicAuth(r.api, r.secret)
}
