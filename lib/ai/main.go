package ai

import (
	"fmt"
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Prompt(apiKey string, prompt string) string {
	ctx := context.Background()
	opt := option.WithAPIKey(apiKey)
	client, _ := genai.NewClient(ctx, opt)
	model := client.GenerativeModel("gemini-1.5-flash")

	text := genai.Text(prompt)

	resp, _ := model.GenerateContent(ctx, text)

	return buildOutput(resp)


}


func buildOutput(resp *genai.GenerateContentResponse) string {
	out := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				out += fmt.Sprintf("%s", part)
			}
		}
	}
	return out
}
