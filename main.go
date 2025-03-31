package main

import (
	"fmt"
	"os"

	"github.com/hjfitz/gen/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'command' subcommand")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "changelog":
		cmd.GenerateChangelog()
	case "readme":
		cmd.GenerateReadme()
	case "commit":
		fallthrough
	case "cmt":
		cmd.GenerateCommit()
	case "arch":
		cmd.GenerateArchitecture()
	default:
		fmt.Printf("Unknown subcommand: \"%s\"\n", command)
		fmt.Printf("Usage: `gen (readme|changelog)`")
		os.Exit(1)
	}
}
