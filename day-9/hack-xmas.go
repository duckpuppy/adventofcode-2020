package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	preamble := lines[:25]
	payload := lines[25:]

	// fmt.Printf("Preamble: %#v\nPayload: %#v", preamble, payload)
	for _, val := range payload {
		if checksum(val, preamble) == false {
			ex := exploit(val, lines)
			fmt.Printf("Exploitable value found: %d\n", val)
			fmt.Printf("Exploit sum: %d\n", ex)
		}
		preamble = append(preamble[1:], val)
	}
}

func readLines(filename string) (lines []int, err error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file into a slice, converting each line to an int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, int(val))
	}

	// Return the slice sorted in reverse order
	return lines, err
}

func checksum(val int, preamble []int) bool {
	for i, n := range preamble {
		if n > val {
			continue
		}
		for _, m := range preamble[i+1:] {
			if n+m == val {
				return true
			}
		}
	}

	return false
}

func exploit(val int, payload []int) int {
	for i, n := range payload {
		fmt.Printf("Checking %d :: ", n)
		if n > val {
			fmt.Println("\tSkipping, too high")
			continue
		}
		smallest := n
		largest := n
		acc := n
		for _, m := range payload[i+1:] {
			acc += m
			if m < smallest {
				smallest = m
			} else if m > largest {
				largest = m
			}
			// fmt.Printf("\tAdding %d for a total of %d\n", m, acc)
			if acc == val {
				fmt.Printf("Found it! %d to %d\n", n, m)
				fmt.Printf("\tSmallest: %d, Largest: %d\n", smallest, largest)
				return smallest + largest
			} else if acc > val {
				fmt.Printf("Accumulator overflow: %d\n", acc)
				break
			}
		}
	}

	return -1
}
