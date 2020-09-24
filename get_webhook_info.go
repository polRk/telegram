package telegram

import (
	"fmt"
)

// GetWebhookInfoResponse returns WebhookInfo on success.
type GetWebhookInfoResponse struct {
	Response
	Result *WebhookInfo `json:"result"`
}

// GetWebhookInfo returns current webhook status.
// On success, returns a WebhookInfo object.
// If the bot is using getUpdates, will return an object with the WebhookInfo.URL field empty.
func (tg *Telegram) GetWebhookInfo() (*WebhookInfo, error) {
	var r GetWebhookInfoResponse

	if err := tg.makeRequest("getWebhookInfo", nil, &r); err != nil {
		return nil, err
	}

	if !r.OK {
		return nil, fmt.Errorf("telegram: %s", r.Description)
	}

	return r.Result, nil
}
