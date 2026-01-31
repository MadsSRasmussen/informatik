package ai

import (
	"encoding/json"
	"fmt"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleSystem    Role = "system"
	RoleAssistant Role = "assistant"
)

func (r *Role) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "user", "system", "assistant":
		*r = Role(s)
		return nil
	default:
		return fmt.Errorf("invalid role: %s\n", s)
	}
}

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type CompletionsRequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type MistralResponseMessageWrapper struct {
	Message Message `json:"message"`
}

type MistralCompletionsResponse struct {
	Choices []MistralResponseMessageWrapper `json:"choices"`
}
