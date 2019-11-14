package http

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// PostInitiatives sends a message to a Slack channel with the initiatives for an encounter
func PostInitiatives(messageBody string) error {
	webHookURL := os.Getenv("SLACK_WEB_HOOK")
	if webHookURL == "" {
		return fmt.Errorf("unable to post the message to Slack due to missing the web hook environment variable")
	}
	resp, err := http.Post(webHookURL, "application/json", bytes.NewBufferString(messageBody))
	if err != nil {
		return fmt.Errorf("error whilst sending the message to the Slack channel: %v", err)
	}
	fmt.Println(resp)
	return nil
}
