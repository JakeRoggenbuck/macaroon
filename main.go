package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

func warn(message string) {
	color.Yellow.Print("WARN: ")
	fmt.Println(message)
}

func main() {
	argLength := len(os.Args[1:]) 

	if argLength == 0 {
		warn("Too few arguments, please provide a filename")
		return
	}

	filename := os.Args[1]
	fmt.Println(filename)
}
