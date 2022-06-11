package main

import (
	"os"
)

type Args struct {
	filename    string
	hasFilename bool
}

func parseArgs() Args {
	argLength := len(os.Args[1:])

	if argLength == 0 {
		fatal("Too few arguments, please provide a filename")
	}

	filename := os.Args[1]

	return Args{
		filename,
		true,
	}
}
