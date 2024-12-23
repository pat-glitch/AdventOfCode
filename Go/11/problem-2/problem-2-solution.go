package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check if a number has an even number of digits
func isEvenDigits(num int) bool {
	digits := 0
	for num > 0 {
		digits++
		num /= 10
	}
	return digits%2 == 0
}

// Split the digits of a number into two parts
func splitDigits(num int) (int, int) {
	// Find the number of digits in num
	digits := 0
	temp := num
	for temp > 0 {
		digits++
		temp /= 10
	}

	// Split number into left and right halves
	power := 1
	for i := 0; i < digits/2; i++ {
		power *= 10
	}

	left := num / power
	right := num % power

	return left, right
}

func main() {
	// Read the initial list of stones from the file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the first line from the file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// Parse the initial stones into a map to store counts
	stoneCounts := make(map[int]int)
	stoneStrings := strings.Fields(line)
	for _, s := range stoneStrings {
		num, _ := strconv.Atoi(s)
		stoneCounts[num]++
	}

	// Perform 75 blinks
	for blink := 0; blink < 75; blink++ {
		// Create a new map to store updated stone counts after the blink
		newStoneCounts := make(map[int]int)

		for stone, count := range stoneCounts {
			if stone == 0 {
				// If the stone value is 0, we add 1
				newStoneCounts[1] += count
			} else if isEvenDigits(stone) {
				// If the number of digits in the stone is even, split it
				left, right := splitDigits(stone)
				newStoneCounts[left] += count
				newStoneCounts[right] += count
			} else {
				// If the number of digits is odd, multiply the stone by 2024
				newStoneCounts[stone*2024] += count
			}
		}

		// Update stoneCounts to the new map after this blink
		stoneCounts = newStoneCounts
	}

	// Output the total number of stones after 75 blinks
	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}

	fmt.Println("Number of stones after 75 blinks:", totalStones)
}
