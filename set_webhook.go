package telegram

type SetWebhookPayload struct {
	URL            string          `json:"url"`
	Certificate    InputFile       `json:"certificate,omitempty"`
	MaxConnections int             `json:"max_connections,omitempty"`
	AllowedUpdates []AllowedUpdate `json:"allowed_updates,omitempty"`
}

type SetWebhookResponse struct {
	Response
	Result bool `json:"result"`
}

func (tg *Telegram) SetWebhook(payload SetWebhookPayload) (bool, error) {
	var r SetWebhookResponse

	if err := tg.makeRequest("setWebhook", payload, &r); err != nil {
		return false, err
	}

	return r.Result, nil
}
