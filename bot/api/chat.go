// Bot API 5.7

package api

type ChatType string

const (
	ChatTypeSender     ChatType = "sender"
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// Chat represents a chat.
type Chat struct {
	// ID is a unique identifier for this chat.
	ID int64 `json:"id"`

	// Type of chat, can be either ChatTypePrivate, ChatTypeGroup, ChatTypeSuperGroup or ChatTypeChannel.
	Type ChatType `json:"type"`

	// Title for ChatTypeSuperGroup, ChatTypeChannel and ChatTypeGroup chats.
	//
	// Optional.
	Title string `json:"title,omitempty"`

	// Username for ChatTypePrivate, ChatTypeSuperGroup and ChatTypeChannel if available.
	//
	// Optional.
	Username string `json:"username,omitempty"`

	// FirstName of the other party in a ChatTypePrivate.
	//
	// Optional.
	FirstName string `json:"first_name,omitempty"`

	// LastName of the other party in a ChatTypePrivate.
	//
	// Optional.
	LastName string `json:"last_name,omitempty"`

	// Photo is a chat photo.
	// Returned only in api.GetChat().
	//
	// Optional.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Bio of the other party in a private chat.
	// Returned only in api.GetChat().
	//
	// Optional.
	Bio string `json:"bio,omitempty"`

	// HasPrivateForwards is True, if privacy settings of the other party in the private chat
	// allows to use tg://user?id=<user_id> links only in chats with the user.
	// Returned only in api.GetChat().
	//
	// Optional.
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`

	// Description for ChatTypeGroup, ChatTypeSuperGroup and ChatTypeChannel chats.
	// Returned only in api.GetChat().
	//
	// Optional.
	Description string `json:"description,omitempty"`

	// InviteLink is a primary invite link, for ChatTypeGroup, ChatTypeSuperGroup and ChatTypeChannel chats.
	// Returned only in api.GetChat().
	//
	// Optional.
	InviteLink string `json:"invite_link,omitempty"`

	// PinnedMessage is the most recent pinned message.
	// Returned only in api.GetChat().
	//
	// Optional.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Permissions are default chat member permissions, for ChatTypeGroup and ChatTypeSuperGroup.
	// Returned only in api.GetChat().
	//
	// Optional.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// SlowModeDelay is the minimum allowed delay between consecutive messages sent by each unprivileged user.
	// For ChatTypeSuperGroup.
	// Returned only in api.GetChat().
	//
	// Optional.
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// MessageAutoDeleteTime is the time in seconds
	// after which all messages sent to the chat will be automatically deleted.
	// Returned only in api.GetChat().
	//
	// Optional.
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`

	// HasProtectedContent is True, if messages from the chat can't be forwarded to other chats.
	// Returned only in api.GetChat().
	//
	// Optional.
	HasProtectedContent int `json:"has_protected_content,omitempty"`

	// StickerSetName ia a name of group sticker set.
	// Returned only in api.GetChat().
	//
	// Optional.
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet is True, if the bot can change the group sticker set.
	// Returned only in api.GetChat().
	//
	// Optional.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`

	// LinkedChatID is a unique identifier for the linked chat,
	// i.e. the discussion group identifier for a channel and vice versa.
	// For ChatTypeSuperGroup and ChatTypeChannel.
	// Returned only in api.GetChat().
	//
	// Optional.
	LinkedChatID int64 `json:"linked_chat_id,omitempty"`

	// Location is the location to which the ChatTypeSuperGroup is connected.
	// For ChatTypeSuperGroup.
	// Returned only in api.GetChat().
	//
	// Optional.
	Location *ChatLocation `json:"location,omitempty"`
}
