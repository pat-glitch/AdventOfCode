package main

import (
	"bufio"
	"fmt"
	"os"
)

const Size = 50

// Point structure to represent grid coordinates
type P struct {
	X, Y int
}

// Check if a point is within bounds
func (p P) Valid() bool {
	return p.X >= 0 && p.X < Size && p.Y >= 0 && p.Y < Size
}

// Parse input from "inputdata.txt" to build a frequency map
func parseInput() map[byte][]P {
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error: Unable to open input file.")
		os.Exit(1)
	}
	defer file.Close()

	freqs := map[byte][]P{}
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j, c := range line {
			if c == '.' {
				continue
			}

			ch := byte(c)

			freqs[ch] = append(freqs[ch], P{i, j})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: Unable to read input file.")
		os.Exit(1)
	}

	return freqs
}

func main() {
	// Parse the input file
	freqs := parseInput()

	// Map to store unique antinodes
	antiNodes := map[P]struct{}{}

	// Process each frequency group
	for _, locs := range freqs {
		// Iterate over all pairs of antennas
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				// Calculate the delta between two antennas
				delta := P{locs[b].X - locs[a].X, locs[b].Y - locs[a].Y}

				// Track the number of points out of bounds
				outOfBound := 0
				for period := 0; outOfBound < 2; period++ {
					outOfBound = 0

					// Calculate antinode 1 by subtracting multiple of delta
					anti1 := P{locs[a].X - period*delta.X, locs[a].Y - period*delta.Y}
					if anti1.Valid() {
						antiNodes[anti1] = struct{}{}
					} else {
						outOfBound++
					}

					// Calculate antinode 2 by adding multiple of delta
					anti2 := P{locs[b].X + period*delta.X, locs[b].Y + period*delta.Y}
					if anti2.Valid() {
						antiNodes[anti2] = struct{}{}
					} else {
						outOfBound++
					}
				}
			}
		}
	}

	// Print the total count of unique antinode locations
	fmt.Println(len(antiNodes))
}
