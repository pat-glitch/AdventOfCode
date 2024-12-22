package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FileName = "inputdata.txt"
	Size     = 141
	Cutoff   = 100
)

func main() {
	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapData := make([][]rune, Size)
	dist := make([][]int, Size)
	for i := 0; i < Size; i++ {
		mapData[i] = make([]rune, Size)
		dist[i] = make([]int, Size)
		for j := 0; j < Size; j++ {
			dist[i][j] = -1
		}
	}

	var sr, sc, er, ec int
	for r := 0; r < Size; r++ {
		scanner.Scan()
		line := scanner.Text()
		for c := 0; c < Size; c++ {
			mapData[r][c] = rune(line[c])
			if mapData[r][c] == 'S' {
				sr, sc = r, c
			} else if mapData[r][c] == 'E' {
				er, ec = r, c
			}
		}
	}

	r, c := sr, sc
	d := 0

	// Traverse the path
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
			if c+2 < Size && mapData[r][c] != '#' && mapData[r][c+1] == '#' && mapData[r][c+2] != '#' {
				var saved int
				if dist[r][c] < dist[r][c+2] {
					saved = dist[r][c+2] - (dist[r][c] + 2)
				} else {
					saved = dist[r][c] - (dist[r][c+2] + 2)
				}
				if saved >= Cutoff {
					numCheats++
				}
			}
			if r+2 < Size && mapData[r][c] != '#' && mapData[r+1][c] == '#' && mapData[r+2][c] != '#' {
				var saved int
				if dist[r][c] < dist[r+2][c] {
					saved = dist[r+2][c] - (dist[r][c] + 2)
				} else {
					saved = dist[r][c] - (dist[r+2][c] + 2)
				}
				if saved >= Cutoff {
					numCheats++
				}
			}
		}
	}

	fmt.Println(numCheats)
}
