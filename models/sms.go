package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SMSRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func sendSMS(apiKey string, sms SMSRequest) error {
	url := "https://api.semaphore.co/api/v4/messages"

	// Convert SMSRequest to JSON
	smsJSON, err := json.Marshal(sms)
	if err != nil {
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(smsJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS: %s", resp.Status)
	}

	return nil
}
