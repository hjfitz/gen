package cmd

import (
	"flag"
	"os"
	"fmt"
	"strconv"

	"github.com/hjfitz/agentic-workflow/lib/git"
	"github.com/hjfitz/agentic-workflow/lib/ai"
	"github.com/hjfitz/agentic-workflow/prompts"

)

func validate(depth *int, apiKey *string) {
	isValid := (*depth != 0 && *apiKey != "")

	if !isValid {
		fmt.Printf("Usage: agentic -a <api-key> -d <commit-depth> -t (optional)\n")
		os.Exit(1)
	}
}

func GenerateChangelog() {
	fs := flag.NewFlagSet("changelog", flag.ExitOnError)

	depth := fs.Int("d", 0, "How many commits to analyse")
	trump := fs.Bool("t", false, "Make changelogs great again")
	apiKey := fs.String("a", "", "Gemini API Key")

	fs.Parse(os.Args[2:])

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}


	ds := strconv.Itoa(*depth)

	/*
	fmt.Printf("Running changelog with trump=%t, depth=%d, apiKey=%s\n", *trump, *depth, *apiKey)
	fmt.Printf("Valid: %t\n", validate(depth, apiKey))
	*/

	validate(depth, apiKey)

	wd, _ := os.Getwd()

	diff := git.GetDiff(wd, ds)

	cp := prompts.GetChangelog(diff, *trump)

	out := ai.Prompt(*apiKey, cp)

	fmt.Println(out)

}
