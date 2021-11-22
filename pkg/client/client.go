package client

import (
	"context"
	"net/http"

	"github.com/gamezop/razorpay-go/pkg/requestor"
	rC "github.com/gamezop/razorpay-go/pkg/resource/contact"
	"github.com/go-playground/validator/v10"
)

type IRazorPayClient interface {
	CreateContact(ctx context.Context, contact rC.RequestCreateContact) (rC.Contact, error)
	// CreateFundingAccount(ctx,)
}

type razorPayClient struct {
	httpClient   *http.Client
	clientHelper requestor.IRazorPayHttpClientHelper
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (r *razorPayClient) CreateContact(ctx context.Context, contact rC.RequestCreateContact) (rC.Contact, error) {
	err := validate.Struct(contact)
	if err != nil {
		return rC.Contact{}, err
	}
	api := requestor.API_CONTACT_CREATE

	var createdContact rC.Contact
	err = r.clientHelper.Do(r.httpClient, api, contact, &createdContact)
	return createdContact, err
}

// please don't use default httpClient
// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
func NewClient(httpClient *http.Client, api, secret string) IRazorPayClient {
	return &razorPayClient{
		httpClient:   httpClient,
		clientHelper: requestor.New(api, secret),
	}
}
