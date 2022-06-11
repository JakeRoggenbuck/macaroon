package main

import (
	"fmt"
)

func main() {
	args := parseArgs()

	if args.hasFilename {
		fmt.Println(args.filename)
	}
}
