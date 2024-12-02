package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Function to check if a row is "safe"
func isSafeRow(row []int) bool {
	if len(row) < 2 {
		return true // Rows with fewer than 2 numbers are automatically safe
	}

	increasing, decreasing := true, true

	for i := 0; i < len(row)-1; i++ {
		diff := int(math.Abs(float64(row[i+1] - row[i])))
		if diff < 1 || diff > 3 {
			fmt.Printf("Unsafe due to difference at index %d: %d to %d\n", i, row[i], row[i+1])
			return false // Unsafe: differences not in [1, 3]
		}
		if row[i] < row[i+1] {
			decreasing = false
		} else if row[i] > row[i+1] {
			increasing = false
		}
	}

	// Unsafe if not monotonic
	if !(increasing || decreasing) {
		fmt.Println("Unsafe due to non-monotonic order")
	}
	return increasing || decreasing
}

// Function to check if the row can be made "safe" by removing one number
func canBeSafeByRemovingOne(row []int) bool {
	for i := 0; i < len(row); i++ {
		// Create a temporary slice excluding the i-th number
		temp := append(row[:i], row[i+1:]...)
		if isSafeRow(temp) {
			fmt.Printf("Row is safe by removing number %d\n", row[i])
			return true
		}
	}
	return false
}

// Function to process the file and determine "safe" rows
func processFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount, unsafeCount := 0, 0

	for scanner.Scan() {
		// Read the line and split it into numbers
		line := scanner.Text()
		tokens := strings.Fields(line)
		row := make([]int, len(tokens))

		// Convert tokens to integers
		for i, token := range tokens {
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
			row[i] = num
		}

		// Debug print row
		fmt.Printf("Processing row: %v\n", row)

		// Check if the row is safe
		if isSafeRow(row) {
			safeCount++
		} else if canBeSafeByRemovingOne(row) {
			safeCount++ // Safe after removing one number
		} else {
			unsafeCount++
		}
	}

	// Output results
	fmt.Printf("Total safe rows: %d\n", safeCount)
	fmt.Printf("Total unsafe rows: %d\n", unsafeCount)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	var filename string
	fmt.Print("Enter the name of the input file: ")
	fmt.Scan(&filename)

	processFile(filename)
}
