package api

type MessageEntityType string

const (
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashTag       MessageEntityType = "hashtag"
	MessageEntityTypeCashTag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeSpoiler       MessageEntityType = "spoiler"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
)

// MessageEntity is one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type is a type of the entity.
	// Currently, can be:
	// MessageEntityTypeMention (@username),
	// MessageEntityTypeHashTag (#hashtag),
	// MessageEntityTypeCashTag ($USD),
	// MessageEntityTypeBotCommand (/start@jobs_bot),
	// MessageEntityTypeURL (https://telegram.org),
	// MessageEntityTypeEmail (do-not-reply@telegram.org),
	// MessageEntityTypePhoneNumber (+1-212-555-0123),
	// MessageEntityTypeBold (bold text),
	// MessageEntityTypeItalic (italic text),
	// MessageEntityTypeUnderline (underlined text),
	// MessageEntityTypeStrikethrough (strikethrough text),
	// MessageEntityTypeSpoiler (spoiler message),
	// MessageEntityTypeCode (monowidth string),
	// MessageEntityTypePre (monowidth block),
	// MessageEntityTypeTextLink (for clickable text URLs),
	// MessageEntityTypeTextMention (for users without usernames).
	Type MessageEntityType `json:"type"`

	// Offset is an offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`

	// Length is a length of the entity in UTF-16 code units.
	Length int `json:"length"`

	// URL is an url that will be opened after user taps on the text.
	// For MessageEntityTypeTextLink only.
	//
	// Optional.
	URL string `json:"url,omitempty"`

	// User is the mentioned user.
	// For MessageEntityTypeTextMention only.
	//
	// Optional.
	User *User `json:"user,omitempty"`

	// Language is the programming language of the entity text.
	// For MessageEntityTypePre only.
	//
	// Optional.
	Language string `json:"language,omitempty"`
}
