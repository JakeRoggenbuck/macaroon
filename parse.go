package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
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
	if len(line) < 4 { return None }
	if line[0:4] == "#set" {
		return Set
	}
	return None 
}

type Macro struct {
	command Command
	name string
	value string
}

func tokenize(line string, command Command, start int) Macro {
	var name string
	var value string
	index := start
	first, second := -1, -1

	for {
		if index + 1 > len(line) { break; }
		cur := line[index]
		if cur == ' ' {
			// Second value has already been found
			if second > 0 {
				break;
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

	name = line[first+1:second]
	value = line[second:]

	return Macro {command, name, value}
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
	var macros[]Macro
	var macro_lines[]int

	file, err := os.Open(filename)
	if err != nil {
		fatal("Error opening file")
	}

	defer file.Close()

	line_num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()

		// Add macros
		if len(current) < 1 { continue; }
		if start := check_start(current); start != -1 {
			command := lex_command(current)
			if command == None { continue; }
			macro := tokenize(current, command, start)
			macros = append(macros, macro)
			macro_lines = append(macro_lines, line_num + 1)
		}

		// Use macros
		for _, mac := range macros {
			if cur_start := strings.Index(current, mac.name); cur_start != -1 {
				if !contains(macro_lines, cur_start) {
					fmt.Println(cur_start)
				}
			}
		}

		line_num += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
