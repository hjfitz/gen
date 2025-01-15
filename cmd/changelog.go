package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/hjfitz/gen/lib/ai"
	"github.com/hjfitz/gen/lib/git"
	"github.com/hjfitz/gen/prompts"
)

func validateChangelogArgs(depth *int, apiKey *string) {
	isValid := (*depth != 0 && *apiKey != "")

	if !isValid {
		fmt.Printf("Usage: gen changelog -a <api-key> -d <commit-depth> -t (optional)\n")
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

	validateChangelogArgs(depth, apiKey)

	wd, _ := os.Getwd()

	diff := git.GetDiff(wd, ds)

	cp := prompts.GetChangelog(diff, *trump)

	out := ai.Prompt(*apiKey, cp)

	fmt.Println(out)

}
