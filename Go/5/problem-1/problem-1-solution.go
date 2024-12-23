package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	a, b int
}

// Function to check if the update is valid according to the rules
func isUpdateValid(line []int, rules []Rule) bool {
	for _, rule := range rules {
		aPos, bPos := -1, 1000

		// Find positions of both a and b in the line
		for i, page := range line {
			if page == rule.a {
				aPos = i
			}
			if page == rule.b {
				bPos = i
			}
		}

		// If both a and b are present and a comes after b, it's invalid
		if aPos != -1 && bPos != -1 && aPos > bPos {
			return false
		}
	}
	return true
}

func main() {
	// Open the input file
	fileName := "inputdata.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var rules []Rule

	// Reading rules until we encounter an empty line (or some separator)
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}

		if line == "" { // If we encounter an empty line, stop reading rules
			break
		}

		// Parse the rule in the form of "a|b"
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule{a, b})
		}
	}

	var sum int
	// Now process the updates
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}

		// Split the update line into individual page numbers
		parts := strings.Split(line, ",")
		lineNumbers := []int{}
		for _, part := range parts {
			page, _ := strconv.Atoi(part)
			lineNumbers = append(lineNumbers, page)
		}

		// Find the middle element
		midpoint := lineNumbers[len(lineNumbers)/2]

		// Check if the update is valid and add the middle element if valid
		if isUpdateValid(lineNumbers, rules) {
			sum += midpoint
		}
	}

	// Output the sum of middle elements of valid updates
	fmt.Printf("Sum of middle elements of valid updates: %d\n", sum)
}
