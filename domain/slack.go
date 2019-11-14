package domain

import (
	"encoding/json"
	"fmt"
)

// Message represents the structure of a Slack message sent as an output of the App
type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func (m *Message) String() (string, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("unable to marshal the message into a string: %v", err)
	}
	return string(data), nil
}

func (m *Message) createFormattedText(characters *[]Character) error {
	m.Text = "Hello World!"
	return nil
}
