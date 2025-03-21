package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Prompt(apiKey string, prompt string) string {
	ctx := context.Background()
	opt := option.WithAPIKey(apiKey)
	client, _ := genai.NewClient(ctx, opt)
	model := client.GenerativeModel("gemini-2.0-flash")

	text := genai.Text(prompt)

	resp, err := model.GenerateContent(ctx, text)

	if err != nil {
		fmt.Printf("Got error: %s\n", err)
	}

	out := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				out += fmt.Sprintf("%s", part)
			}
		}
	}
	return strings.TrimSpace(out)
}
