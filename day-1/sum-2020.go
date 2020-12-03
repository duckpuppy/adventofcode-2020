package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	// Check for the 2 and 3 number sums
	for i, n := range lines {
		for _, m := range lines[i+1:] {
			if n+m == 2020 {
				fmt.Printf("Found 2 numbers that sum to 2020: %d+%d: The product is %d\n", n, m, n*m)
			}
			for _, o := range lines[i+2:] {
				if n+m+o == 2020 {
					fmt.Printf("Found 3 numbers that sum to 2020: %d+%d+%d: The product is %d\n", n, m, o, n*m*o)
				}
			}
		}
	}
}

func readLines(filename string) (lines []int64, err error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file into a slice, converting each line to an int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, val)
	}

	// Return the slice sorted in reverse order
	sort.Slice(lines, func(i, j int) bool { return lines[j] < lines[i] })
	return lines, err
}
