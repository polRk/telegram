// Bot API 5.7

package api

// ShippingAddress is a shipping address.
type ShippingAddress struct {
	// CountryCode is a	ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	// State, if applicable.
	State string `json:"state"`

	// City.
	City string `json:"city"`

	// StreetLine1 is a first line for the address.
	StreetLine1 string `json:"street_line1"`

	// StreetLine2 is a second line for the address.
	StreetLine2 string `json:"street_line2"`

	// PostCode is an address post code.
	PostCode string `json:"post_code"`
}
