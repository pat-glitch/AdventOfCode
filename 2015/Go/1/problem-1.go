package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("inputdata.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, World!")

	instructions := string(data)
	fmt.Println("Instructions loaded successfully")

	// Part-1: Find the floor Santa ends up on
	floor := 0
	// Part-2: Find the position of the first character that causes Santa to enter the basement
	for _, char := range instructions {
		if char == '(' {
			floor++ // Go up a floor
		} else if char == ')' {
			floor-- // Go down one floor
		}
	}
	fmt.Printf("Santa ends up on floor: %d\n", floor)

	// Part-2: Find the position of the first character that causes Santa to enter the basement
	floor = 0
	position := 0
	for i, char := range instructions {
		if char == '(' {
			floor++ // Go up a floor
		} else if char == ')' {
			floor-- // Go down one
		}

		// Check if Santa has entered the basement
		if floor == -1 {
			position = i + 1
			break
		}
	}
	fmt.Printf("Santa enters the basement at position: %d\n", position)
}
