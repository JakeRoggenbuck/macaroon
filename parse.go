package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fast_forward_empty_space(current string) int {
	index := 0

	for {
		if index+1 > len(current) {
			break
		}
		cur := current[index]
		if cur != '\t' && cur != ' ' {
			break
		}

		index += 1

		if index > 999 {
			fatal("Too much space before line")
		}
	}

	return index
}

func check_start(current string) int {
	index := fast_forward_empty_space(current)

	if current[index] == '#' {
		return index
	}

	return -1
}

type Command int

const (
	Set Command = iota + 1
	None
)

func lex_command(line string) Command {
	if len(line) < 4 {
		return None
	}
	if line[0:4] == "#set" {
		return Set
	}
	return None
}

type Macro struct {
	command Command
	name    string
	value   string
}

func tokenize(line string, command Command, start int) Macro {
	var name string
	var value string
	index := start
	first, second := -1, -1

	for {
		if index+1 > len(line) {
			break
		}
		cur := line[index]
		if cur == ' ' {
			// Second value has already been found
			if second > 0 {
				break
			}

			// Assign index to either first or second based on if they have been found yet
			if first < 0 {
				first = index
			} else {
				second = index
			}
		}

		index += 1
	}

	name = line[first+1 : second]
	value = line[second+1:]

	return Macro{command, name, value}
}

func contains(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}

func parse(filename string) {
	var macros []Macro
	var macro_lines []int

	in_file, err := os.Open(filename)
	if err != nil {
		fatal("Error opening input file")
	}

	out_file, err := os.Create(filename + ".go")
	if err != nil {
		fatal("Error writing to output file")
	}

	defer in_file.Close()
	defer out_file.Close()

	line_num := 0
	scanner := bufio.NewScanner(in_file)
	for scanner.Scan() {
		current := scanner.Text()

		// Add macros
		if len(current) > 1 {
			if start := check_start(current); start != -1 {
				command := lex_command(current)
				if command == None {
					continue
				}
				macro := tokenize(current, command, start)
				macros = append(macros, macro)
				macro_lines = append(macro_lines, line_num)
			}
		}

		// Use macros
		for _, mac := range macros {
			if cur_start := strings.Index(current, mac.name); cur_start != -1 {
				if !contains(macro_lines, line_num) {
					new_line := current[:cur_start] + mac.value + current[cur_start+len(mac.name):]
					out_file.WriteString(new_line + "\n")
					continue
				}
			}

			if !contains(macro_lines, line_num) {
				out_file.WriteString(current + "\n")
			}
		}

		line_num += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
