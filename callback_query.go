// Bot API 5.7

package telegram

// CallbackQuery is an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	// ID is a unique identifier for this query.
	ID string `json:"id"`

	// From is a Sender.
	From *User `json:"from"`

	// Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old.
	//
	// Optional.
	Message *Message `json:"message,omitempty"`

	// InlineMessageID is an identifier of the message sent via the bot in inline mode, that originated the query.
	//
	// Optional.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ChatInstance a global identifier,
	// uniquely corresponding to the chat to which the message with the callback button was sent.
	// Useful for high scores in games.
	//
	// Optional.
	ChatInstance string `json:"chat_instance"`

	// Data associated with the callback button.
	// Be aware that a bad client can send arbitrary data in this field.
	//
	// Optional.
	Data string `json:"data,omitempty"`

	// GameShortName is a short name of a Game to be returned, serves as the unique identifier for the game.
	//
	// Optional.
	GameShortName string `json:"game_short_name,omitempty"`
}
