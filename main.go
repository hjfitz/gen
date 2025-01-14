package main

import (
	"fmt"
	"os"

	"github.com/hjfitz/agentic-workflow/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'command' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "changelog":
		cmd.GenerateChangelog()
	case "readme":
		cmd.GenerateReadme()
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
	}
}
