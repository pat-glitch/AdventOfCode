package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to calculate the wrapping paper reqiured for a single box
func calculateWrappingPaper(l, w, h int) int {
	side1 := l * w
	side2 := w * h
	side3 := h * l

	surfaceArea := 2*side1 + 2*side2 + 2*side3
	slack := min(side1, min(side2, side3))

	return surfaceArea + slack
}

// Function to calculate the ribbon required for a single box
func calculateRibbon(l, w, h int) int {
	// Smallest perimeter
	side1 := 2 * (l + w)
	side2 := 2 * (w + h)
	side3 := 2 * (h + l)

	smallestPerimeter := min(side1, min(side2, side3))

	// Volume for the bow
	volume := l * w * h

	return smallestPerimeter + volume
}

// Function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalWrappingPaper := 0
	totalRibbon := 0

	// Read dimension from the file
	for scanner.Scan() {
		dimension := scanner.Text()
		dims := strings.Split(dimension, "x")

		// Parse dimensions to integers
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])

		// Calculate the wrapping paper required for the box
		totalWrappingPaper += calculateWrappingPaper(l, w, h)

		// Calculate the ribbon required for the box
		totalRibbon += calculateRibbon(l, w, h)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	fmt.Printf("Total wrapping paper required: %d\n", totalWrappingPaper)
	fmt.Printf("Total ribbon required: %d\n", totalRibbon)
}
