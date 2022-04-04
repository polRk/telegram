// Bot API 5.7

package api

// OrderInfo is an information about an order.
type OrderInfo struct {
	// Name is a user's name.
	//
	// Optional.
	Name string `json:"name,omitempty"`

	// PhoneNumber is a user's phone number.
	//
	// Optional.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email is a user's email.
	//
	// Optional.
	Email string `json:"email,omitempty"`

	// ShippingAddress is a user's shipping address.
	//
	// Optional.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}
