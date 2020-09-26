package telegram

import "fmt"

const (
	AllowedUpdateMessage            AllowedUpdate = "message"
	AllowedUpdateEditedMessage      AllowedUpdate = "edited_message"
	AllowedUpdateChannelPost        AllowedUpdate = "channel_post"
	AllowedUpdateEditedChannelPost  AllowedUpdate = "edited_channel_post"
	AllowedUpdateInlineQuery        AllowedUpdate = "inline_query"
	AllowedUpdateChosenInlineResult AllowedUpdate = "chosen_inline_result"
	AllowedUpdateCallbackQuery      AllowedUpdate = "callback_query"
	AllowedUpdateShippingQuery      AllowedUpdate = "shipping_query"
	AllowedUpdatePreCheckoutQuery   AllowedUpdate = "pre_checkout_query"
	AllowedUpdatePoll               AllowedUpdate = "poll"
	AllowedUpdatePollAnswer         AllowedUpdate = "poll_answer"
)

// ErrBadAllowedUpdateValue indicates an invalid value.
var ErrBadAllowedUpdateValue = fmt.Errorf("telegram: bad value for AllowedUpdate type")

// AllowedUpdate is an Update type
type AllowedUpdate string

// WebhookInfo represents information about the current status of a webhook.
type WebhookInfo struct {
	URL                  string          `json:"url"`
	HasCustomCertificate bool            `json:"has_custom_certificate,omitempty"`
	PendingUpdateCount   int             `json:"pending_update_count,omitempty"`
	LastErrorDate        int             `json:"last_error_date,omitempty"`
	LastErrorMessage     int             `json:"last_error_message,omitempty"`
	MaxConnections       int             `json:"max_connections,omitempty"`
	AllowedUpdates       []AllowedUpdate `json:"allowed_updates,omitempty"`
}

// Validate returns an error if value is invalid.
func (au AllowedUpdate) Validate() error {
	switch au {
	case AllowedUpdateMessage,
		AllowedUpdateEditedMessage,
		AllowedUpdateChannelPost,
		AllowedUpdateEditedChannelPost,
		AllowedUpdateInlineQuery,
		AllowedUpdateChosenInlineResult,
		AllowedUpdateCallbackQuery,
		AllowedUpdateShippingQuery,
		AllowedUpdatePreCheckoutQuery,
		AllowedUpdatePoll,
		AllowedUpdatePollAnswer:
		return nil
	}

	return ErrBadAllowedUpdateValue
}
