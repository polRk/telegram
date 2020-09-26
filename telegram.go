package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Telegram represents a telegram bot api client.
type Telegram struct {
	baseURL         string
	token           string
	client          *http.Client
	bufferSize      int
	shutdownChannel chan interface{}
}

// Response represents a response from the Telegram API.
type Response struct {
	OK          bool        `json:"ok"`
	Result      interface{} `json:"result,omitempty"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
}

// NewTelegram returns a new Telegram.
func NewTelegram(token string, bufferSize int) *Telegram {
	return &Telegram{
		baseURL:         fmt.Sprintf("https://api.telegram.org/bot%s", token),
		token:           token,
		client:          &http.Client{},
		bufferSize:      bufferSize,
		shutdownChannel: make(chan interface{}),
	}
}

func (tg *Telegram) StartPolling(uu []AllowedUpdate) chan *Update {
	ch := make(chan *Update, tg.bufferSize)

	go func() {
		var lastUpdateID int

		for {
			select {
			case <-tg.shutdownChannel:
				close(ch)
				return
			default:
			}

			payload := GetUpdatesPayload{Offset: lastUpdateID, AllowedUpdates: uu}
			updates, err := tg.GetUpdates(payload)
			if err != nil {
				println("Err", err.Error())
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.UpdateID >= lastUpdateID {
					lastUpdateID = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func (tg *Telegram) StopPolling() {
	tg.shutdownChannel <- nil
}

// ListenAndServe listens and push incoming updates to *Update channel.
func (tg *Telegram) ListenAndServe(addr string) (chan *Update, error) {
	ch := tg.HandleUpdates(fmt.Sprintf("/%s", tg.token))

	if err := http.ListenAndServe(addr, nil); err != nil {
		return nil, err
	}

	return ch, nil
}

// HandleUpdates returns the *Update channel.
// HandleUpdates adds handler to http.DefaultServeMux.
func (tg *Telegram) HandleUpdates(pattern string) chan *Update {
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

// makeRequest makes a request to a specific endpoint with token.
func (tg *Telegram) makeRequest(method string, payload interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", tg.baseURL, method)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
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
