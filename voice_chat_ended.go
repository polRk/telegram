// Bot API 5.7

package telegram

// VoiceChatEnded is a service message about a voice chat ended in the chat.
type VoiceChatEnded struct {
	// Duration is a voice chat duration in seconds.
	Duration int `json:"duration"`
}
