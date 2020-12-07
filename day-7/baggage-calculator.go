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

type Bag struct {
	color        string
	contained_by []*Bag
	contains     map[string]int
}

func NewBag(color string) *Bag {
	return &Bag{color: color, contained_by: []*Bag{}, contains: map[string]int{}}
}

type BagMap map[string]*Bag

// This is so clunky - gotta be a better way to do this
func (b BagMap) ParseBagRule(rule string) {
	re := regexp.MustCompile(`^(?P<color>([a-z]\s?)+)\sbags\scontain\s(?P<contains>.*)\.$`)
	if re.MatchString(rule) {
		matches := re.FindStringSubmatch(rule)
		color := matches[re.SubexpIndex("color")]
		contains := strings.Split(matches[re.SubexpIndex("contains")], ", ")
		// fmt.Printf("Color: %s\n", color)
		// Create the bag if it doesn't exist
		if _, exists := b[color]; exists == false {
			b[color] = NewBag(color)
		}
		bag_re := regexp.MustCompile(`^(?P<number>\d+)\s(?P<color>([a-z]\s?)+)\sbag(s)?`)
		for _, c := range contains {
			if c == "no other bags" {
				return
			}

			if bag_re.MatchString(c) {
				c_matches := bag_re.FindStringSubmatch(c)
				c_color := c_matches[bag_re.SubexpIndex("color")]
				c_num, _ := strconv.Atoi(c_matches[bag_re.SubexpIndex("number")])
				if _, exists := b[c_color]; exists == false {
					b[c_color] = NewBag(c_color)
				}
				// fmt.Printf("\tCan hold: %s\n", c_color)
				b[color].contains[c_color] = c_num
				b[c_color].contained_by = append(b[c_color].contained_by, b[color])
			}
		}
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var bags BagMap = make(BagMap)

	for scanner.Scan() {
		bags.ParseBagRule(scanner.Text())
	}

	// fmt.Printf("%v", getContainingColors(bags["shiny gold"]))
	fmt.Printf("%d bag colors can hold shiny gold\n", len(bags.GetContainingColors("shiny gold")))
	fmt.Printf("%d other bags will be in shiny gold\n", bags.GetHeldBags("shiny gold"))
}

func (b BagMap) GetHeldBags(bag_color string) int {
	direct_bags := 0
	total_bags := 0
	// fmt.Printf("%s bags hold:\n", bag_color)
	for color, number := range b[bag_color].contains {
		// fmt.Printf("\t%d %s bags\n", number, color)
		total_bags = total_bags + number + (number * b.GetHeldBags(color))
		direct_bags += number
	}

	// fmt.Printf("%s bags directly hold %d other bags\n\n", bag_color, direct_bags)
	return total_bags
}

func (b BagMap) GetContainingColors(bag_color string) []string {
	colors := []string{}
	// fmt.Printf("%s can be contained by %d other colors\n", bag_color, len(b[bag_color].contained_by))
	for _, bag := range b[bag_color].contained_by {
		// fmt.Printf("%s can be contained by %s\n", bag.color, b.color)
		colors = append(colors, bag.color)
		colors = append(colors, b.GetContainingColors(bag.color)...)
	}

	// fmt.Printf("%v\n", colors)
	return unique(colors)
}

func unique(list []string) []string {
	unique_map := map[string]bool{}
	unique_list := []string{}

	for v := range list {
		if unique_map[list[v]] != true {
			unique_map[list[v]] = true
			unique_list = append(unique_list, list[v])
		}
	}

	return unique_list
}
