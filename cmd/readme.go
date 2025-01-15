package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/hjfitz/gen/lib/ai"
	"github.com/hjfitz/gen/lib/fs"
	"github.com/hjfitz/gen/prompts"
)

func validateReadmeArgs(apiKey *string) {
	isValid := *apiKey != ""

	if !isValid {
		fmt.Printf("Usage: gen readme -a <api-key> -t (optional)\n")
		os.Exit(1)
	}
}

func GenerateReadme() {
	fls := flag.NewFlagSet("readme", flag.ExitOnError)

	trump := fls.Bool("t", false, "Make your readme great again")
	apiKey := fls.String("a", "", "Gemini API Key")

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}

	fls.Parse(os.Args[2:])

	validateReadmeArgs(apiKey)

	wd, _ := os.Getwd()

	cb := fs.GetCodebase(wd)

	rp := prompts.GetReadme(cb, *trump)

	out := ai.Prompt(*apiKey, rp)

	fmt.Println(out)
}
