package telegram

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type Chat struct {
	ID               int              `json:"id"`
	Type             ChatType         `json:"type"`
	Title            string           `json:"title,omitempty"`
	Username         string           `json:"username,omitempty"`
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	ChatPhoto        *ChatPhoto       `json:"chat_photo,omitempty"`          // Returned only in Telegram.GetChat().
	Description      string           `json:"description,omitempty"`         // Returned only in Telegram.GetChat().
	InviteLink       string           `json:"invite_link,omitempty"`         // Returned only in Telegram.GetChat().
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`      // Returned only in Telegram.GetChat().
	Permissions      *ChatPermissions `json:"permissions,omitempty"`         // Returned only in Telegram.GetChat().
	SlowModeDelay    int              `json:"slow_mode_delay,omitempty"`     // Returned only in Telegram.GetChat().
	StickerSetName   string           `json:"sticker_set_name,omitempty"`    // Returned only in Telegram.GetChat().
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"` // Returned only in Telegram.GetChat().
}
