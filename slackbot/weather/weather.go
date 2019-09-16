package weather

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Client provides an http connection for communicating with the Slack API.``
type Client struct {
	host       string
	httpClient *http.Client
}

type baseResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type dataResp struct {
	Datas []Data `json:"data"`
}

type Weather struct {
	Description string `json:"description"`
}

type Data struct {
	Temp    float32 `json:"temp"`
	Weather Weather `json:"weather"`
}

// NewClient accepts an httpClient to faciliate making http requests to Slack.
// Client does not attempt to evaluate the response body, leaving that to the caller.
func NewClient(client *http.Client, host string) *Client {
	return &Client{host: host, httpClient: client}
}

// PostMessage calls the chat.postMessage Slack endpoint to post a message into a channel
func (c *Client) Get(ctx context.Context) (Data, error) {
	var req, err = http.NewRequest("GET", c.host, bytes.NewReader([]byte{})) //bytes.NewReader(body)
	if err != nil {
		return Data{}, err
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var token = os.Getenv("WeatherToken")

	req.Header.Add("x-rapidapi-key", token)
	req.Header.Add("x-rapidapi-host", "weatherbit-v1-mashape.p.rapidapi.com")

	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return Data{}, err
	}
	defer resp.Body.Close()

	var datas = dataResp{}
	err = json.NewDecoder(resp.Body).Decode(&datas)
	if err != nil {
		return Data{}, err
	}
	if len(datas.Datas) < 1 {
		return Data{}, errors.New("no results")
	}
	fmt.Println("Got Weather in Denver:", datas.Datas[0])

	return datas.Datas[0], nil
}
