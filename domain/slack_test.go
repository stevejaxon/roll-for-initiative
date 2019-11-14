package domain

import (
	"testing"

	"github.com/stevejaxon/roll-for-initiative/http"
)

func TestPostMessage(t *testing.T) {
	// Setup
	message := &Message{
		Channel: "general",
	}
	message.createFormattedText(nil)
	body, err := message.String()
	if err != nil {
		t.Fatal("error marshalling the message as a string")
	}
	http.PostInitiatives(body)
}
