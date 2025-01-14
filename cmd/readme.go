package cmd

import (
	"flag"
	"os"
	"fmt"
)

func GenerateReadme() {
	fs := flag.NewFlagSet("subcommand1", flag.ExitOnError)
	option := fs.String("option", "default", "An option for subcommand1")
	fs.Parse(os.Args[2:])

	fmt.Printf("Running subcommand1 with option: %s\n", *option)
}
