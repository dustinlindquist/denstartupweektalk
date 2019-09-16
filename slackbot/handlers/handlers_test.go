package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dustinlindquist/denstartupweektalk/slackbot/handlers"
	"github.com/dustinlindquist/denstartupweektalk/slackbot/mocks"
	"github.com/dustinlindquist/denstartupweektalk/slackbot/weather"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type handlerMocks struct {
	slack   handlers.Slack
	weather handlers.Weather
}

func (m handlerMocks) assertExpectations(t *testing.T) {
	if m.slack != nil {
		mock.AssertExpectationsForObjects(t, m.slack)
	}
	if m.weather != nil {
		mock.AssertExpectationsForObjects(t, m.weather)
	}
}

func TestHandler_Message(t *testing.T) {
	t.Parallel() // flag used for runing this top level test in parallell with other top level tests.
	var chanID = "CNCGNLJ68"
	var tests = []struct {
		name        string
		payload     string
		expRespCode int
		mocks       handlerMocks
	}{
		{
			name:        "success",
			payload:     `{"message": "Hello Denver Startup Week!"}`,
			expRespCode: http.StatusOK,
			mocks: func() handlerMocks {
				var weatherMock = new(mocks.Weather)
				var weatherResp = weather.Data{
					Temp: 91.11,
					Weather: weather.Weather{
						Description: "Scattered clouds",
					},
				}
				weatherMock.On("Get", mock.Anything).Return(weatherResp, nil)

				var inputText = fmt.Sprintf(`Hello Denver Startup Week! It's a beatiful %d °F and %s`, 91, weatherResp.Weather.Description)
				var slackMock = new(mocks.Slack)
				slackMock.On("SendMessage", mock.Anything, chanID, inputText).Return(nil)
				return handlerMocks{slack: slackMock, weather: weatherMock}
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc     // need this to avoid the go-routine and loop varible common mistake.
			t.Parallel() // flag used for running each of these table (sub) tests in paralell.

			var handler = handlers.New(tc.mocks.slack, tc.mocks.weather)
			var server = httptest.NewServer(handler.Message())
			defer server.Close()

			var url = server.URL + "/message"
			var req, err = http.NewRequest("POST", url, strings.NewReader(tc.payload))
			assert.NoError(t, err)

			var client = http.DefaultClient
			resp, err := client.Do(req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expRespCode, resp.StatusCode)

			tc.mocks.assertExpectations(t)
		})
	}
}
