// Bot API 5.7

package api

// VoiceChatParticipantsInvited is a service message about new members invited to a voice chat.
type VoiceChatParticipantsInvited struct {
	//Users is an array of new members that were invited to the voice chat.
	//
	// Optional.
	Users []*User `json:"users"`
}
