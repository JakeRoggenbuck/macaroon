package main

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
)

func warn(message string) {
	color.Yellow.Print("WARN: ")
	fmt.Println(message)
}

func main() {
	var filename string
	
	flag.StringVar(&filename, "f", "default", "provide a file to preprocess")
	flag.Parse()

	if filename == "default" {
		warn("Please provide a filename")
		return
	}

	fmt.Println(filename)
}
