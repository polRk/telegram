// Bot API 5.7

package api

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Currency is a three-letter ISO 4217 currency code.
	// See: https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`

	// TotalAmount is a total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload is a bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID is an identifier of the shipping option chosen by the user.
	//
	// Optional.
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo is an order info provided by the user.
	//
	// Optional.
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// TelegramPaymentChargeID is a Telegram payment identifier.
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// ProviderPaymentChargeID is a Provider payment identifier.
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}
