// Bot API 5.7

package api

// ForceReply showing a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces
// without having to sacrifice privacy mode.
type ForceReply struct {
	// ForceReply shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'.
	ForceReply bool `json:"force_reply"`

	// InputFieldPlaceholder is the placeholder to be shown in the input field when the reply is active; 1-64 characters.
	//
	// Optional
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective is force reply from specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}