package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leogtzr/shellcligen"
)

func run() error {
	input := flag.String("input", "", "script input file")

	flag.Parse()

	if len(*input) == 0 {
		return shellcligen.ErrMissingRequiredArgument
	}

	cli, err := shellcligen.ParseCLIProgram(*input)
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}

	for _, opt := range cli.Options {
		fmt.Println(opt)
		fmt.Println(opt.ConflictsWith)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
