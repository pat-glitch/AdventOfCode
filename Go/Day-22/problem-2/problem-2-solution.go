package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PRUNER = (1 << 24) - 1
)

// mixWithShift implements the mixing operation
func mixWithShift(n int64, shift int) int64 {
	if shift > 0 {
		return n ^ (n << uint(shift))
	}
	return n ^ (n >> uint(-shift))
}

// prune keeps only the last 24 bits
func prune(n int64) int64 {
	return n & PRUNER
}

// difference calculates consecutive differences in a slice
func difference(prices []int) []int {
	diffs := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		diffs[i-1] = prices[i] - prices[i-1]
	}
	return diffs
}

// windowed creates sliding windows of size 4
func windowed(arr []int, size int) [][4]int {
	if len(arr) < size {
		return nil
	}
	result := make([][4]int, len(arr)-size+1)
	for i := 0; i <= len(arr)-size; i++ {
		var window [4]int
		copy(window[:], arr[i:i+size])
		result[i] = window
	}
	return result
}

func main() {
	// Read input from file
	data, err := os.ReadFile("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Split into lines and convert to numbers, handling both \n and \r\n
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(content), "\n")
	nums := make([]int64, len(lines))
	for i, line := range lines {
		// Clean any remaining \r if present
		line = strings.TrimSpace(line)
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing number on line %d: %v\n", i+1, err)
			return
		}
		nums[i] = n
	}

	var part1 int64
	priceInfo := make([]map[[4]int]int, len(nums))

	// Process each number
	for i, n := range nums {
		prices := make([]int, 2001) // 0 + 2000 iterations
		prices[0] = int(n % 10)

		current := n
		for j := 0; j < 2000; j++ {
			current = prune(mixWithShift(current, 6))
			current = prune(mixWithShift(current, -5))
			current = prune(mixWithShift(current, 11))
			prices[j+1] = int(current % 10)
		}

		part1 += current

		// Calculate price differences
		priceDiff := difference(prices)
		priceDict := make(map[[4]int]int)

		// Create windows and store first occurrences
		windows := windowed(priceDiff, 4)
		for j, window := range windows {
			if _, exists := priceDict[window]; !exists {
				priceDict[window] = prices[j+4]
			}
		}

		priceInfo[i] = priceDict
	}

	fmt.Println("Part 1:", part1)

	// Find part 2
	var part2 int

	// Collect all unique windows
	windowSet := make(map[[4]int]bool)
	for _, pi := range priceInfo {
		for k := range pi {
			windowSet[k] = true
		}
	}

	// Find maximum sum
	for window := range windowSet {
		sum := 0
		for _, pi := range priceInfo {
			if val, ok := pi[window]; ok {
				sum += val
			}
		}
		if sum > part2 {
			part2 = sum
		}
	}

	fmt.Println("Part 2:", part2)
}
