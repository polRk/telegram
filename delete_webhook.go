package telegram

// DeleteWebhookResponse returns True on success
type DeleteWebhookResponse struct {
	Response
	Result bool `json:"result"`
}

// DeleteWebhook removes webhook integration
func (tg *Telegram) DeleteWebhook() (bool, error) {
	var r DeleteWebhookResponse

	if err := tg.makeRequest("deleteWebhook", nil, &r); err != nil {
		return false, err
	}

	return r.Result, nil
}
