package ai

type AIClient interface {
	GetCompletion(r *CompletionsRequestBody) (*Message, error)
}
