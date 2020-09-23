package telegram

import (
	"errors"
	"fmt"
)

type GetMeResponse struct {
	Response
	Result *User `json:"result"`
}

func (t *Telegram) GetMe() (*User, error) {
	var r GetMeResponse

	if err := t.makeRequest("getMe", nil, &r); err != nil {
		return nil, err
	}

	if r.OK == false {
		return nil, errors.New(fmt.Sprintf("telegram: %s", r.Description))
	}

	return r.Result, nil
}
