/**
This file handles the generation from AI of some content for the blog
*/

package blog

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

type GenerateContentOptions struct {
	Instruction string
	Max_Length  int
}

func (blog *Blog) GenerateContent(options GenerateContentOptions) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
