package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
	defer file.Close()

	var passport string
	valid_passports_all_fields := 0
	valid_passports_data := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(SplitAtEmptyLine())
	for scanner.Scan() {
		passport = strings.Join(strings.Split(scanner.Text(), "\n"), " ")
		fmt.Printf("passport %s: ", passport)
		if valid, _ := isValidPassport(passport, false); valid {
			valid_passports_all_fields++
		}
		if valid, errors := isValidPassport(passport, true); valid {
			valid_passports_data++
			fmt.Println("VALID")
		} else {
			fmt.Printf("INVALID\n%s\n", strings.Join(errors, "\n"))
		}
	}

	fmt.Printf("++++++++++++++++++++++++++++++++\nValid passports -- No missing fields: %d, valid data: %d\n", valid_passports_all_fields, valid_passports_data)
}

func isValidPassport(passport string, validateFields bool) (bool, []string) {
	passport_map := parsePassport(passport)
	if valid, validation_errors := checkRequiredFields(passport_map, validateFields); valid {
		return true, nil
	} else {
		return false, validation_errors
	}
}

func getRequiredFields(includeOptional bool) []string {
	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	if includeOptional {
		return append(required_fields, "cid")
	}
	return required_fields
}

func checkRequiredFields(passport map[string]string, validateFields bool) (bool, []string) {
	var validation_errors []string
	valid := true

	for _, field := range getRequiredFields(false) {
		if val, ok := passport[field]; !ok {
			validation_errors = append(validation_errors, fmt.Sprintf("missing %s", field))
			valid = false
		} else {
			if validateFields {
				if field_valid := validateField(field, val); field_valid == false {
					valid = false
					validation_errors = append(validation_errors, fmt.Sprintf("%s is invalid", field))
				}
			}
		}
	}
	return valid, validation_errors
}

func validateField(field, value string) bool {
	switch field {
	case "byr":
		return checkValidYear(value, 1920, 2002)
	case "iyr":
		return checkValidYear(value, 2010, 2020)
	case "eyr":
		return checkValidYear(value, 2020, 2030)
	case "hgt":
		re := regexp.MustCompile(`^(?P<num>\d+)(?P<unit>in|cm)$`)
		if matched := re.MatchString(value); matched {
			matches := re.FindStringSubmatch(value)
			val, _ := strconv.Atoi(matches[re.SubexpIndex("num")])
			unit := matches[re.SubexpIndex("unit")]
			switch unit {
			case "cm":
				if val >= 150 && val <= 193 {
					return true
				}
			case "in":
				if val >= 59 && val <= 76 {
					return true
				}
			}
		}
		return false
	case "hcl":
		re := regexp.MustCompile(`^#[0-9a-f]{6}`)
		return re.MatchString(value)
	case "ecl":
		valid := false
		for _, val := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if value == val {
				valid = true
			}
		}
		return valid
	case "pid":
		re := regexp.MustCompile(`^[0-9]{9}$`)
		return re.MatchString(value)
	case "cid":
		return true
	}
	return false
}

func checkValidYear(value string, min, max int) bool {
	if len(value) != 4 {
		return false
	}
	if val, err := strconv.Atoi(value); err != nil {
		return false
	} else {
		if val <= max && val >= min {
			return true
		} else {
			return false
		}
	}
}

func parsePassport(passport string) map[string]string {
	passport_map := make(map[string]string)
	for _, field := range strings.Split(passport, " ") {
		parts := strings.Split(field, ":")
		passport_map[parts[0]] = parts[1]
	}
	// fmt.Printf("%v", passport_map)

	return passport_map
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
