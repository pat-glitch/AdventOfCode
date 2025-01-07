package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions for up, down, left, and right
var directions = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func add(loc1, loc2 [2]int) [2]int {
	return [2]int{loc1[0] + loc2[0], loc1[1] + loc2[1]}
}

func calculateSides(borderPairs map[[2][2]int]struct{}) int {
	sides := 0
	for len(borderPairs) > 0 {
		for pair := range borderPairs {
			loc, out := pair[0], pair[1]
			delete(borderPairs, pair)

			// Calculate right and left directions
			di, dj := out[0], out[1]
			right := [2]int{dj, -di}
			left := [2]int{-dj, di}

			// Traverse the right side
			rl := add(loc, right)
			for {
				if _, exists := borderPairs[[2][2]int{rl, out}]; !exists {
					break
				}
				delete(borderPairs, [2][2]int{rl, out})
				rl = add(rl, right)
			}

			// Traverse the left side
			ll := add(loc, left)
			for {
				if _, exists := borderPairs[[2][2]int{ll, out}]; !exists {
					break
				}
				delete(borderPairs, [2][2]int{ll, out})
				ll = add(ll, left)
			}

			sides++
			break
		}
	}
	return sides
}

func calculatePart2Cost(data map[[2]int]rune) int {
	visited := make(map[[2]int]bool)
	locToRegion := make(map[[2]int]map[[2]int]bool)
	totalCost := 0

	for loc, char := range data {
		if visited[loc] {
			continue
		}

		// New region
		newRegion := make(map[[2]int]bool)
		locToRegion[loc] = newRegion
		stack := [][2]int{loc}
		border := make(map[[2][2]int]struct{})

		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[cur] {
				continue
			}
			visited[cur] = true
			newRegion[cur] = true

			for _, dir := range directions {
				nei := add(cur, dir)
				if data[nei] == char {
					if !visited[nei] {
						stack = append(stack, nei)
					}
				} else {
					border[[2][2]int{cur, dir}] = struct{}{}
				}
			}
		}

		area := len(newRegion)
		sides := calculateSides(border)
		totalCost += area * sides
	}

	return totalCost
}

func main() {
	// Read input from a file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Parse the grid into a map
	data := make(map[[2]int]rune)
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			data[[2]int{row, col}] = char
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Calculate Part 2 cost
	part2Cost := calculatePart2Cost(data)
	fmt.Println("Part 2 Total Cost:", part2Cost)
}
