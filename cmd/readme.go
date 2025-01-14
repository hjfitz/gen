package cmd

import (
	"flag"
	"os"
	"fmt"

	"github.com/hjfitz/agentic-workflow/lib/fs"
)

func validateReadmeGenArgs(apiKey *string) bool {
	return *apiKey != ""
}

func GenerateReadme() {
	fls := flag.NewFlagSet("readme", flag.ExitOnError)

	//trump := fs.Bool("t", false, "Make changelogs great again")
	apiKey := fls.String("a", "", "Gemini API Key")

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}

	fls.Parse(os.Args[2:])

	validateReadmeGenArgs(apiKey)

	fmt.Printf("Running subcommand1 with option: %s\n", *apiKey)

	wd, _ := os.Getwd()
	
	cb := fs.GetCodebase(wd, []string{})

	fmt.Println(cb)
}
