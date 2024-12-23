package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Evaluate an expression left-to-right, ignoring operator precedence
func evaluateExpression(expression string) int64 {
	var result, current int64
	var lastOperator byte = '+'

	for i := 0; i < len(expression); i++ {
		char := expression[i]
		if char >= '0' && char <= '9' {
			current = current*10 + int64(char-'0')
		} else if char == '+' || char == '*' {
			if lastOperator == '+' {
				result += current
			} else if lastOperator == '*' {
				result *= current
			}
			current = 0
			lastOperator = char
		}
	}

	// Final operation
	if lastOperator == '+' {
		result += current
	} else if lastOperator == '*' {
		result *= current
	}

	return result
}

// Generate all possible operator combinations and validate the result
func checkEquation(result int64, numbers []int) bool {
	combinations := 1 << (len(numbers) - 1) // Total combinations of '+' and '*'

	for i := 0; i < combinations; i++ {
		var expression strings.Builder
		for j, num := range numbers {
			expression.WriteString(strconv.Itoa(num))
			if j < len(numbers)-1 {
				if (i & (1 << j)) != 0 {
					expression.WriteByte('*')
				} else {
					expression.WriteByte('+')
				}
			}
		}

		if evaluateExpression(expression.String()) == result {
			fmt.Printf("Valid: %s = %d\n", expression.String(), result)
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var totalSum int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		// Parse the result from the left-hand side
		result, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			continue
		}

		// Parse the numbers from the right-hand side
		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			numbers[i] = num
		}

		// Check the equation
		if checkEquation(result, numbers) {
			totalSum += result
		}
	}

	fmt.Printf("Total sum of proven true results: %d\n", totalSum)
}
