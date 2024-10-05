package discordlogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DiscordLogger struct {
	webhookUrl string
	logErrors  bool
}

type discordWebhookPayload struct {
	Content string `json:"content"`
}

func NewDiscordLogger(webhookUrl string, logErrors bool) *DiscordLogger {
	return &DiscordLogger{webhookUrl: webhookUrl, logErrors: logErrors}
}

func (d *DiscordLogger) Printf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	err := d.sendWebhookMessage(message)
	if (err != nil) && d.logErrors {
		log.Printf("discord wh log error: %v", err)
	}
}

func (d *DiscordLogger) Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	err := d.sendWebhookMessage("**Error:** " + message)
	if (err != nil) && d.logErrors {
		log.Printf("discord wh error error: %v", err)
	}
}

func (d *DiscordLogger) sendWebhookMessage(message string) error {
	payload := discordWebhookPayload{Content: message}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", d.webhookUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send discord webhook, status code: %d", resp.StatusCode)
	}

	return nil
}
