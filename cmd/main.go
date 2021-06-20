package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/leogtzr/shellcligen"
)

var (
	input = flag.String("input", "", "script input file")
)

func main() {
	flag.Parse()

	if len(*input) == 0 {
		fmt.Fprintln(os.Stderr, "error: missing required input script file")
		os.Exit(3)
	}

	cli, err := shellcligen.ParseCLIProgram(*input)
	if err != nil {
		log.Fatal(err)
	}

	for _, opt := range cli.Options {
		fmt.Println(opt.ArgsNum)
	}
}
