// Bot API 5.7

package telegram

// Dice is an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based.
	Emoji string `json:"emoji"`

	// Value of the dice,
	// 1-6 for “🎲”,
	// “🎯” and “🎳” base emoji,
	// 1-5 for “🏀” and “⚽” base emoji,
	// 1-64 for “🎰” base emoji.
	Value int `json:"value"`
}
