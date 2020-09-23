package telegram

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	VCard       string `json:"v_card,omitempty"`
}
