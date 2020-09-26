package telegram

import (
	"fmt"
)

// GetUpdatesPayload represents a request payload for method Telegram.GetUpdates.
type GetUpdatesPayload struct {
	Offset         int             `json:"offset,omitempty"`
	Limit          int             `json:"limit,omitempty"`
	Timeout        int             `json:"timeout,omitempty"`
	AllowedUpdates []AllowedUpdate `json:"allowed_updates,omitempty"`
}

// Validate returns an error if payload is invalid.
func (p GetUpdatesPayload) Validate() error {
	if p.Limit != 0 && (p.Limit < 1 || p.Limit > 100) {
		return fmt.Errorf("telegram[GetUpdatesPayload]: Limit should be >= 1 and <= 100")
	}

	if p.Timeout < 0 {
		return fmt.Errorf("telegram[GetUpdatesPayload]: Timeout should be > 0")
	}

	for _, allowedUpdate := range p.AllowedUpdates {
		if err := allowedUpdate.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetUpdatesResponse represents a request response.
// Returns an array of *Update in Result fiend on success.
type GetUpdatesResponse struct {
	Response
	Result []*Update `json:"result"`
}

// GetUpdates returns incoming updates.
// An Array of *Update objects is returned.
func (tg *Telegram) GetUpdates(payload GetUpdatesPayload) ([]*Update, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	var r GetUpdatesResponse

	if err := tg.makeRequest("getUpdates", payload, &r); err != nil {
		return nil, err
	}

	if !r.OK {
		return nil, fmt.Errorf("telegram[GetUpdates]: %s", r.Description)
	}

	return r.Result, nil
}
