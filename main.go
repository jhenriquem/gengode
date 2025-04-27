package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

var (
	apiKey  string = ""
	modelIA string = ""
)

func RequestForIA(message string) OpenRouterResponse {
	url := "https://openrouter.ai/api/v1/chat/completions"

	// Create request payload
	requestBody := OpenRouterRequest{
		Model: modelIA,
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
		panic(err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse response
	var result OpenRouterResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

// Initial function, variable assignment
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey = os.Getenv("OPENROUTER_KEY")
	modelIA = os.Getenv("OPENROUTER_MODEL")
}

func main() {
	message := "Hi!"
	result := RequestForIA(message)

	if len(result.Choices) > 0 {
		fmt.Println("Response:", result.Choices[0].Message.Content)
	}
}
