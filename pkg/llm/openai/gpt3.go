package openai

import (
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/sashabaranov/go-gpt3"
)

var (
	oaiToken = os.Getenv("OAI_TOKEN")
)

func Prompt(prompt string) string {
	fmt.Println("prompting openai with: ", prompt)
	c := gpt3.NewClient(oaiToken)
	ctx := context.Background()

	req := gpt3.CompletionRequest{
		Model:     gpt3.GPT3Ada, // or gpt3.GPT3Babbage, gpt3.GPT3Curie, gpt3.GPT3Davinci
		MaxTokens: 10,           // token are word-like "chunks" of text [10, 20, 50, 100, 200]
		Prompt:    prompt,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "error"
	}

	log.Println(resp)
	log.Println(resp.Choices[0].Text)

	return resp.Choices[0].Text
}
