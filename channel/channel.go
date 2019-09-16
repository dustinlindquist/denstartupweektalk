package channel

import (
	"strings"
)

type SlackChannel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *SlackChannel) SetName(name string) {
	var cleanedName = cleanName(name)
	c.Name = cleanedName
}

// cleanName trims leading and trailing whitespace and lowers all characters
// used to clean channel names for checking if a #channel-name already exists.
func cleanName(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	return name
}
