package resource

type IFSCInfo struct {
	Bank     string `json:"BANK"`
	Branch   string `json:"BRANCH"`
	Centre   string `json:"CENTRE"`
	District string `json:"DISTRICT"`
	State    string `json:"STATE"`
	Address  string `json:"ADDRESS"`
	Contact  string `json:"CONTACT"`
	City     string `json:"CITY"`
	IFSC     string `json:"IFSC"`
	UPI      bool   `json:"UPI"`
	RTGS     bool   `json:"RTGS"`
	MICR     string `json:"MICR"`
	NEFT     bool   `json:"NEFT"`
	SWIFT    string `json:"SWIFT"`
	IMPS     bool   `json:"IMPS"`
	BankCode string `json:"BANKCODE"`
}
