package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// 1. Read the whole file at once
	// Since the input is one single long line(or a few comma-separated ones)
	// it's easier to read the whole thing into a byte slice first
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert bytes to string and clean up any whitespace(newlines)
	data := strings.TrimSpace(string(content))

	// 2. Split the data by commas to get ranges like "11-22"
	// strings.Split creates a "slice"(array) of substrings
	rangeStrings := strings.Split(data, ",")

	totalSum := 0

	// 3. Loop through each range string
	for _, rng := range rangeStrings {
		// rng is something like "95-115"
		// Split by hyphen to get the two numbersS
		bounds := strings.Split(rng, "-")

		// Convert the strings to numbers
		startNum, _ := strconv.Atoi(bounds[0])
		endNum, _ := strconv.Atoi(bounds[1])

		// 4. Iterate through the range and sum the numberstam
		for n := startNum; n <= endNum; n++ {
			if isInvalidID(n) {
				totalSum += n
			}
		}
	}
	fmt.Printf("Total Sum of Invalid IDs: %d\n", totalSum)
	fmt.Printf("Finished in %s\n", time.Since(start))
}

// isInvalidID checks if the number fits the "pattern repeated twice" rule
// examples: 55(invalid), 1212(invalid), 123123(invalid), 1234(valid), 101(invalid)
func isInvalidID(n int) bool {
	// Convert the number to string to inspect digits easily
	s := strconv.Itoa(n)

	//Rule 1: Must be even length
	// A number like 101(length 3) cannot be two equal halves
	if len(s)%2 != 0 {
		return false
	}

	//Rule 2: Check if first half equals second half
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	// Rule 3: Check if they match
	return firstHalf == secondHalf
}
