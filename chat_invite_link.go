// Bot API 5.7

package telegram

// ChatInviteLink is an invitation link for a chat.
type ChatInviteLink struct {
	// InviteLink.
	// If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// Creator of the link.
	Creator *User `json:"creator"`

	// CreatesJoinRequest is True, if users joining the chat via the link need to be approved by chat administrators.
	CreatesJoinRequest bool `json:"creates_join_request"`

	// IsPrimary is True, if the link is primary.
	IsPrimary bool `json:"is_primary"`

	// IsRevoked is True, if the link is revoked.
	IsRevoked bool `json:"is_revoked"`

	// Name is an invitation link name.
	//
	// Optional.
	Name string `json:"name,omitempty"`

	// ExpireDate is a point in time (Unix timestamp) when the link will expire or has been expired.
	//
	// Optional.
	ExpireDate int `json:"expire_date,omitempty"`

	// MemberLimit is a maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999.
	//
	// Optional.
	MemberLimit int `json:"member_limit,omitempty"`

	// PendingJoinRequestCount is a number of pending join requests created using this link.
	//
	// Optional.
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}

func (l *ChatInviteLink) Validate() bool {
	if l.MemberLimit < 0 || l.MemberLimit > 99999 {
		return false
	}

	return true
}
