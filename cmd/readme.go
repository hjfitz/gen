package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/hjfitz/agentic-workflow/lib/ai"
	"github.com/hjfitz/agentic-workflow/lib/fs"
	"github.com/hjfitz/agentic-workflow/prompts"
)

func validateReadmeGenArgs(apiKey *string) bool {
	return *apiKey != ""
}

func GenerateReadme() {
	fls := flag.NewFlagSet("readme", flag.ExitOnError)

	trump := fls.Bool("t", false, "Make your readme great again")
	apiKey := fls.String("a", "", "Gemini API Key")

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}

	fls.Parse(os.Args[2:])

	validateReadmeGenArgs(apiKey)

	wd, _ := os.Getwd()

	cb := fs.GetCodebase(wd)

	rp := prompts.GetReadme(cb, *trump)

	out := ai.Prompt(*apiKey, rp)

	fmt.Println(out)
}
