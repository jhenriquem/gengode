package gengode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// API request structure
type OpenRouterRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

// API response structure
type OpenRouterResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func RequestForIA(message string) (OpenRouterResponse, error) {
	url := os.Getenv("OPENROUTER_URL")

	// Create request payload
	requestBody := OpenRouterRequest{
		Model: os.Getenv("OPENROUTER_MODEL"),
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{"user", message},
		},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return OpenRouterResponse{}, fmt.Errorf("Error converting request body : %s", err.Error())
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return OpenRouterResponse{}, fmt.Errorf("Error creating http request: %s", err.Error())
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENROUTER_KEY"))
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OpenRouterResponse{}, fmt.Errorf("Error sending http resquest: %s", err.Error())
	}
	defer resp.Body.Close()

	// Parse response
	var result OpenRouterResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return OpenRouterResponse{}, fmt.Errorf("Error parsing response: %s", err.Error())
	}

	return result, nil
}
