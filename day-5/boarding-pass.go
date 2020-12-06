package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
	defer file.Close()

	var seat_list []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Decoding %s\n", scanner.Text())
		seat_id := decodeSeat(scanner.Text())

		seat_list = append(seat_list, seat_id)
	}
	sort.Ints(seat_list)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("Highest seat ID: %d\n", seat_list[len(seat_list)-1])
	for i := 1; i < len(seat_list)-1; i++ {
		if seat_list[i+1] != seat_list[i]+1 {
			fmt.Printf("My seat ID: %d\n", seat_list[i]+1)
		}
	}
}

func decodeSeat(seat_coords string) int {
	num_rows := 128
	seats_per_row := 8

	// Because of the input, we can make some assumptions
	upper_row := num_rows - 1
	lower_row := 0
	for c := 0; c < 7; c++ {
		switch seat_coords[c] {
		case 'F':
			upper_row = lower_row + ((upper_row - lower_row) / 2)
		case 'B':
			lower_row = upper_row - ((upper_row - lower_row) / 2)
		}
		fmt.Printf("\tFound %s: New partition %d-%d\n", string(seat_coords[c]), lower_row, upper_row)
	}

	upper_seat := seats_per_row - 1
	lower_seat := 0
	for c := 7; c < 10; c++ {
		switch seat_coords[c] {
		case 'L':
			upper_seat = lower_seat + ((upper_seat - lower_seat) / 2)
		case 'R':
			lower_seat = upper_seat - ((upper_seat - lower_seat) / 2)
		}
		fmt.Printf("\tFound %s: New partition %d-%d\n", string(seat_coords[c]), lower_seat, upper_seat)
	}
	seat_id := lower_row*8 + lower_seat
	fmt.Printf("\tSeat ID: %d\n", seat_id)
	return seat_id
}
