package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Telegram struct {
	baseURL string
	token   string
	client  *http.Client
}

type Response struct {
	OK          bool        `json:"ok"`
	Result      interface{} `json:"result,omitempty"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
}

func NewTelegram(token string) *Telegram {
	return &Telegram{
		baseURL: fmt.Sprintf("https://api.telegram.org/bot%s", token),
		token:   token,
		client:  &http.Client{},
	}
}

func (t Telegram) makeRequest(method string, payload interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", t.baseURL, method)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := t.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
