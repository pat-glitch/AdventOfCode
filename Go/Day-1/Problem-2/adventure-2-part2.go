package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateSimilarityScore(filename string) (int64, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return -1, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Map to store frequencies of numbers in the right list
	frequency := make(map[int]int)
	var totalScore int64

	// First pass: Count occurrences of numbers in the right list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			rightNumber, _ := strconv.Atoi(numbers[1])
			frequency[rightNumber]++
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return -1, fmt.Errorf("error reading file: %w", err)
	}

	// Rewind the file
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	// Second pass: Calculate similarity score
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			leftNumber, _ := strconv.Atoi(numbers[0])
			totalScore += int64(leftNumber) * int64(frequency[leftNumber])
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return -1, fmt.Errorf("error reading file: %w", err)
	}

	return totalScore, nil
}

func main() {
	// Prompt user for input file name
	var filename string
	fmt.Print("Enter the name of the input file: ")
	fmt.Scan(&filename)

	// Calculate similarity score
	similarityScore, err := calculateSimilarityScore(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total similarity score: %d\n", similarityScore)
}
