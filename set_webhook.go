package telegram

import "fmt"

// SetWebhookPayload represents a request payload for method Telegram.SetWebhook.
type SetWebhookPayload struct {
	URL            string          `json:"url"`
	Certificate    InputFile       `json:"certificate,omitempty"`
	MaxConnections int             `json:"max_connections,omitempty"`
	AllowedUpdates []AllowedUpdate `json:"allowed_updates,omitempty"`
}

// Validate returns an error if payload is invalid.
func (p SetWebhookPayload) Validate() error {
	if len(p.URL) == 0 {
		return fmt.Errorf("telegram[SetWebhookPayload]: URL is required")
	}

	if p.MaxConnections != 0 && (p.MaxConnections < 1 || p.MaxConnections > 100) {
		return fmt.Errorf("telegram[SetWebhookPayload]: MaxConnections should be >= 1 and <= 100")
	}

	for _, allowedUpdate := range p.AllowedUpdates {
		if err := allowedUpdate.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// SetWebhookResponse represents a request response.
//  Returns True on success.
type SetWebhookResponse struct {
	Response
	Result bool `json:"result"`
}

// SetWebhook specifies a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts.
func (tg *Telegram) SetWebhook(payload SetWebhookPayload) (bool, error) {
	if err := payload.Validate(); err != nil {
		return false, err
	}

	var r SetWebhookResponse

	if err := tg.makeRequest("setWebhook", payload, &r); err != nil {
		return false, err
	}

	return r.Result, nil
}
