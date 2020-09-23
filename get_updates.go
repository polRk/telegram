package telegram

import (
	"errors"
	"fmt"
)

type GetUpdatesPayload struct {
	Offset         int             `json:"offset,omitempty"`
	Limit          int             `json:"limit,omitempty"`
	Timeout        int             `json:"timeout,omitempty"`
	AllowedUpdates []AllowedUpdate `json:"allowed_updates,omitempty"`
}

type GetUpdatesResponse struct {
	Response
	Result []*Update `json:"result"`
}

func (tg *Telegram) GetUpdates(payload GetUpdatesPayload) ([]*Update, error) {
	var r GetUpdatesResponse

	if err := tg.makeRequest("getUpdates", payload, &r); err != nil {
		return nil, err
	}

	if r.OK == false {
		return nil, errors.New(fmt.Sprintf("telegram: %s", r.Description))
	}

	return r.Result, nil
}
