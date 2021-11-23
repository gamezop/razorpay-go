package resource

type RequestCreateContact struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email"`
	Contact     string `json:"contact"`
	Type        string `json:"type"`
	ReferenceID string `json:"reference_id"`
	// Notes       map[string]string `json:"notes"`
}

// disabling notes for now, response in dev doesnt not match docs
// // {"level":"trace","api":"rzHttpClient.Do.ReadResponse","responseBody":"{\"id\":\"cont_IOyXrG1t6TXFwD\",\"entity\":\"contact\",\"name\":\"Prithvihv\",\"contact\":\"+919902508248\",\"email\":\"phv@gamezop.co\",\"type\":\"customer\",\"reference_id\":\"gzpCode\",\"batch_id\":null,\"active\":true,\"notes\":[],\"created_at\":1637655340}","time":"2021-11-23T13:49:25+05:30"}
// expecting notes to be an map
type Contact struct {
	ID          string  `json:"id"`
	Entity      string  `json:"entity"`
	Name        string  `json:"name"`
	Contact     string  `json:"contact"`
	Email       string  `json:"email"`
	Type        string  `json:"type"`
	ReferenceID string  `json:"reference_id"`
	BatchID     *string `json:"batch_id"`
	Active      bool    `json:"active"`
	// Notes       map[string]string `json:"notes"`
	CreatedAt int `json:"created_at"`
}
