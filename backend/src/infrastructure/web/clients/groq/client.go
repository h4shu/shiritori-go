package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/h4shu/shiritori-go/infrastructure/web/models"
)

const (
	GroqUrl = "https://api.groq.com/openai/v1/chat/completions"
)

type (
	GroqClient struct {
		client *http.Client
	}

	ErrGroqFailed struct{}
)

func NewGroqClient() *GroqClient {
	client := &http.Client{}
	return &GroqClient{
		client: client,
	}
}

func (c *GroqClient) ValidateJapanese(word string) (bool, error) {
	mesg := models.GroqMessage{
		Role:    "user",
		Content: fmt.Sprintf("「%s」は日本語として正しいですか？正しい場合は \"T\" 、正しくない場合は \"F\" の一文字のみを返却してください。", word),
	}
	greq := models.GroqRequest{
		Messages:    []models.GroqMessage{mesg},
		Model:       "llama3-8b-8192",
		Temperature: 0,
		Max_tokens:  1,
		Top_p:       0,
	}

	reqb, err := json.Marshal(greq)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", GroqUrl, bytes.NewBuffer(reqb))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))
	log.Printf("%v\n", req) // TEST
	res, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	log.Println(string(resb)) // TEST
	var gres models.GroqResponse
	err = json.Unmarshal(resb, &gres)
	if err != nil {
		return false, err
	}

	if len(gres.Choices) == 0 {
		return false, &ErrGroqFailed{}
	}
	content := gres.Choices[0].Message.Content
	switch content {
	case "T":
		return true, nil
	case "F":
		return false, nil
	}
	return false, &ErrGroqFailed{}
}

func (*ErrGroqFailed) Error() string {
	return "Groq Error: 判定に失敗しました"
}
