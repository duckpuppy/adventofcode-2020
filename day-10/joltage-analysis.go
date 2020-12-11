package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	joltages, err := readLines("input")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// laptop_joltage := joltages[len(joltages)-1] + 3
	one_diff := 0
	three_diff := 0
	previous_jotage := 0
	for _, joltage := range joltages {
		// fmt.Printf("Checking difference between %d and %d: ", joltage, previous_jotage)
		if joltage-previous_jotage == 1 {
			// fmt.Print("1")
			one_diff++
			// fmt.Printf(" (%d total)\n", one_diff)
		} else if joltage-previous_jotage == 3 {
			// fmt.Print("3")
			three_diff++
			// fmt.Printf(" (%d total)\n", three_diff)
		} else {
			// fmt.Printf("Difference is not 1 or 3: %d and %d\n", joltage, previous_jotage)
		}
		previous_jotage = joltage
	}
	// Add the final 3-jolt difference between the highest adapter and the charging port
	three_diff++

	// This method takes several shortcuts based on what we know about this data set:
	// It starts with 0 (the plugin on the plane)
	// The lowest rated adapter is always valid to plug into the plane
	// The only adapter that can plug into the tablet is the highest rated one
	// There are no outliers - there's always at least one valid adapter for any adapter in the list
	// A valid adapter will be no more than 3 adapters behind in the list, as the minimum difference is 1
	// The entries in the map cache will already be populated, no need to test because we KNOW we've already visited them
	// This could absolutely been done with recursion, but I didn't - recursion vs a map for a cache, either way consumes resources
	path_cache := map[int]int{}
	for i, j := range joltages {
		fmt.Printf("Checking %d\n", j)
		if i == 0 {
			fmt.Printf("First entry: %d\n", j)
			path_cache[j] = 1
			continue
		}
		path_cache[j] = 0
		if j <= 3 {
			// This is also a valid path to 0 itself, so we count it
			path_cache[j] += 1
		}
		for _, v := range joltages[int(math.Max(float64(i-3), 0)):i] {
			if j-v <= 3 {
				fmt.Printf("\t%d is a valid node, adding path_cache (%d)\n", v, path_cache[v])
				path_cache[j] += path_cache[v]
			}
		}
		fmt.Printf("\tFinal path count: %d\n", path_cache[j])
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("Count of one-different joltages (%d) times count of three-different joltages (%d): %d\n", one_diff, three_diff, one_diff*three_diff)
	fmt.Printf("Number of valid paths: %d\n", path_cache[joltages[len(joltages)-1]])
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
		val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, int(val))
	}

	// Return the slice sorted in ascending order
	sort.Ints(lines)
	return lines, err
}

func tracePaths(joltages []int, joltage int) (bool, int) {
	fmt.Printf("Checking %d\n", joltage)
	// If we're at the last entry, either it's valid or not
	if len(joltages) == 1 {
		fmt.Println("Hit the final entry")
		if joltages[0] <= 3 {
			return true, 1
		}
		return false, 0
	}

	valid := false
	num_paths := 0
	for i, j := range joltages {
		// We assume the slice is sorted in reverse order, so we only check until we get too large a difference
		if joltage-j > 3 {
			break
		}
		// If the joltage is 2 or 3, it's valid, but there may be more paths that continue from this point
		if joltage == 2 || joltage == 3 {
			fmt.Printf("Hit valid joltage (%d), checking the rest: ", joltage)
			_, extra_paths := tracePaths(joltages[1:], joltage)
			fmt.Printf("Found %d extra paths\n", extra_paths)
			return true, 1 + extra_paths
		} else {
			if valid_paths, new_paths := tracePaths(joltages[i+1:], j); valid_paths == true {
				valid = valid || valid_paths
				num_paths += new_paths
			}
		}
	}

	return valid, num_paths
}
