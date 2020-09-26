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

// Validate returns an error if value is invalid.
func (value AllowedUpdate) Validate() error {
	switch value {
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
