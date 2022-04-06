// Bot API 5.7

package telegram

// VoiceChatScheduled is a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	//StartDate is a point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator.
	StartDate int `json:"start_date"`
}
