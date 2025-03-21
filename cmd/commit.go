package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/hjfitz/gen/lib/ai"
	"github.com/hjfitz/gen/lib/exec"
	//"github.com/hjfitz/gen/lib/exec"
	"github.com/hjfitz/gen/lib/git"
	"github.com/hjfitz/gen/prompts"
)

func GenerateCommit() {

	fs := flag.NewFlagSet("commit", flag.ExitOnError)
	apiKey := fs.String("a", "", "Gemini API Key")
	shouldPrint := fs.Bool("p", false, "Whether to print the commit instead of making it")

	fs.Parse(os.Args[2:])

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}

	diff := git.GetMostRecentChanges()

	cp := prompts.GetCommit(diff)

	msg := ai.Prompt(*apiKey, cp)

	if *shouldPrint {
		fmt.Println(msg)
	} else {
		command := fmt.Sprintf("git commit -m \"%s\"", msg)
		exec.Exec(command)
	}
}
