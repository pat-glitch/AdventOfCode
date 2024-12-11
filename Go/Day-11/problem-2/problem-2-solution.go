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

	// Perform 25 blinks
	for blink := 0; blink < 75; blink++ {
		var newStones []int
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if isEvenDigits(stone) {
				left, right := splitDigits(stone)
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}

	fmt.Println("Number of stones after 75 blinks:", len(stones))
}

// Check if a number has an even number of digits
func isEvenDigits(num int) bool {
	digits := 0
	for num > 0 {
		digits++
		num /= 10
	}
	return digits%2 == 0
}

// Split the digits of a number in half
func splitDigits(num int) (int, int) {
	digits := []int{}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}
	mid := len(digits) / 2

	left := 0
	for i := 0; i < mid; i++ {
		left = left*10 + digits[i]
	}

	right := 0
	for i := mid; i < len(digits); i++ {
		right = right*10 + digits[i]
	}

	return left, right
}
