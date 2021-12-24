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
	CreateFundingAccountBank(ctx context.Context, faBank r.RequestFundingAccountBank) (
		r.FundingAccountBank,
		error,
	)
	Payout(ctx context.Context, faUPI r.RequestPayout) (
		result r.Payout,
		rawResponse string,
		statusCode int,
		err error,
	)
	GetPayout(ctx context.Context, payoutId string) (
		result r.Payout,
		rawResponse string,
		statusCode int,
		err error,
	)

	GetFundingAccountById(ctx context.Context, fundingAccountId string) (
		result r.FundingAccount,
		rawResponse string,
		statusCode int,
		err error,
	)
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
	err = rzp.clientHelper.DoCtx(ctx, rzp.httpClient, api, []string{}, contact, &createdContact)
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
	err = rzp.clientHelper.DoCtx(ctx, rzp.httpClient, api, []string{}, faUPI, &createdFundingAccountUPI)
	return createdFundingAccountUPI, err
}

func (rzp *razorPayClient) CreateFundingAccountBank(ctx context.Context, faBank r.RequestFundingAccountBank) (
	r.FundingAccountBank,
	error,
) {
	err := validate.Struct(faBank)
	if err != nil {
		return r.FundingAccountBank{}, err
	}
	api := requestor.API_FUNDING_ACCOUNT_CREATE

	var createdFundingAccountBank r.FundingAccountBank
	err = rzp.clientHelper.DoCtx(ctx, rzp.httpClient, api, []string{}, faBank, &createdFundingAccountBank)
	return createdFundingAccountBank, err
}

func (rzp *razorPayClient) Payout(ctx context.Context, faUPI r.RequestPayout) (
	result r.Payout,
	rawResponse string,
	statusCode int,
	err error,
) {
	err = validate.Struct(faUPI)
	if err != nil {
		return
	}
	api := requestor.API_PAYOUT_CREATE

	rawResponse, statusCode, err = rzp.clientHelper.DoReturnExtraCtx(ctx, rzp.httpClient, api, []string{}, faUPI, &result)
	return
}

func (rzp *razorPayClient) GetPayout(ctx context.Context, payoutId string) (
	result r.Payout,
	rawResponse string,
	statusCode int,
	err error,
) {
	api := requestor.API_PAYOUT_GET
	rawResponse, statusCode, err = rzp.clientHelper.DoReturnExtraCtx(ctx, rzp.httpClient, api, []string{payoutId}, nil, &result)
	return
}

func (rzp *razorPayClient) GetFundingAccountById(ctx context.Context, fundingAccountId string) (
	result r.FundingAccount,
	rawResponse string,
	statusCode int,
	err error,
) {
	api := requestor.API_FUNDING_ACCOUNT_GET

	rawResponse, statusCode, err = rzp.clientHelper.DoReturnExtraCtx(ctx, rzp.httpClient, api, []string{fundingAccountId}, nil, &result)
	return
}

// please don't use default httpClient
// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
func NewClient(httpClient *http.Client, api, secret string) IRazorPayClient {
	return &razorPayClient{
		httpClient:   httpClient,
		clientHelper: requestor.New(api, secret),
	}
}
