package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dustinlindquist/denstartupweektalk/slackbot/weather"
)

// Slack holds the methods for interacting with Slack's api.
type Slack interface {
	SendMessage(ctx context.Context, channel, msg string) error
}

// Weather holds the methods
type Weather interface {
	Get(ctx context.Context) (weather.Data, error)
}

// Handler is a struct that has all the handler functions and implementations of the dependencies for the Workflow API.
type Handler struct {
	slack   Slack
	weather Weather
}

// New returns an WorkflowHandler with wired up dependencies and routes set.
func New(slack Slack, weather Weather) *Handler {
	var h = &Handler{
		slack:   slack,
		weather: weather,
	}
	return h
}

type message struct {
	Message string `json:"message"`
}

func (h *Handler) Message() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		data, err := h.weather.Get(ctx)
		if err != nil {
			fmt.Println("error getting weather", err)
		}

		var msg message
		if err = json.NewDecoder(r.Body).Decode(&msg); err != nil {
			fmt.Println("error decoding input", err)
		}
		defer r.Body.Close()

		var text = fmt.Sprintf(`%s It's a beatiful %d °F and %s`, msg.Message, int(data.Temp), data.Weather.Description)

		err = h.slack.SendMessage(ctx, "CNCGNLJ68", text)
		if err != nil {
			fmt.Println("Error posing message:", err)
		}

	})
}
