package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type State struct {
	pos      Point
	distance int
	cheated  bool
}

var directions = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
	grid, start, end := readInput("inputdata.txt")
	result := findShortestPath(grid, start, end)
	fmt.Println("Shortest path with cheat:", result)
}

func readInput(filename string) ([][]rune, Point, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	var start, end Point

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		if idx := strings.IndexRune(line, 'S'); idx != -1 {
			start = Point{idx, y}
		}
		if idx := strings.IndexRune(line, 'E'); idx != -1 {
			end = Point{idx, y}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid, start, end
}

func findShortestPath(grid [][]rune, start, end Point) int {
	rows, cols := len(grid), len(grid[0])
	queue := []State{{pos: start, distance: 0, cheated: false}}
	visited := make(map[State]bool)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.pos == end {
			return cur.distance
		}

		if visited[cur] {
			continue
		}
		visited[cur] = true

		for _, d := range directions {
			next := Point{cur.pos.x + d.x, cur.pos.y + d.y}
			if next.x < 0 || next.x >= cols || next.y < 0 || next.y >= rows {
				continue
			}

			cell := grid[next.y][next.x]
			if cell == '#' && cur.cheated {
				continue
			}

			newState := State{pos: next, distance: cur.distance + 1, cheated: cur.cheated}
			if cell == '#' && !cur.cheated {
				newState.cheated = true
			}

			queue = append(queue, newState)
		}
	}

	return -1 // No path found
}
