package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to check if a design can be formed using the available patterns
func isDesignPossible(patterns []string, design string) bool {
	// Dynamic Programming to check if the design can be formed
	dp := make([]bool, len(design)+1)
	dp[0] = true // Base case: Empty design can always be formed

	for i := 1; i <= len(design); i++ {
		for _, pattern := range patterns {
			if len(pattern) <= i && dp[i-len(pattern)] {
				if design[i-len(pattern):i] == pattern {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(design)]
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read towel patterns
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")

	// Skip the blank line
	scanner.Scan()

	designs := []string{}

	// Read designs
	for scanner.Scan() {
		design := scanner.Text()
		if design != "" {
			designs = append(designs, design)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Count the number of possible designs
	possibleCount := 0

	for _, design := range designs {
		if isDesignPossible(patterns, design) {
			possibleCount++
		}
	}

	fmt.Printf("Number of possible designs: %d\n", possibleCount)
}
