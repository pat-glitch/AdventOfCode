package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

// Holds all the info about our maze - the grid and start/end points
type Maze struct {
	grid  [][]string
	start Point
	end   Point
}

// Represents a state in our search - where we are, which way we're facing,
// current score and the path we took to get here
type QueueItem struct {
	pos   Point
	dir   int
	score int
	path  []Point
}

// The four possible directions we can face/move in
var (
	directions = []Direction{
		{0, -1}, // up    (0)
		{1, 0},  // right (1)
		{0, 1},  // down  (2)
		{-1, 0}, // left  (3)
	}
)

// Maze elements and costs for moves
const (
	Wall     = "#"
	Start    = "S"
	End      = "E"
	TurnCost = 1000 // rotating 90 degrees costs 1000
	MoveCost = 1    // moving forward costs 1
	StartDir = 1    // start facing right (index 1 in directions)
)

func main() {
	input := readInput("inputdata.txt")
	maze := parseMaze(input)

	part1Score := findLowestScore(maze)
	fmt.Println("Answer for Part 1:", part1Score)

	paths := findAllOptimalPaths(maze, part1Score)
	uniqueTiles := countUniqueTiles(paths)
	fmt.Println("Answer for Part 2:", uniqueTiles)
}

func (p Point) add(d Direction) Point {
	return Point{p.x + d.dx, p.y + d.dy}
}

func (p Point) key(dir int) string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, dir)
}

func (m Maze) isValid(p Point) bool {
	return p.y >= 0 && p.y < len(m.grid) &&
		p.x >= 0 && p.x < len(m.grid[0]) &&
		m.grid[p.y][p.x] != Wall
}

func (m Maze) isEnd(p Point) bool {
	return p == m.end
}

func parseMaze(input []string) Maze {
	grid := make([][]string, len(input))
	var start, end Point

	for y, line := range input {
		grid[y] = strings.Split(line, "")
		for x, ch := range grid[y] {
			switch ch {
			case Start:
				start = Point{x, y}
			case End:
				end = Point{x, y}
			}
		}
	}

	return Maze{grid, start, end}
}

// Finds the lowest possible score to reach the end
func findLowestScore(m Maze) int {
	queue := []QueueItem{{m.start, StartDir, 0, nil}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].score < queue[j].score
		})

		current := queue[0]
		queue = queue[1:]

		if m.isEnd(current.pos) {
			return current.score
		}

		key := current.pos.key(current.dir)
		if visited[key] {
			continue
		}
		visited[key] = true

		// Try moving forward in current direction
		nextPos := current.pos.add(directions[current.dir])
		if m.isValid(nextPos) {
			queue = append(queue, QueueItem{
				nextPos,
				current.dir,
				current.score + MoveCost,
				nil,
			})
		}

		// Try both possible 90 degree turns
		queue = append(queue,
			QueueItem{current.pos, (current.dir + 1) % 4, current.score + TurnCost, nil},
			QueueItem{current.pos, (current.dir + 3) % 4, current.score + TurnCost, nil},
		)
	}

	return -1
}

// Similar to findLowestScore but keeps track of all paths that achieve the target score
func findAllOptimalPaths(m Maze, targetScore int) [][]Point {
	queue := []QueueItem{{m.start, StartDir, 0, []Point{m.start}}}
	visited := make(map[string]int)
	var paths [][]Point

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > targetScore {
			continue
		}

		key := current.pos.key(current.dir)
		if score, exists := visited[key]; exists && score < current.score {
			continue
		}
		visited[key] = current.score

		if m.isEnd(current.pos) && current.score == targetScore {
			paths = append(paths, current.path)
			continue
		}

		// Try moving forward
		nextPos := current.pos.add(directions[current.dir])
		if m.isValid(nextPos) {
			newPath := make([]Point, len(current.path))
			copy(newPath, current.path)
			queue = append(queue, QueueItem{
				nextPos,
				current.dir,
				current.score + MoveCost,
				append(newPath, nextPos),
			})
		}

		// Handle turns
		for _, newDir := range []int{(current.dir + 1) % 4, (current.dir + 3) % 4} {
			queue = append(queue, QueueItem{
				current.pos,
				newDir,
				current.score + TurnCost,
				current.path,
			})
		}
	}

	return paths
}

// Counts how many unique positions appear in any of the optimal paths
func countUniqueTiles(paths [][]Point) int {
	unique := make(map[string]bool)
	for _, path := range paths {
		for _, p := range path {
			unique[p.key(0)] = true
		}
	}
	return len(unique)
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
