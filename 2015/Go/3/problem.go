package main

import (
	"bufio"
	"fmt"
	"os"
)

// Position represents a coordinate on the grid
type Position struct {
	x, y int
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the input directions
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	// Track visited houses using a map
	visitedHouses := make(map[Position]bool)

	// Starting positions for Santa and Robo-Santa
	santaPosition := Position{0, 0}
	roboPosition := Position{0, 0}
	visitedHouses[santaPosition] = true

	// Process each direction alternately between Santa and Robo-Santa
	for i, direction := range directions {
		if i%2 == 0 {
			santaPosition = move(santaPosition, direction)
			visitedHouses[santaPosition] = true
		} else {
			roboPosition = move(roboPosition, direction)
			visitedHouses[roboPosition] = true
		}
	}

	// Count the number of unique houses visited
	numberOfHouses := len(visitedHouses)

	fmt.Println("Number of houses that receive at least one present:", numberOfHouses)
}

// move updates the position based on the given direction
func move(pos Position, direction rune) Position {
	switch direction {
	case '^':
		pos.y++
	case 'v':
		pos.y--
	case '>':
		pos.x++
	case '<':
		pos.x--
	}
	return pos
}
