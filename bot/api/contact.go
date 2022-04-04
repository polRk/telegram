package api

// Contact is a phone contact.
type Contact struct {
	// PhoneNumber is a Contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// FirstName is a Contact's first name.
	FirstName string `json:"first_name"`

	// LastName is a Contact's last name.
	//
	// Optional.
	LastName string `json:"last_name,omitempty"`

	// UserID is a Contact's user identifier in Telegram.
	//
	// Optional.
	UserID int64 `json:"user_id,omitempty"`

	// VCard is an additional data about the contact in the form of a vCard.
	//
	// Optional.
	VCard string `json:"v_card,omitempty"`
}
