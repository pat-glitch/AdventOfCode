package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

// ButtonMove represents the movement of a button
type ButtonMove struct {
	Name      string
	TokenCost int
	XMove     int
	YMove     int
}

// Machine represents a single claw machine
type Machine struct {
	ButtonA ButtonMove
	ButtonB ButtonMove
	PrizeX  int
	PrizeY  int
}

// calculateMinTokens finds the minimum tokens to reach the prize
func calculateMinTokens(machine Machine) (int, error) {
	minTokens := math.MaxInt32

	// Use GCD to optimize the search space
	gcdX := gcd(machine.ButtonA.XMove, machine.ButtonB.XMove)
	gcdY := gcd(machine.ButtonA.YMove, machine.ButtonB.YMove)

	// Determine maximum reasonable search iterations using GCD
	maxIterations := (machine.PrizeX/gcdX + machine.PrizeY/gcdY) * 2

	for a := 0; a <= maxIterations; a++ {
		for b := 0; b <= maxIterations; b++ {
			// Calculate total X and Y movements
			totalX := a*machine.ButtonA.XMove + b*machine.ButtonB.XMove
			totalY := a*machine.ButtonA.YMove + b*machine.ButtonB.YMove

			// Check if we've reached the prize exactly
			if totalX == machine.PrizeX && totalY == machine.PrizeY {
				// Calculate total tokens spent
				tokens := a*machine.ButtonA.TokenCost + b*machine.ButtonB.TokenCost

				// Update minimum tokens if found
				if tokens < minTokens {
					minTokens = tokens
				}
			}
		}
	}

	// If no solution found
	if minTokens == math.MaxInt32 {
		return 0, fmt.Errorf("no solution found for machine with prize X=%d, Y=%d", machine.PrizeX, machine.PrizeY)
	}

	return minTokens, nil
}

// gcd calculates the Greatest Common Divisor using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// parseInputFile reads the input file and returns a slice of Machines
func parseInputFile(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)

	// Regex patterns to extract values
	buttonARegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var currentMachine Machine
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if line == "" {
			continue
		}

		// Match and parse Button A
		if matches := buttonARegex.FindStringSubmatch(line); matches != nil {
			xMove, _ := strconv.Atoi(matches[1])
			yMove, _ := strconv.Atoi(matches[2])
			currentMachine.ButtonA = ButtonMove{
				Name:      "A",
				TokenCost: 3,
				XMove:     xMove,
				YMove:     yMove,
			}
			lineCount++
		}

		// Match and parse Button B
		if matches := buttonBRegex.FindStringSubmatch(line); matches != nil {
			xMove, _ := strconv.Atoi(matches[1])
			yMove, _ := strconv.Atoi(matches[2])
			currentMachine.ButtonB = ButtonMove{
				Name:      "B",
				TokenCost: 1,
				XMove:     xMove,
				YMove:     yMove,
			}
			lineCount++
		}

		// Match and parse Prize
		if matches := prizeRegex.FindStringSubmatch(line); matches != nil {
			prizeX, _ := strconv.Atoi(matches[1])
			prizeY, _ := strconv.Atoi(matches[2])
			currentMachine.PrizeX = prizeX
			currentMachine.PrizeY = prizeY
			lineCount++
		}

		// When we have processed all details for a machine, add it to machines
		if lineCount == 3 {
			machines = append(machines, currentMachine)
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

func main() {
	// Read machines from input file
	machines, err := parseInputFile("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Calculate minimum tokens for each machine
	totalTokens := 0
	for i, machine := range machines {
		tokens, err := calculateMinTokens(machine)
		if err != nil {
			fmt.Printf("Machine %d: %v\n", i+1, err)
			continue
		}
		fmt.Printf("Machine %d - Minimum tokens: %d\n", i+1, tokens)
		totalTokens += tokens
	}

	fmt.Printf("\nTotal minimum tokens for all machines: %d\n", totalTokens)
}
