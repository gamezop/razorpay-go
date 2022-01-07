package resource

type RequestCreateContact struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email"`
	Contact     string `json:"contact"`
	Type        string `json:"type"`
	ReferenceID string `json:"reference_id"`
	// Notes       map[string]string `json:"notes"`
}

// disabling notes for now, response in dev doesn't not match docs
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
