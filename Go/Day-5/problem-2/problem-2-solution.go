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

// Function to check if an update is in the correct order
func isUpdateValid(line []int, rules []Rule) bool {
	for _, rule := range rules {
		aPos, bPos := -1, -1

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

// Function to perform topological sorting for a single update
func topologicalSort(line []int, rules []Rule) []int {
	// Create a map to store the positions of elements and a boolean array to track if they are sorted
	sorted := false
	for !sorted {
		sorted = true
		for _, rule := range rules {
			// If the current rule is violated, swap the positions of the pages
			for i := 0; i < len(line)-1; i++ {
				for j := i + 1; j < len(line); j++ {
					if line[i] == rule.b && line[j] == rule.a {
						// Swap the pages
						line[i], line[j] = line[j], line[i]
						sorted = false
					}
				}
			}
		}
	}
	return line
}

// Function to find the middle element of a valid update
func getMiddleElement(line []int) int {
	return line[len(line)/2] // Use integer division for consistent results
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

	// Reading rules from the input file
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

		// If the update is invalid, sort it
		if !isUpdateValid(lineNumbers, rules) {
			// Sort the update using the rules
			sortedLine := topologicalSort(lineNumbers, rules)
			// Get the middle element of the sorted update and add to sum
			midpoint := getMiddleElement(sortedLine)
			sum += midpoint
		}
	}

	// Output the sum of middle elements of corrected (invalid) updates
	fmt.Printf("Sum of middle elements of corrected updates: %d\n", sum)
}
