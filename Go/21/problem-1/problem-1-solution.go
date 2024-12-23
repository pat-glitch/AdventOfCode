package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Vector represents a 2D position
type Vector struct {
	x, y int
}

// Add adds two vectors
func (v Vector) Add(other Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y}
}

// Sub subtracts two vectors
func (v Vector) Sub(other Vector) Vector {
	return Vector{v.x - other.x, v.y - other.y}
}

// Global variables
var positions = map[string]Vector{
	"7": {0, 0}, "8": {0, 1}, "9": {0, 2},
	"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
	"1": {2, 0}, "2": {2, 1}, "3": {2, 2},
	"0": {3, 1}, "A": {3, 2},
	"^": {0, 1}, "a": {0, 2},
	"<": {1, 0}, "v": {1, 1}, ">": {1, 2},
}

var directions = map[string]Vector{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

// memoKey represents a key for memoization
type memoKey struct {
	str   string
	depth int
}

var memoCache = make(map[memoKey]int)

// generateMoveSet generates all possible move sequences between two positions
func generateMoveSet(start, end, avoid Vector) []string {
	delta := end.Sub(start)

	// Generate basic moves sequence
	sequence := ""
	// Vertical moves
	if delta.x < 0 {
		sequence += strings.Repeat("^", -delta.x)
	} else {
		sequence += strings.Repeat("v", delta.x)
	}
	// Horizontal moves
	if delta.y < 0 {
		sequence += strings.Repeat("<", -delta.y)
	} else {
		sequence += strings.Repeat(">", delta.y)
	}

	// Generate all permutations that don't hit avoid position
	permutations := generatePermutations(sequence)
	validMoves := make([]string, 0)

	for _, perm := range permutations {
		if !hitsAvoidPosition(start, perm, avoid) {
			validMoves = append(validMoves, perm+"a")
		}
	}

	if len(validMoves) == 0 {
		return []string{"a"}
	}
	return validMoves
}

// generatePermutations generates all unique permutations of a string
func generatePermutations(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}

	result := make(map[string]bool)
	chars := strings.Split(s, "")
	var generatePerms func([]string, []string)

	generatePerms = func(current []string, remaining []string) {
		if len(remaining) == 0 {
			result[strings.Join(current, "")] = true
			return
		}

		for i := range remaining {
			newCurrent := append(current, remaining[i])
			newRemaining := append(append([]string{}, remaining[:i]...), remaining[i+1:]...)
			generatePerms(newCurrent, newRemaining)
		}
	}

	generatePerms([]string{}, chars)

	perms := make([]string, 0, len(result))
	for perm := range result {
		perms = append(perms, perm)
	}
	return perms
}

// hitsAvoidPosition checks if a sequence of moves hits the avoid position
func hitsAvoidPosition(start Vector, moves string, avoid Vector) bool {
	current := start
	for _, move := range moves {
		dir := directions[string(move)]
		current = current.Add(dir)
		if current == avoid {
			return true
		}
	}
	return false
}

// minLength calculates the minimum length of moves needed
func minLength(str string, depth int) int {
	key := memoKey{str, depth}
	if val, exists := memoCache[key]; exists {
		return val
	}

	avoid := Vector{3, 0}
	if depth > 0 {
		avoid = Vector{0, 0}
	}

	var current Vector
	if depth == 0 {
		current = positions["A"]
	} else {
		current = positions["a"]
	}

	length := 0
	for _, char := range str {
		next := positions[string(char)]
		moveSet := generateMoveSet(current, next, avoid)

		if depth == 2 { // Part 1 limit
			length += len(moveSet[0])
		} else {
			minLen := int(^uint(0) >> 1) // Max int
			for _, move := range moveSet {
				l := minLength(move, depth+1)
				if l < minLen {
					minLen = l
				}
			}
			length += minLen
		}
		current = next
	}

	memoCache[key] = length
	return length
}

func main() {
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var codes []string

	// Read codes from file
	for scanner.Scan() {
		code := strings.TrimSpace(scanner.Text())
		if len(code) == 4 && code[3] == 'A' {
			codes = append(codes, code)
		}
	}

	totalComplexity := 0

	// Calculate complexity for part 1
	for _, code := range codes {
		length := minLength(code, 0)
		numeric, _ := strconv.Atoi(code[:3])
		complexity := length * numeric
		totalComplexity += complexity
		fmt.Printf("Code: %s, Length: %d, Numeric: %d, Complexity: %d\n",
			code, length, numeric, complexity)
	}

	fmt.Printf("\nPart 1 Total Complexity: %d\n", totalComplexity)
}
