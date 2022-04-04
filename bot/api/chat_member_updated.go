// Bot API 5.7

package api

// ChatMemberUpdated is a changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat the user belongs to.
	Chat *Chat `json:"chat"`

	// From is a performer of the action, which resulted in the change.
	From *User `json:"from"`

	// Date the change was done in Unix time
	Date int `json:"date"`

	// OldChatMember is previous information about the chat member.
	OldChatMember *ChatMember `json:"old_chat_member"`

	// NewChatMember is new information about the chat member.
	NewChatMember *ChatMember `json:"new_chat_member"`

	// InviteLink is a chat invite link, which was used by the user to join the chat.
	// For joining by invite link events only.
	//
	// Optional.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
