// Bot API 5.7

package telegram

// Game is a Telegram game.
type Game struct {
	// Title of the game.
	Title string `json:"title"`

	// Description of the game.
	Description string `json:"description"`

	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo"`

	// Text is a brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game
	// when the bot calls setGameScore, or manually edited using editMessageText.
	// 0-4096 characters.
	//
	// Optional.
	Text string `json:"text,omitempty"`

	// TextEntities is a special entities that appear in text, such as usernames, URLs, bot commands, etc.
	//
	// Optional.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`

	// Animation that will be displayed in the game message in chats.
	// Upload via BotFather.
	//
	// Optional.
	Animation *Animation `json:"animation,omitempty"`
}

func (g *Game) Validate() bool {
	if len(g.Text) > 4096 {
		return false
	}

	return true
}
