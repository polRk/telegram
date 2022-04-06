// Bot API 5.7

package telegram

// ChatJoinRequest is a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat to which the request was sent.
	Chat *Chat `json:"chat"`

	// User that sent the join request.
	User *User `json:"user"`

	// Date the request was sent in Unix time.
	Date int `json:"date"`

	// Bio of the user.
	//
	// Optional.
	Bio string `json:"bio,omitempty"`

	// InviteLink is a chat invite link that was used by the user to send the join request.
	//
	// Optional.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
