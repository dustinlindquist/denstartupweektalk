package slack

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type msgInput struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// Client provides an http connection for communicating with the Slack API.``
type Client struct {
	host       string
	httpClient *http.Client
}

type dataResp struct {
	Datas []Data `json:"data"`
}

type Data struct {
	Temp   float32 `json:"temp"`
	Weater struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// NewClient accepts an httpClient to faciliate making http requests to Slack.
// Client does not attempt to evaluate the response body, leaving that to the caller.
func NewClient(client *http.Client, host string) *Client {
	return &Client{host: host, httpClient: client}
}

// PostMessage calls the chat.postMessage Slack endpoint to post a message into a channel
func (c *Client) SendMessage(ctx context.Context, channel, msg string) error {
	var err = godotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}
	var token = os.Getenv("SlackToken")
	var params = fmt.Sprintf("?token=%s&channel=%s&text=%s", token, channel, url.QueryEscape(msg))
	var full = fmt.Sprintf("%s/api/chat.postMessage%s", c.host, params)
	req, err := http.NewRequest("POST", full, nil)
	if err != nil {
		return errors.New("error creating request.")
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	_, err = c.httpClient.Do(req)
	if err != nil {
		return errors.New("error doing request.")
	}
	fmt.Println("Send Slack Message.")
	return nil
}
