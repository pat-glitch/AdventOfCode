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
	// Check consecutive differences
	for i := 0; i < len(row)-1; i++ {
		diff := int(math.Abs(float64(row[i+1] - row[i])))
		if diff < 1 || diff > 3 {
			return false // Unsafe condition
		}
	}

	// Check monotonicity
	increasing, decreasing := true, true
	for i := 0; i < len(row)-1; i++ {
		if row[i] < row[i+1] {
			decreasing = false
		} else if row[i] > row[i+1] {
			increasing = false
		}
	}

	if !increasing && !decreasing {
		return false // Not monotonic, hence unsafe
	}

	return true // Safe if all conditions pass
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

		// Check if the row is safe
		if isSafeRow(row) {
			safeCount++
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
