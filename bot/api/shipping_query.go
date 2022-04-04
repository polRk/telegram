// Bot API 5.7

package api

// ShippingQuery is an information about an incoming shipping query.
type ShippingQuery struct {
	// ID is a unique query identifier.
	ID string `json:"id"`

	// From is a User who sent the query.
	From *User `json:"from"`

	// InvoicePayload is a Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingAddress is the User specified shipping address.
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}
