package telegram

import "fmt"

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

// Validate returns an error if value is invalid.
func (b InlineKeyboardButton) Validate() error {
	optionalCount := 0

	if len(b.Text) == 0 {
		return fmt.Errorf("telegram[InlineKeyboardButton]: Text is required")
	}

	if len(b.URL) != 0 {
		optionalCount += 1
	}

	if b.LoginURL != nil {
		optionalCount += 1

		if err := b.LoginURL.Validate(); err != nil {
			return err
		}
	}

	if len(b.CallbackData) != 0 {
		optionalCount += 1
	}

	if len(b.SwitchInlineQuery) != 0 {
		optionalCount += 1
	}

	if len(b.SwitchInlineQueryCurrentChat) != 0 {
		optionalCount += 1
	}

	if b.CallbackGame != nil {
		optionalCount += 1
	}

	if b.Pay != false {
		optionalCount += 1
	}

	if optionalCount > 1 {
		return fmt.Errorf("telegram[InlineKeyboardButton]: you must use exactly one of the optional fields")
	}

	return nil
}
