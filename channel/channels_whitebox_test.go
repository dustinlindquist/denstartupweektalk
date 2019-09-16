package channel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCleanName(t *testing.T) {
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
		},
		{
			name:       "uppercase_letters", // sub-test names should use "_" as a delimiter.
			nameIn:     "GENerAL",
			expNameOut: "general",
		},
		{
			name: "spaces_and_uppercase",
			nameIn: " 	 GenERaL  		",
			expNameOut: "general",
		},
	}

	for _, tc := range tests {
		tc := tc // need this to avoid the go-routine and loop varible common mistake.
		t.Run(tc.name, func(t *testing.T) {
			var result = cleanName(tc.nameIn)
			assert.Equal(t, tc.expNameOut, result)
		})
	}

}
