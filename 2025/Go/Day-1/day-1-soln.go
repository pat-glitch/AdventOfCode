package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize the dial state
	// The dial has numbers 0-99, so the modulus is 100
	const modSize = 100
	currentPos := 50 // The dial starts at 50
	zeroHits := 0    // Counter for the password

	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Parse the instruction
		// line[0] is the direction 'L' or 'R'
		// line[1:] is the number (string format)
		direction := line[0]
		amountStr := line[1:]

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			log.Fatalf("Invalid number in line %s: %v", line, err)
		}

		// Apply the rotation
		if direction == 'R' {
			// Right moves towards higher numbers (Addition)
			currentPos = (currentPos + amount) % modSize
		} else if direction == 'L' {
			// Left moves towards lower numbers (Subtraction)
			// Go's % operator can return negative numbers for negative operands.
			// Example: -18 % 100 = -18 in Go.
			// To fix this, we subtract, then add the modulus to ensure it's positive,
			// then mod again just to be safe.
			currentPos = (currentPos - amount) % modSize
			if currentPos < 0 {
				currentPos += modSize
			}
		}

		// Check if we landed on 0
		if currentPos == 0 {
			zeroHits++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The Actual Password is: %d\n", zeroHits)
}
