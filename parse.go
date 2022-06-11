package main

import (
	"os"
	"bufio"
	"fmt"
)

func parse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fatal("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		fmt.Println(current)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
