package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Function to interpret and perform multiplication
func interpretMult(mulStr string) int {
	// Regular expression to extract numbers from mul(x,y)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	match := re.FindStringSubmatch(mulStr)

	// Convert the captured numbers to integers
	if len(match) == 3 {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		return x * y
	}
	return 0
}

func p2() {
	var total int = 0
	var do_yes bool = true // By default, mul instructions are enabled
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Regular expression to match `do()`, `don't()`, and `mul(x,y)`
	re := regexp.MustCompile("(do\\(\\))|(mul\\([0-9]+,[0-9]+\\))|(don't\\(\\))")

	// Reading the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all matching substrings
		found := re.FindAllString(line, -1)

		// Iterate through all matches in the line
		for i := 0; i < len(found); i++ {
			if len(found[i]) == 4 { // "do()" instruction
				do_yes = true
			} else if found[i][:5] == "don't" { // "don't()" instruction
				do_yes = false
			} else { // This is a "mul(x,y)" instruction
				if do_yes {
					total += interpretMult(found[i]) // Add result of multiplication
				}
			}
		}
	}

	// Print the total sum of all valid multiplications
	fmt.Println(total)
}

func main() {
	p2()
}
