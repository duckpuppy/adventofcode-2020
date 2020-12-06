package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
	defer file.Close()

	sum_all := 0
	sum_everyone := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(SplitAtEmptyLine())
	for scanner.Scan() {
		answers := strings.Split(scanner.Text(), "\n")
		unique_answers := unique(strings.Join(answers, ""))
		fmt.Printf("Checking %s\n", answers)
		sum_all += len(unique_answers)

		if len(answers) == 1 {
			sum_everyone += len(answers[0])
		} else {
			for _, char := range answers[0] {
				// For each rune in the first string
				in_all := true
				for pos, set := range answers {
					if pos == 0 {
						continue // skip the first one
					}
					fmt.Printf("\tChecking if %s contains %s: %v\n", set, string(char), strings.Contains(set, string(char)))
					in_all = in_all && strings.Contains(set, string(char))
				}
				if in_all {
					fmt.Printf("\t*** %s is in all of %v\n", string(char), answers)
					sum_everyone++
				} else {
					fmt.Printf("\t*** %s is NOT in all of %v\n", string(char), answers)
				}
			}
		}
	}
	fmt.Printf("Sum of all counts: %d\n", sum_all)
	fmt.Printf("Sum of everyone counts: %d\n", sum_everyone)
}

func SplitAtEmptyLine() func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte("\n\n")
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}

func unique(list string) []byte {
	unique_map := map[byte]bool{}
	unique_list := []byte{}

	for v := range list {
		if unique_map[list[v]] != true {
			unique_map[list[v]] = true
			unique_list = append(unique_list, list[v])
		}
	}

	return unique_list
}
