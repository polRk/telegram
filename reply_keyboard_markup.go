// Bot API 5.7

package telegram

// ReplyKeyboardMarkup is a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an Array of KeyboardButton objects.
	Keyboard [][]*KeyboardButton `json:"keyboard"`

	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit.
	// Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	//
	// Optional.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat -
	// the user can press a special button in the input field to see the custom keyboard again.
	// Defaults to false.
	//
	// Optional.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder is the placeholder to be shown in the input field when the keyboard is active.
	// 1-64 characters.
	//
	// Optional.
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective is showing the keyboard to specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}
