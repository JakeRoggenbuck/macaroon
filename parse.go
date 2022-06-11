package main

import (
	"os"
	"bufio"
	"fmt"
)

func fast_forward_empty_space(current string) int {
	index := 0

	for {
		if index + 1 > len(current) { break; }
		cur := current[index]
		if cur != '\t' && cur != ' ' { break; }

		index += 1

		if index > 999 {
			fatal("Too much space before line")
		}
	}

	return index
}

func check_start(current string) bool {
	index := fast_forward_empty_space(current)

	if current[index] == '#' {
		return true
	}

	return false
}

type Command int

const (
	Set Command = iota + 1
	None
)

func lex(line string) Command {
	if len(line) < 4 { return None }
	if line[0:4] == "#set" {
		return Set
	}
	return None 
}

func parse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fatal("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		if len(current) < 1 { continue; }
		if status := check_start(current); status {
			fmt.Println(current)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
