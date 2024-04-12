package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

const Usage = `gotest is a command-line tool that makes test names more human-readable.

Usage:

	gotest [ARGS]

This will run 'go test -json [ARGS]' in the current directory and format the results in a readable
way. You can use any arguments that 'go test' accepts.`

func main() {

	color.NoColor = false // force enable colors regardless of the environment|terminal settings

	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Println(Usage)
		os.Exit(0)
	}
	tr := NewTransformer()
	tr.Execute(os.Args[1:])

	// If tests failed, exit with status 1
	if !tr.PASS {
		os.Exit(1)
	}
}
