package client

import (
	"context"
	"net/http"

	"github.com/gamezop/razorpay-go/pkg/requestor"
	r "github.com/gamezop/razorpay-go/pkg/resource"
	"github.com/go-playground/validator/v10"
)

type IRazorPayClient interface {
	CreateContact(ctx context.Context, contact r.RequestCreateContact) (r.Contact, error)
	CreateFundingAccountUPI(ctx context.Context, faUPI r.RequestFundingAccountUPI) (
		r.FundingAccountUPI,
		error,
	)
	Payout(ctx context.Context, faUPI r.RequestPayout) (r.Payout, error)
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

func (rzp *razorPayClient) CreateContact(ctx context.Context, contact r.RequestCreateContact) (r.Contact, error) {
	err := validate.Struct(contact)
	if err != nil {
		return r.Contact{}, err
	}
	api := requestor.API_CONTACT_CREATE

	var createdContact r.Contact
	err = rzp.clientHelper.Do(rzp.httpClient, api, contact, &createdContact)
	return createdContact, err
}

func (rzp *razorPayClient) CreateFundingAccountUPI(ctx context.Context, faUPI r.RequestFundingAccountUPI) (
	r.FundingAccountUPI,
	error,
) {
	err := validate.Struct(faUPI)
	if err != nil {
		return r.FundingAccountUPI{}, err
	}
	api := requestor.API_FUNDING_ACCOUNT_CREATE

	var createdFundingAccountUPI r.FundingAccountUPI
	err = rzp.clientHelper.Do(rzp.httpClient, api, faUPI, &createdFundingAccountUPI)
	return createdFundingAccountUPI, err
}

func (rzp *razorPayClient) Payout(ctx context.Context, faUPI r.RequestPayout) (
	r.Payout,
	error,
) {
	err := validate.Struct(faUPI)
	if err != nil {
		return r.Payout{}, err
	}
	api := requestor.API_PAYOUT_CREATE

	var createdFundingAccountUPI r.Payout
	err = rzp.clientHelper.Do(rzp.httpClient, api, faUPI, &createdFundingAccountUPI)
	return createdFundingAccountUPI, err
}

// please don't use default httpClient
// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
func NewClient(httpClient *http.Client, api, secret string) IRazorPayClient {
	return &razorPayClient{
		httpClient:   httpClient,
		clientHelper: requestor.New(api, secret),
	}
}
