package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MistralClient struct {
	baseUrl string
	apiKey  string
}

func NewMistralClient(getenv func(string) string) (*MistralClient, error) {
	apiKey := getenv("MISTRAL_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("No MISTRAL_API_KEY environment variable found")
	}

	return &MistralClient{
		apiKey:  apiKey,
		baseUrl: "https://api.mistral.ai",
	}, nil
}

func (c *MistralClient) GetCompletion(r *CompletionsRequestBody) (*Message, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v1/chat/completions", c.baseUrl),
		bytes.NewBuffer(data),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%d\n", res.StatusCode)

	var rspMsg MistralCompletionsResponse
	if err := json.Unmarshal(body, &rspMsg); err != nil {
		return nil, err
	}

	if len(rspMsg.Choices) == 0 {
		return nil, fmt.Errorf("Response message had no choices")
	}

	return &rspMsg.Choices[0].Message, nil
}
