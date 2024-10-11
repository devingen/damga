package main

import (
	"fmt"
	"os"
)

const desc = `Usage: damga <tg>
Examples:
  damga tg add --id abcd --target 1.2.3.4
  damga tg remove --id abcd --target 1.2.3.4
`

func main() {

	var err error

	if len(os.Args) == 1 {
		fmt.Print(desc)
		return
	}
	resource := os.Args[1]
	switch resource {
	case "tg":
		err = handleTargetCommand()
	case "-h", "--help":
		fmt.Print(desc)
	case "-v", "--version":
		fmt.Println("damga v0.0.1")
	default:
		fmt.Println("Unknown command.")
		fmt.Print(desc)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
