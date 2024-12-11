package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the initial list of stones from the file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// Parse the initial stones
	stoneStrings := strings.Fields(line)
	var stones []int
	for _, s := range stoneStrings {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}

	// Perform 75 blinks with aggressive optimization
	for blink := 0; blink < 75; blink++ {
		newStones := make([]int, 0, len(stones)*3) // Estimate exponential growth
		for _, stone := range stones {
			switch {
			case stone == 0:
				newStones = append(newStones, 1)
			case isEvenDigits(stone):
				left, right := splitDigitsDirect(stone)
				newStones = append(newStones, left, right)
			default:
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}

	fmt.Println("Number of stones after 75 blinks:", len(stones))
}

// Check if a number has an even number of digits directly
func isEvenDigits(num int) bool {
	digits := 0
	for num > 0 {
		digits++
		num /= 10
	}
	return digits%2 == 0
}

// Split the digits of a number directly without conversion
func splitDigitsDirect(num int) (int, int) {
	// Count total digits
	digits := 0
	temp := num
	for temp > 0 {
		digits++
		temp /= 10
	}

	// Split point
	mid := digits / 2

	// Calculate left and right parts
	left, right, divider := 0, 0, 1
	for i := 0; i < mid; i++ {
		divider *= 10
	}
	left = num / divider
	right = num % divider

	return left, right
}
