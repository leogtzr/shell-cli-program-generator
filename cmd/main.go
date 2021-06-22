package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/leogtzr/shellcligen"
)

func main() {
	input := flag.String("input", "", "script input file")

	flag.Parse()

	if len(*input) == 0 {
		fmt.Fprintln(os.Stderr, "error: missing required input script file")
		os.Exit(1)
	}

	cli, err := shellcligen.ParseCLIProgram(*input)
	if err != nil {
		log.Fatal(err)
	}

	for _, opt := range cli.Options {
		fmt.Println(opt.ArgsNum)
	}
}
