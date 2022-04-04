// Bot API 5.7

package api

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard is an array of button rows, each represented by an Array of InlineKeyboardButton objects.
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// Validate returns an error if value is invalid.
func (k *InlineKeyboardMarkup) Validate() error {
	for i := range k.InlineKeyboard {
		for j := range k.InlineKeyboard[i] {
			if err := k.InlineKeyboard[i][j].Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
