package telegram

type AllowedUpdate string

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

type WebhookInfo struct {
	Url                  string          `json:"url"`
	HasCustomCertificate bool            `json:"has_custom_certificate,omitempty"`
	PendingUpdateCount   int             `json:"pending_update_count,omitempty"`
	LastErrorDate        int             `json:"last_error_date,omitempty"`
	LastErrorMessage     int             `json:"last_error_message,omitempty"`
	MaxConnections       int             `json:"max_connections,omitempty"`
	AllowedUpdates       []AllowedUpdate `json:"allowed_updates,omitempty"`
}
