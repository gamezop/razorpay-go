package resource

import "encoding/json"

type FundingAccountType string

const (
	FA_UPI  FundingAccountType = "vpa"
	FA_BANK FundingAccountType = "bank_account"
)

var ALL_FA = []string{
	string(FA_UPI),
	string(FA_BANK),
}

type RequestFundingAccount struct {
	AccountType string `json:"account_type" validate:"required"`
	ContactID   string `json:"contact_id" validate:"required"`
}

// UPI
type RequestFundingAccountUPIDetails struct {
	Address string `json:"address" validate:"required"`
}

// BANK
type RequestFundingAccountBankDetails struct {
	Name          string `json:"name"`
	Ifsc          string `json:"ifsc"`
	AccountNumber string `json:"account_number"`
}

type RequestFundingAccountUPI struct {
	RequestFundingAccount
	UPIDetails RequestFundingAccountUPIDetails `json:"vpa" validate:"required"`
}

type RequestFundingAccountBank struct {
	RequestFundingAccount
	BankDetails RequestFundingAccountBankDetails `json:"bank_account" validate:"required"`
}

type fundingAccountCommon struct {
	ID          string  `json:"id"`
	Entity      string  `json:"entity"`
	ContactID   string  `json:"contact_id"`
	AccountType string  `json:"account_type"`
	Active      bool    `json:"active"`
	BatchID     *string `json:"batch_id"`
	CreatedAt   int     `json:"created_at"`
}

func (f fundingAccountCommon) GetContactId() string {
	return f.ContactID
}
func (f fundingAccountCommon) GetFundingAccountId() string {
	return f.ID
}

func (f fundingAccountCommon) GetAccountType() string {
	return f.AccountType
}

type FundingAccountUPI struct {
	fundingAccountCommon
	Vpa struct {
		Username string `json:"username"`
		Handle   string `json:"handle"`
		Address  string `json:"address"`
	} `json:"vpa"`
}

func (f FundingAccountUPI) GetIdentifierForAccount() string {
	return f.Vpa.Address
}
func (f fundingAccountCommon) GetMode() string {
	return string(PAYOUT_MODE_UPI)
}

type FundingAccountBank struct {
	fundingAccountCommon
	BankAccount struct {
		Ifsc          string `json:"ifsc"`
		BankName      string `json:"bank_name"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		// Notes         []interface{} `json:"notes"`
	} `json:"bank_account"`
}

func (f FundingAccountBank) GetIdentifierForAccount() string {
	return f.BankAccount.AccountNumber
}

func (f FundingAccountBank) GetMode() string {
	return string(PAYOUT_MODE_IMPS)
}

type FundingAccount struct {
	fundingAccountCommon
	BankAccount json.RawMessage `json:"bank_account"`
	Vpa         json.RawMessage `json:"vpa"`
}
