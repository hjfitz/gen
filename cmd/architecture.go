package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dreampuf/mermaid.go"

	"github.com/hjfitz/gen/lib/ai"
	"github.com/hjfitz/gen/lib/fs"
	"github.com/hjfitz/gen/prompts"
)

func GenerateArchitecture() {
	fset := flag.NewFlagSet("arch", flag.ExitOnError)
	apiKey := fset.String("a", "", "Gemini API Key")
	should_render := fset.Bool("r", false, "Whether to render the diagram or not")

	fset.Parse(os.Args[2:])

	if *apiKey == "" {
		*apiKey = os.Getenv("GEMINI_API_KEY")
	}

	wd, _ := os.Getwd()
	tf := fs.GetIAC(wd)

	prompt := prompts.GetArchitectureGenPrompt(tf)

	out := ai.Prompt(*apiKey, prompt)

	if out[0] == '`' {
		// remove the first line
		out_lines := strings.Split(out, "\n")
		out = ""
		for i, line := range out_lines {
			if i == 0 || i == (len(out_lines) - 1) {
				continue
			}
			out += line
			out += "\n"

		}

		out = strings.TrimSpace(out)
	}


	if *should_render {
		ctx := context.Background()
		re, _ := mermaid_go.NewRenderEngine(ctx)
	}


	fmt.Println(out)
}
