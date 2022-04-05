// Bot API 5.7

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

type TelegramBotAPI struct {
	token string
}

type TelegramResponse[R any] struct {
	OK          bool   `json:"ok"`
	Result      R      `json:"result"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters  struct {
		MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
		RetryAfter      int   `json:"retry_after,omitempty"`
	} `json:"parameters,omitempty"`
}

func NewClient(token string) *TelegramBotAPI {
	return &TelegramBotAPI{
		token: token,
	}
}

func (t *TelegramBotAPI) requestURL(method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", t.token, method)
}

// GetMe returns basic information about the bot in form of a User object.
func (t *TelegramBotAPI) GetMe() (*User, error) {
	url := t.requestURL("getMe")

	return get[User](url)
}

// SendMessage sending text messages.
// On success, the sent Message is returned.
func (t *TelegramBotAPI) SendMessage(body *Message) (*Message, error) {
	url := t.requestURL("sendMessage")

	return post[Message, Message](url, body)
}

type AnswerCallbackQueryBody struct {
	// CallbackQueryId is a unique identifier for the query to be answered.
	CallbackQueryId string `json:"callback_query_id"`

	// Text of the notification.
	// If not specified, nothing will be shown to the user, 0-200 characters.
	//
	// Optional.
	Text string `json:"text,omitempty"`

	// ShowAlert if True, an alert will be shown by the client instead of a notification at the top of the chat screen.
	// Defaults to false.
	//
	// Optional.
	ShowAlert bool `json:"show_alert,omitempty"`

	// URL that will be opened by the user's client.
	// If you have created a Game and accepted the conditions via @Botfather,
	// specify the URL that opens your game — note that this will only work if the query comes from a callback_game button.
	// Otherwise, you may use links like t.me/your_bot?start=XXXX.
	//
	// Optional.
	URL string `json:"url,omitempty"`

	// CacheTime is the maximum amount of time in seconds that the result of the callback query may be cached client-side.
	// Telegram apps will support caching starting in version 3.14.
	// Defaults to 0.
	//
	// Optional.
	CacheTime int `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery answer to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// On success, True is returned.
func (t *TelegramBotAPI) AnswerCallbackQuery(body *AnswerCallbackQueryBody) (*Message, error) {
	url := t.requestURL("answerCallbackQuery")

	return post[AnswerCallbackQueryBody, Message](url, body)
}

func get[RESP any](url string) (*RESP, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return decode[RESP](resp.Body)
}

func post[REQ, RESP any](url string, body *REQ) (*RESP, error) {
	var requestBody []byte

	if body != nil {
		var err error

		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return decode[RESP](resp.Body)
}

func decode[RESP any](body io.ReadCloser) (*RESP, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var data TelegramResponse[RESP]
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	if data.OK {
		return &data.Result, nil
	}

	return nil, errors.New(data.Description)
}
