package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define the regular expression for valid mul(x,y)
	// Matches strings like mul(123,456) where x and y are 1-3 digit numbers
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Total sum of multiplications
	totalSum := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all matches in the current line
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// Convert the captured groups (numbers) to integers
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			// Add the product to the total sum
			totalSum += x * y
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the total sum of all valid mul(x,y) sequences
	fmt.Printf("Total Sum of all valid mul(x,y): %d\n", totalSum)
}
