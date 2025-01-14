package cmd

import (
	"flag"
	"os"
	"fmt"
	"strconv"

	"github.com/hjfitz/agentic-workflow/lib/git"
)

func validate(depth *int, apiKey *string) bool {
	return *depth != 0 && *apiKey != ""
}

func GenerateChangelog() {
	fs := flag.NewFlagSet("changelog", flag.ExitOnError)

	depth := fs.Int("d", 0, "How many commits to analyse")
	trump := fs.Bool("t", false, "Make changelogs great again")
	apiKey := fs.String("a", "", "Gemini API Key")

	fs.Parse(os.Args[2:])

	ds := strconv.Itoa(*depth)

	fmt.Printf("Running changelog with trump=%t, depth=%d, apiKey=%s\n", *trump, *depth, *apiKey)
	fmt.Printf("Valid: %t\n", validate(depth, apiKey))

	if (!validate(depth, apiKey)) {
		fmt.Printf("Usage: agentic -a <api-key> -d <commit-depth> -t (optional)\n")
		os.Exit(1)
	}

	wd, _ := os.Getwd()

	fmt.Println(wd)

	diff := git.GetDiff(wd, ds)

	fmt.Println(diff)

}
