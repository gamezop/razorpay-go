package resource

type PAYOUT_MODE string

const (
	PAYOUT_MODE_UPI  PAYOUT_MODE = "UPI"
	PAYOUT_MODE_IMPS PAYOUT_MODE = "IMPS"
)

type PURPOSE string

const (
	PURPOSE_PAYOUT PAYOUT_MODE = "payout"
)

type CURRENCY string

const (
	CURRENCY_INR CURRENCY = "INR"
)

type STATUS string

const (
	STATUS_PROCESSED  STATUS = "processed"
	STATUS_PROCESSING STATUS = "processing"
	STATUS_REVERSED   STATUS = "reversed"
)

// for UPI

type RequestPayout struct {
	AccountNumber     string `json:"account_number" validate:"required"`
	FundAccountID     string `json:"fund_account_id" validate:"required"`
	Amount            int    `json:"amount" validate:"required"`
	Currency          string `json:"currency" validate:"required"`
	Mode              string `json:"mode" validate:"required"`
	Purpose           string `json:"purpose" validate:"required"`
	QueueIfLowBalance bool   `json:"queue_if_low_balance"`
	ReferenceID       string `json:"reference_id"`
	Narration         string `json:"narration"`
	// Notes             map[string]string
}

type Payout struct {
	ID            string `json:"id"`
	Entity        string `json:"entity"`
	FundAccountID string `json:"fund_account_id"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	// Notes         map[string]string
	Fees          int     `json:"fees"`
	Tax           int     `json:"tax"`
	Status        string  `json:"status"`
	Utr           *string `json:"utr"` // bank unique transaction identifier
	Mode          string  `json:"mode"`
	Purpose       string  `json:"purpose"`
	ReferenceID   string  `json:"reference_id"`
	Narration     string  `json:"narration"`
	BatchID       *string `json:"batch_id"`
	FailureReason *string `json:"failure_reason"`
	CreatedAt     int     `json:"created_at"`
}
