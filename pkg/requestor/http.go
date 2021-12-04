package requestor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type API string

var (
	API_CONTACT_CREATE         API = "API_CONTACT_CREATE"
	API_FUNDING_ACCOUNT_CREATE API = "API_FUNDING_ACCOUNT_CREATE"
	API_PAYOUT_CREATE          API = "API_PAYOUT_CREATE"
	API_PAYOUT_GET             API = "API_PAYOUT_GET"
)

// this can be generic sdk builder
type IRazorPayHttpClientHelper interface {
	GetMethod(api API) string
	GetPath(api API, urlParams []string) string

	AddAuth(req *http.Request)
	Do(httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) error
	DoCtx(ctx context.Context, httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) error
	DoReturnExtra(httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) (
		rawResp string,
		statusCode int,
		err error,
	)
	DoReturnExtraCtx(ctx context.Context, httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) (
		rawResp string,
		statusCode int,
		err error,
	)
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
func (r *razorPayHttpClientHelper) DoReturnExtraCtx(ctx context.Context, httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) (
	rawResp string,
	statusCode int,
	err error,
) {
	res, err := r.doRequest(ctx, httpClient, api, urlParams, reqBody)
	if err != nil {
		return rawResp, statusCode, err
	}
	rawResp, statusCode, err = ReadResponse(res, resp)
	return rawResp, statusCode, err
}
func (r *razorPayHttpClientHelper) DoReturnExtra(httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) (
	rawResp string,
	statusCode int,
	err error,
) {
	return r.DoReturnExtraCtx(context.Background(), httpClient, api, urlParams, reqBody, resp)
}

func (r *razorPayHttpClientHelper) DoCtx(ctx context.Context, httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) error {
	res, err := r.doRequest(ctx, httpClient, api, urlParams, reqBody)
	if err != nil {
		return err
	}
	_, _, err = ReadResponse(res, resp)
	if err != nil {
		return err
	}
	return nil
}
func (r *razorPayHttpClientHelper) Do(httpClient *http.Client, api API, urlParams []string, reqBody interface{}, resp interface{}) error {
	return r.DoCtx(context.Background(), httpClient, api, urlParams, reqBody, resp)
}

func (r *razorPayHttpClientHelper) doRequest(ctx context.Context, httpClient *http.Client, api API, urlParams []string, reqBody interface{}) (*http.Response, error) {
	bArr, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	log.Trace().
		Str("api", "rzHttpClient.doRequest").
		Str("requestBody", string(bArr)).
		Msg("send message")

	httpReq, err := http.NewRequestWithContext(
		ctx,
		r.GetMethod(api),
		r.GetPath(api, urlParams),
		bytes.NewReader(bArr),
	)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "gzp-rzpay")
	r.AddAuth(httpReq)
	if err != nil {
		return nil, err
	}
	res, err := httpClient.Do(httpReq)
	return res, err
}

func (r *razorPayHttpClientHelper) GetMethod(api API) string {
	switch api {
	case API_CONTACT_CREATE:
		return "POST"
	case API_FUNDING_ACCOUNT_CREATE:
		return "POST"
	case API_PAYOUT_CREATE:
		return "POST"
	case API_PAYOUT_GET:
		return "GET"
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
	case API_FUNDING_ACCOUNT_CREATE:
		return "/fund_accounts"
	case API_PAYOUT_CREATE:
		return "/payouts"
	case API_PAYOUT_GET:
		return fmt.Sprintf("/payouts/%s", urlParams[0])
	default:
		panic(fmt.Sprintf("unknown path %s", api))
	}
}

func (r *razorPayHttpClientHelper) AddAuth(req *http.Request) {
	req.SetBasicAuth(r.api, r.secret)
}
