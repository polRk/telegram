package telegram

import (
	"fmt"
)

// GetUpdatesPayload represents a request payload for method Telegram.GetUpdates.
type GetUpdatesPayload struct {
	// Offset represents an identifier of the first update to be returned.
	// Must be greater by one than the highest among the identifiers of previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned.
	// An update is considered confirmed as soon as getUpdates is called with an offset higher than its Update.ID.
	// The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue.
	// All previous updates will forgotten.
	Offset int `json:"offset,omitempty"`

	// Limit represents the number of updates to be retrieved.
	// Values between 1-100 are accepted.
	// Defaults to 100.
	Limit int `json:"limit,omitempty"`

	// Timeout represents timeout in seconds for long polling.
	// Defaults to 0, i.e. usual short polling.
	// Should be positive, short polling should be used for testing purposes only.
	Timeout int `json:"timeout,omitempty"`

	// AllowedUpdates represents the list of AllowedUpdate you want your bot to receive.
	AllowedUpdates []AllowedUpdate `json:"allowed_updates,omitempty"`
}

// Validate returns an error if payload is invalid.
func (p GetUpdatesPayload) Validate() error {
	for _, allowedUpdate := range p.AllowedUpdates {
		if err := allowedUpdate.Validate(); err != nil {
			return err
		}
	}

	if p.Limit != 0 && (p.Limit < 1 || p.Limit > 100) {
		return fmt.Errorf("telegram: Limit should be >= 1 and <= 100")
	}

	if p.Timeout < 0 {
		return fmt.Errorf("telegram: Timeout should be > 0")
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
func (tg *Telegram) GetUpdates(payload GetUpdatesPayload) ([]*Update, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	var r GetUpdatesResponse

	if err := tg.makeRequest("getUpdates", payload, &r); err != nil {
		return nil, err
	}

	if !r.OK {
		return nil, fmt.Errorf("telegram: %s", r.Description)
	}

	return r.Result, nil
}
