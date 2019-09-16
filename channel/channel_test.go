package channel_test // separate package from the demo package where the func we're testing lives.

import (
	"testing"
	"time"

	"github.com/dustinlindquist/denstartupweektalk/channel"
	"github.com/stretchr/testify/assert"
)

func TestSetName(t *testing.T) {
	// t.Parallel() // flag used for runing this top level test in parallell with other top level tests.
	var tests = []struct {
		name       string
		nameIn     string
		expNameOut string
		duration   time.Duration
	}{
		{
			name: "spaces",
			nameIn: " 	general   ",
			expNameOut: "general",
			duration:   2 * time.Second,
		},
		{
			name:       "uppercase_letters", // sub-test names should use "_" as a delimiter.
			nameIn:     "GENerAL",
			expNameOut: "general",
			duration:   3 * time.Second,
		},
		{
			name: "spaces_and_uppercase",
			nameIn: " 	 GenERaL  		",
			expNameOut: "general",
			duration:   3 * time.Second,
		},
	}

	for _, tc := range tests {
		tc := tc // need this to avoid the go-routine and loop varible common mistake.
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // flag used for running each of these table (sub) tests in paralell.
			// time.Sleep(tc.duration)
			var channel = &channel.SlackChannel{
				ID:   "12345",
				Name: "random",
			}

			channel.SetName(tc.nameIn)
			assert.Equal(t, tc.expNameOut, channel.Name)
			assert.Equal(t, "12345", channel.ID) // assert the id didn't change.
		})
	}

}

// emoji text: `   	   😀😁😂🤣-GeNeRaL`
