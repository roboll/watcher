package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/roboll/watcher/handlers"
)

var (
	webhookURL    = os.Getenv("WEBHOOK_URL")
	webhookMethod = os.Getenv("WEBHOOK_VERB")
)

func init() {
	handlers.Register("webhook", Webhook)
}

func Webhook(path string) error {
	log.Printf("webhook: %s", path)
	req, err := http.NewRequest(webhookMethod, webhookURL, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("webhook: status code was %d", resp.StatusCode)
	}
	return nil
}
