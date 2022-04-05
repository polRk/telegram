// Bot API 5.7

package api

// ReplyKeyboardRemove remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button.
type ReplyKeyboardRemove struct {
	// RemoveKeyboard requests clients to remove the custom keyboard.
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Selective is removing the keyboard for specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	//
	// Optional.
	Selective bool `json:"selective,omitempty"`
}