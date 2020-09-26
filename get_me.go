package telegram

import (
	"fmt"
)

// GetMeResponse represents a request response.
// Returns *User in Result fiend on success.
type GetMeResponse struct {
	Response
	Result *User `json:"result"`
}

// GetMe returns basic information about the bot in form of a User object.
func (tg *Telegram) GetMe() (*User, error) {
	var r GetMeResponse

	if err := tg.makeRequest("getMe", nil, &r); err != nil {
		return nil, err
	}

	if !r.OK {
		return nil, fmt.Errorf("telegram[GetMe]: %s", r.Description)
	}

	return r.Result, nil
}
