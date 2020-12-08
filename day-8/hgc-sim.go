package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	program, err := readLines("input")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	acc, err := Run(program)
	if err != nil {
		log.Printf("ERROR: %s", err)
	}
	log.Printf("Accumulator: %d\n", acc)

	log.Printf("++++++++++++++++++++++++++++++++++++++++++++++++")
	log.Println("Fixing Loop:")
	acc = FixLoop(program)
	log.Printf("Fixed: Result is %d\n", acc)
}

func Run(program []string) (int, error) {
	acc := 0 // Global accumulator
	pc := 0  // Program counter

	previous_pc := map[int]bool{}

	for {
		// Exit if the pc is outside the program
		if pc >= len(program) || pc < 0 {
			break
		}

		if previous_pc[pc] == true { // We've visited this line before
			return acc, fmt.Errorf("loop encountered at %d", pc)
		}

		previous_pc[pc] = true // Implement a simple visitor pattern
		op, val := parseStatement(program[pc])
		switch op {
		case "acc":
			acc += val
			pc++
		case "jmp":
			pc += val
		case "nop":
			pc++
		}
	}
	return acc, nil
}

func FixLoop(program []string) int {
	for index, line := range program {
		op, _ := parseStatement(line)
		if op == "jmp" || op == "nop" {
			program1 := make([]string, len(program))
			copy(program1, program)
			switch op {
			case "jmp":
				program1[index] = strings.Replace(program1[index], "jmp", "nop", 1)
			case "nop":
				program1[index] = strings.Replace(program1[index], "nop", "jmp", 1)
			}
			acc, err := Run(program1)
			if err != nil {
				log.Printf("INFO: Fixing line %d (%s) did not help\n", index, line)
				continue
			}
			log.Printf("INFO: Fixing line %d (%s) worked!\n", index, line)
			return acc
		}
	}
	return -1
}

func parseStatement(statement string) (string, int) {
	re := regexp.MustCompile(`^(?P<op>jmp|acc|nop)\s(?P<val>[+-]\d+)$`)
	if re.MatchString(statement) {
		matches := re.FindStringSubmatch(statement)
		op := matches[re.SubexpIndex("op")]
		val, _ := strconv.Atoi(matches[re.SubexpIndex("val")])
		return op, val
	}

	return "nop", 0 // Arbirtrarily decided that a syntax error is a noop
}

func readLines(filename string) (lines []string, err error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file into a slice of lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, err
}
