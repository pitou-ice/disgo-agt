package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	PROMPT_TEMPLATE = "<im_start>system\n%s\nYou're talking to %s.<|im_end|>\n<|im_start|>user\n%s<|im_end|>\n<|im_start|>assistant\n"
	N_PREDICT       = 134
	STOP_WORD       = "<|im_end|>"
	TEMPERATURE     = 2.0
	CACHED_PROMPT   = true
)

var SystemPrompt string

// Struct representing the JSON body
type RequestBody struct {
	Prompt       string  `json:"prompt"`
	NPredict     int     `json:"n_predict"`
	Stop         string  `json:"stop"`
	Temperature  float32 `json:"temperature"`
	CachedPrompt bool    `json:"cached_prompt"`
}

// Define a struct to represent the structure of the JSON data
type ResponseData struct {
	Content string `json:"content"`
}

func GetCompletion(userName string, userPrompt string) (string, error) {
	// Create JSON body
	requestBody := RequestBody{
		Prompt:       fmt.Sprintf(PROMPT_TEMPLATE, SystemPrompt, userName, userPrompt),
		NPredict:     N_PREDICT,
		Stop:         STOP_WORD,
		Temperature:  TEMPERATURE,
		CachedPrompt: CACHED_PROMPT,
	}

	// Convert struct to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "Error marshalling JSON:", err
	}

	// Make HTTP POST request
	response, err := http.Post("http://localhost:8080/completion", "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return "Error making HTTP POST:", err
	}
	defer response.Body.Close()

	var data ResponseData

	// Parse the JSON data directly from response.Body
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "Error parsing JSON:", err
	}

	return strings.TrimSuffix(data.Content, STOP_WORD), nil
}
