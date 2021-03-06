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
