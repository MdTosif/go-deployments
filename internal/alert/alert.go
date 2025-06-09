package alert

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MdTosif/go-deployments/internal/config"
)

// Payload structure sent to Slack webhook
type slackPayload struct {
    Text string `json:"text"`
}

// Alert sends a Slack message using the given webhook URL.
// It reads SLACK_WEBHOOK_URL from the environment.
func Alert(message string) {
    webhook := config.Cfg.Slack.WebhookURL;

    if webhook == "" {
        log.Println("⚠️ Slack webhook is not set")
        return
    }

    payload := slackPayload{Text: message}
    body, err := json.Marshal(payload)
    if err != nil {
        log.Printf("❌ Error marshalling Slack payload: %v\n", err)
        return
    }

    client := &http.Client{Timeout: 5 * time.Second}
    req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(body))
    if err != nil {
        log.Printf("❌ Failed to create Slack request: %v\n", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("❌ Error sending Slack message: %v\n", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Printf("❌ Slack responded with non-OK status: %s\n", resp.Status)
    } else {
        log.Println("✅ Slack alert sent successfully")
    }
}

// Helper to read env variables
func getenv(key string) string {
    return os.Getenv(key)
}
