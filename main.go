package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

func fatal(message string) {
	color.Yellow.Print("FATAL: ")
	fmt.Println(message)
	os.Exit(1)
}

func warn(message string) {
	color.Yellow.Print("WARN: ")
	fmt.Println(message)
}

type Args struct {
	filename string
	hasFilename bool
}

func parseArgs() Args {
	argLength := len(os.Args[1:])

	if argLength == 0 {
		fatal("Too few arguments, please provide a filename")
	}

	filename := os.Args[1]

	return Args {
		filename,
		true,
	}
}

func main() {
	args := parseArgs()

	if args.hasFilename {
		fmt.Println(args.filename)
	}
}
