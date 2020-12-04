package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	slopes := [][]int{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}

	tree_product := 1 // Initialize to 1 for indentity product on first loop
	for _, pair := range slopes {
		trees := checkSlope(lines, pair[0], pair[1])
		fmt.Printf("Trees encountered for slope %d down and %d right: %d\n", pair[1], pair[0], trees)
		tree_product *= trees
	}

	fmt.Printf("Product of all trees encountered on all slopes: %d\n", tree_product)
}

func checkSlope(slopemap []string, x_step, y_step int) int {
	x := 0
	trees := 0

	for row := 0; row < len(slopemap); row += y_step {
		space := string(slopemap[row][x])
		// fmt.Printf("Checking %d, %d: %s\n", x, row, space)
		if space == "#" {
			trees++
		}
		x = (x + x_step) % len(slopemap[row])
		// fmt.Printf("Next x is: %d\n", x)
	}

	return trees
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
