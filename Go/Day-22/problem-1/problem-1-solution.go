package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Constants
const Modulus = 16777216
const Iterations = 2000

// Function to calculate the next secret number
func nextSecret(secret int) int {
	// Step 1: Multiply by 64, mix, and prune
	secret ^= (secret * 64) % Modulus
	secret %= Modulus

	// Step 2: Divide by 32, mix, and prune
	secret ^= (secret / 32) % Modulus
	secret %= Modulus

	// Step 3: Multiply by 2048, mix, and prune
	secret ^= (secret * 2048) % Modulus
	secret %= Modulus

	return secret
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read initial secret numbers
	var secrets []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to parse number: %v", err)
		}
		secrets = append(secrets, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Simulate and sum the 2000th secret number for each buyer
	total := 0
	for _, secret := range secrets {
		for i := 0; i < Iterations; i++ {
			secret = nextSecret(secret)
		}
		total += secret
	}

	// Output the result
	fmt.Println("Sum of 2000th secret numbers:", total)
}
