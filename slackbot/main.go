package main

import (
	"fmt"
	"net/http"

	"github.com/dustinlindquist/denstartupweektalk/slackbot/handlers"
	"github.com/dustinlindquist/denstartupweektalk/slackbot/slack"
	"github.com/dustinlindquist/denstartupweektalk/slackbot/weather"
)

func main() {
	var weather = weather.NewClient(http.DefaultClient, "https://weatherbit-v1-mashape.p.rapidapi.com/current?lat=39.74&lon=-104.99&lang=en&units=i")
	var slack = slack.NewClient(http.DefaultClient, "https://slack.com")

	var handler = handlers.New(slack, weather)
	fmt.Println("Service Starting")
	if err := http.ListenAndServe("localhost:9000", handler.Message()); err != nil {
		panic(err)
	}
}
