package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	// Track visitd houses using a map
	visitedHouses := make(map[Position]bool)

	// Santa's starting position
	currentPosition := Position{0, 0}
	visitedHouses[currentPosition] = true

	// Process each direction
	for _, direction := range directions {
		switch direction {
		case '^':
			currentPosition.y++
		case 'v':
			currentPosition.y--
		case '>':
			currentPosition.x++
		case '<':
			currentPosition.x--
		}
		visitedHouses[currentPosition] = true
	}

	// Count the number of unique houses visited
	numberOfHouses := len(visitedHouses)

	// Print the result
	fmt.Println("No. of houses that recieve at least one present:", numberOfHouses)
}
