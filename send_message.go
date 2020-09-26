package telegram

import (
	"fmt"
)

// SendMessagePayload represents a request payload for method Telegram.SendMessage.
type SendMessagePayload struct {
	ChatID                int                           `json:"chat_id"`
	Text                  string                        `json:"text"`
	ParseMode             ParseMode                     `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool                          `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool                          `json:"disable_notification,omitempty"`
	ReplyToMessageID      int                           `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{ Validate() error } `json:"reply_markup,omitempty"`
}

// Validate returns an error if payload is invalid.
func (p SendMessagePayload) Validate() error {
	if p.ChatID == 0 {
		return fmt.Errorf("telegram[SendMessagePayload]: ChatID is required")
	}

	if len(p.Text) == 0 {
		return fmt.Errorf("telegram[SendMessagePayload]: Text is required")
	}

	if len(p.Text) > 4096 {
		return fmt.Errorf("telegram[SendMessagePayload]: Text length should be less than 4096")
	}

	if len(p.ParseMode) != 0 {
		if err := p.ParseMode.Validate(); err != nil {
			return err
		}
	}

	if p.ReplyMarkup != nil {
		switch p.ReplyMarkup.(type) {
		case *InlineKeyboardMarkup, *ReplyKeyboardMarkup, *ReplyKeyboardRemove, *ForceReply:
			if err := p.ReplyMarkup.Validate(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("telegram[SendMessagePayload]: Text is required")
		}
	}

	return nil
}

// SendMessageResponse represents a request response.
// Returns an array of *Message in Result fiend on success.
type SendMessageResponse struct {
	Response
	Result *Message `json:"result"`
}

// SendMessage sends text messages.
// On success, the sent *Message is returned.
func (tg *Telegram) SendMessage(payload SendMessagePayload) (*Message, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	var r SendMessageResponse

	if err := tg.makeRequest("sendMessage", payload, &r); err != nil {
		return nil, err
	}

	if !r.OK {
		return nil, fmt.Errorf("telegram[SendMessage]: %s", r.Description)
	}

	return r.Result, nil
}
