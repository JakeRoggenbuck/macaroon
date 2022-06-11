package main

func main() {
	args := parseArgs()

	if !args.hasFilename {
		fatal("File not provided")
	}

	parse(args.filename)

}
