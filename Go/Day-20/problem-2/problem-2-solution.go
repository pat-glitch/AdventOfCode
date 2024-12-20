package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	FileName = "inputdata.txt" // Input file
	Size     = 141             // Map size (rows/columns)
	MaxCheat = 20              // Maximum cheat duration in taxicab distance
	Cutoff   = 100             // Minimum picoseconds saved for a valid cheat
)

var (
	mapData [Size][Size]rune
	dist    [Size][Size]int
)

func main() {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sr, sc, er, ec := -1, -1, -1, -1

	// Read map and initialize distances
	for r := 0; r < Size; r++ {
		if !scanner.Scan() {
			fmt.Println("Error reading input file")
			return
		}
		line := scanner.Text()
		for c := 0; c < Size; c++ {
			mapData[r][c] = rune(line[c])
			dist[r][c] = -1 // Unvisited cells
			if mapData[r][c] == 'S' {
				sr, sc = r, c // Starting position
			}
			if mapData[r][c] == 'E' {
				er, ec = r, c // Ending position
			}
		}
	}

	// Calculate distances along the direct path
	r, c, d := sr, sc, 0
	for r != er || c != ec {
		dist[r][c] = d
		if r-1 >= 0 && mapData[r-1][c] != '#' && dist[r-1][c] == -1 {
			r--
		} else if c+1 < Size && mapData[r][c+1] != '#' && dist[r][c+1] == -1 {
			c++
		} else if r+1 < Size && mapData[r+1][c] != '#' && dist[r+1][c] == -1 {
			r++
		} else if c-1 >= 0 && mapData[r][c-1] != '#' && dist[r][c-1] == -1 {
			c--
		}
		d++
	}
	dist[er][ec] = d

	numCheats := 0

	// Check for cheats
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if mapData[r][c] != '#' {
				numCheats += getCheatsFromPoint(r, c)
			}
		}
	}

	// Output the total number of cheats
	fmt.Println(numCheats)
}

// Calculate cheats from a specific point
func getCheatsFromPoint(r, c int) int {
	result := 0

	// Iterate over all points within MaxCheat taxicab distance
	for i := -MaxCheat; i <= MaxCheat; i++ {
		for j := -MaxCheat; j <= MaxCheat; j++ {
			cheatLen := int(math.Abs(float64(i)) + math.Abs(float64(j))) // Taxicab distance
			if cheatLen <= MaxCheat && r+i < Size && r+i >= 0 && c+j < Size && c+j >= 0 && mapData[r+i][c+j] != '#' {
				d1 := dist[r][c]
				d2 := dist[r+i][c+j]
				saved := 0
				if d1+cheatLen < d2 {
					saved = d2 - (d1 + cheatLen)
				}
				if saved >= Cutoff {
					result++
				}
			}
		}
	}
	return result
}
