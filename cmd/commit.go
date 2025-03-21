package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/hjfitz/gen/lib/ai"
	"github.com/hjfitz/gen/lib/git"
	"github.com/hjfitz/gen/prompts"
)

func GenerateCommit() {

	fs := flag.NewFlagSet("commit", flag.ExitOnError)
	apiKey := fs.String("a", "", "Gemini API Key")

	fs.Parse(os.Args[2:])

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}

	wd, _ := os.Getwd()
	diff := git.GetDiff(wd, "2")

	cp := prompts.GetCommit(diff)

	out := ai.Prompt(*apiKey, cp)

	fmt.Println(out)
}
