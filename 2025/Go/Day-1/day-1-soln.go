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
	totalHits := 0

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
		amount, _ := strconv.Atoi(amountStr)

		// 1. Count "Full Circle" hits
		// A full circle is completed every 100 steps
		fullCircles := amount / modSize
		totalHits += fullCircles

		// 2. Calculate the remaining movement
		remainder := amount % modSize

		// 3. Check if the remainder is enough to cross 0
		if direction == 'R' {
			// To hit 0 going right, we need to reach 100.
			// Distance to next 0 = 100 - currentPos
			// Example: At 90, we need 10 steps.
			// Special Case: If we are AT0, we need 100 steps to hit 0 again.
			disttoZero := (modSize - currentPos)

			if remainder >= disttoZero {
				totalHits++
			}

			// Update current position
			currentPos = (currentPos + remainder) % modSize

		} else if direction == 'L' {
			// To hit 0 going Left, we simply need to reach 0.
			// Distance to next  0 = currentPos
			// Example: At 10, we need 10 steps.
			// Special Case: If we are AT0, we need 100 steps to hit 0 again.
			disttoZero := currentPos
			if currentPos == 0 {
				disttoZero = modSize
			}

			// If we move enough left to consume the distance, we hit 0.
			if remainder >= disttoZero {
				totalHits++
			}

			// Update current position
			currentPos = (currentPos - remainder) % modSize
			if currentPos < 0 {
				currentPos += modSize
			}
		}

	}
	fmt.Printf("The Part 2 password is: %d\n", totalHits)
}
