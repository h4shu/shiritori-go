package models

type (
	GroqMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	GroqRequest struct {
		Messages    []GroqMessage `json:"messages"`
		Model       string        `json:"model"`
		Temperature float32       `json:"temperature"`
		Max_tokens  int           `json:"max_tokens"`
		Top_p       float32       `json:"top_p"`
	}
	GroqChoice struct {
		Message GroqMessage `json:"message"`
	}
	GroqResponse struct {
		Choices []GroqChoice `json:"choices"`
	}
)
