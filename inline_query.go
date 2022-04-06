// Bot API 5.7

package telegram

// InlineQuery is an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// ID is a unique identifier for this query.
	ID string `json:"id"`

	// From is a Sender.
	From *User `json:"from"`

	// Query is a text of the query (up to 256 characters).
	Query string `json:"query"`

	// Offset of the results to be returned, can be controlled by the bot.
	Offset string `json:"offset"`

	// ChatType is a type of the chat, from which the inline query was sent.
	// Can be either ChatTypeSender for a private chat with the inline query sender,
	// ChatTypePrivate, ChatTypeGroup, ChatTypeSuperGroup, or ChatTypeChannel.
	// The chat type should be always known for requests sent from official clients
	// and most third-party clients, unless the request was sent from a secret chat.
	//
	// Optional.
	ChatType ChatType `json:"chat_type,omitempty"`

	// Location is a sender location, only for bots that request user location.
	//
	// Optional.
	Location *Location `json:"location,omitempty"`
}

func (i *InlineQuery) Validate() bool {
	if len(i.Query) > 256 {
		return false
	}

	return true
}
