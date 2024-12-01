package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Prompt user for input file name
	var filename string
	fmt.Print("Enter the name of the input file: ")
	fmt.Scan(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize slices for the two lists
	var list1, list2 []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			// Convert strings to integers
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate the difference list and sum
	diffList := make([]int, len(list1))
	sum := 0
	for i := range list1 {
		diffList[i] = int(math.Abs(float64(list1[i] - list2[i])))
		sum += diffList[i]
	}

	// Output the results
	fmt.Println("Sorted first list:", list1)
	fmt.Println("Sorted second list:", list2)
	fmt.Println("Difference list:", diffList)
	fmt.Printf("Sum of the difference list: %d\n", sum)
}
