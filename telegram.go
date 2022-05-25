package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// APIEndpoint is the endpoint for all API methods.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"

	// FileEndpoint is the endpoint for downloading a file from Telegram server.
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

// httpClient defines the minimal interface needed for a http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	token        string
	apiEndpoint  string
	fileEndpoint string
	httpclient   httpClient
}

// Option defines an option for a Client
type Option func(*Client)

// OptionHTTPClient - provide a custom http client to the client.
func OptionHTTPClient(client httpClient) func(*Client) {
	return func(c *Client) {
		c.httpclient = client
	}
}

// OptionAPIURL set the apiEndpoint for the client.
func OptionAPIURL(u string) func(*Client) {
	return func(c *Client) { c.apiEndpoint = u }
}

// OptionFileURL set the fileEndpoint for the client.
func OptionFileURL(u string) func(*Client) {
	return func(c *Client) { c.fileEndpoint = u }
}

// New builds a telegram bot api client from the provided token and options.
func New(token string, options ...Option) *Client {
	s := &Client{
		token:        token,
		apiEndpoint:  APIEndpoint,
		fileEndpoint: FileEndpoint,
		httpclient:   &http.Client{},
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}

// MakeRequest makes a request to a specific endpoint with our token.
func (c *Client) MakeRequest(ctx context.Context, method string, body interface{}) (*APIResponse, error) {
	endpoint := fmt.Sprintf(c.apiEndpoint, c.token, method)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, buf)
	if err != nil {
		return &APIResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return &apiResp, err
	}

	if !apiResp.Ok {
		return &apiResp, &Error{
			Code:               apiResp.ErrorCode,
			Message:            apiResp.Description,
			ResponseParameters: apiResp.Parameters,
		}
	}

	return &apiResp, nil
}
