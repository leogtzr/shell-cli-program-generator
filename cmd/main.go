package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leogtzr/shellcligen"
)

func run() error {
	inputFile := flag.String("input", "", "configuration file")
	outputFile := flag.String("output", "", "output directory where the script will be generated")

	flag.Parse()

	if len(*inputFile) == 0 {
		return shellcligen.ErrMissingRequiredArgument
	}

	if len(*outputFile) == 0 {
		return shellcligen.ErrMissingRequiredArgument
	}

	cli, err := shellcligen.ParseCLIProgram(*inputFile, *outputFile)
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
