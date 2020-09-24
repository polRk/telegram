package telegram

import (
	"fmt"
)

type ParseMode string

const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML     ParseMode = "HTML"
)

type SendMessagePayload struct {
	ChatID                int         `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type SendMessageResponse struct {
	Response
	Result *Message `json:"result"`
}

func (tg *Telegram) SendMessage(payload SendMessagePayload) (*Message, error) {
	var r SendMessageResponse

	if err := tg.makeRequest("sendMessage", payload, &r); err != nil {
		return nil, err
	}

	if !r.OK {
		return nil, fmt.Errorf("telegram: %s", r.Description)
	}

	return r.Result, nil
}
