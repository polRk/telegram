package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Telegram struct {
	baseURL    string
	token      string
	client     *http.Client
	bufferSize int
}

type Response struct {
	OK          bool        `json:"ok"`
	Result      interface{} `json:"result,omitempty"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
}

func NewTelegram(token string, bufferSize int) *Telegram {
	return &Telegram{
		baseURL:    fmt.Sprintf("https://api.telegram.org/bot%s", token),
		token:      token,
		client:     &http.Client{},
		bufferSize: bufferSize,
	}
}

func (tg *Telegram) Listen(pattern string) chan *Update {
	ch := make(chan *Update, tg.bufferSize)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var update Update

		if json.NewDecoder(r.Body).Decode(&update) != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ch <- &update
	})

	return ch
}

func (tg Telegram) makeRequest(method string, payload interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", tg.baseURL, method)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := tg.client.Post(url, "application/json", bytes.NewBuffer(body))
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
