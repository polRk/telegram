// Bot API 5.7

package telegram

// MessageAutoDeleteTimerChanged is a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// MessageAutoDeleteTime is a new auto-delete time in seconds for messages in the chat.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}
