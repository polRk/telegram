package api

// Invoice contains basic information about an invoice.
type Invoice struct {
	// Title is a product name.
	Title string `json:"title"`

	// Description is a product description.
	Description string `json:"description"`

	// StartParameter is a unique bot deep-linking parameter that can be used to generate this invoice.
	StartParameter string `json:"start_parameter"`

	// Currency is a three-letter ISO 4217 currency code.
	// See: https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`

	// TotalAmount is a total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int64 `json:"total_amount"`
}
