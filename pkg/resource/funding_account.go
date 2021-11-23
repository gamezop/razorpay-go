package resource

type FundingAccountType string

const (
	FA_UPI  FundingAccountType = "vpa"
	FA_BANK FundingAccountType = "bank_account"
)

type RequestFundingAccount struct {
	AccountType string `json:"account_type" validate:"required"`
	ContactID   string `json:"contact_id" validate:"required"`
}

// UPI
type RequestFundingAccountUPIDetails struct {
	Address string `json:"address" validate:"required"`
}

type RequestFundingAccountUPI struct {
	RequestFundingAccount
	Vpa RequestFundingAccountUPIDetails `json:"vpa" validate:"required"`
}

type fundingAccount struct {
	ID          string  `json:"id"`
	Entity      string  `json:"entity"`
	ContactID   string  `json:"contact_id"`
	AccountType string  `json:"account_type"`
	Active      bool    `json:"active"`
	BatchID     *string `json:"batch_id"`
	CreatedAt   int     `json:"created_at"`
}

type FundingAccountUPI struct {
	fundingAccount
	Vpa struct {
		Username string `json:"username"`
		Handle   string `json:"handle"`
		Address  string `json:"address"`
	} `json:"vpa"`
}

func (f *FundingAccountUPI) GetContactId() string {
	return f.ContactID
}

func (f *FundingAccountUPI) GetAccountType() string {
	return "vpa"
}

func (f *FundingAccountUPI) GetIdentifierForAccount() string {
	return f.Vpa.Address
}

func (f *FundingAccountUPI) GetFundingAccountId() string {
	return f.ID
}

// NEFT
