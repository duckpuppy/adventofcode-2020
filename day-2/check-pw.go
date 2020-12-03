package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	old_count := 0
	new_count := 0
	for _, fmt := range lines {
		old, new, err := checkPw(fmt)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		if old {
			old_count++
		}
		if new {
			new_count++
		}
	}

	fmt.Printf("Total lines: %d\n", len(lines))
	fmt.Printf("%d passwords are valid in the old algorithm\n", old_count)
	fmt.Printf("%d passwords are valid in the new algorithm\n", new_count)
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

	// Return the slice sorted in reverse order
	sort.Slice(lines, func(i, j int) bool { return lines[j] < lines[i] })
	return lines, err
}

func checkPw(pass_fmt string) (bool, bool, error) {
	var parts = strings.Split(pass_fmt, ": ")
	format := parts[0]
	pass := parts[1]

	old_valid := false
	new_valid := false

	re := regexp.MustCompile(`(?P<low>\d+)-(?P<high>\d+)\s(?P<char>\w)`)
	if re.MatchString(format) {
		matches := re.FindStringSubmatch(format)
		low, _ := strconv.ParseInt(matches[re.SubexpIndex("low")], 10, 32)
		high, _ := strconv.ParseInt(matches[re.SubexpIndex("high")], 10, 32)
		char := matches[re.SubexpIndex("char")]
		num := strings.Count(pass, char)

		fmt.Printf("Checking %s for validity:\n", pass_fmt)
		fmt.Print("    OLD FORMAT: ")
		if low <= int64(num) && int64(num) <= high {
			fmt.Printf("VALID\n")
			old_valid = true
		} else {
			fmt.Printf("INVALID: %s should appear %d to %d times, but appears %d\n", char, low, high, num)
		}

		fmt.Print("    NEW FORMAT: ")
		testStr := string(pass[low-1]) + string(pass[high-1])
		if (string(pass[low-1]) == char) != (string(pass[high-1]) == char) {
			fmt.Printf("VALID: %s\n", testStr)
			new_valid = true
		} else {
			fmt.Printf("INVALID: %s should appear either location %d or %d, but does not (or is in both): %s\n", char, low, high, testStr)
		}
	} else {
		return false, false, fmt.Errorf("Invalid password format string: %s", pass_fmt)
	}

	return old_valid, new_valid, nil
}
