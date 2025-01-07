package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Directions represent cardinal directions (East, North, West, South)
type Direction int

const (
	East Direction = iota
	North
	West
	South
)

// Coordinate represents a position in the maze
type Coordinate struct {
	x, y int
}

// State represents the current game state
type State struct {
	pos   Coordinate
	dir   Direction
	score int
	path  []rune
	index int // for heap
}

// PriorityQueue to manage states
type PriorityQueue []*State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Predefined direction movement deltas
var dirDeltas = map[Direction]Coordinate{
	East:  {x: 1, y: 0},
	North: {x: 0, y: -1},
	West:  {x: -1, y: 0},
	South: {x: 0, y: 1},
}

func solveMaze(maze [][]rune) int {
	rows, cols := len(maze), len(maze[0])

	// Find start and end positions
	var start, end Coordinate
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if maze[y][x] == 'S' {
				start = Coordinate{x, y}
			}
			if maze[y][x] == 'E' {
				end = Coordinate{x, y}
			}
		}
	}

	// Track visited states
	visited := make(map[string]int)

	// Priority queue for Dijkstra-like search
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Initial state: start at 'S' facing East
	initialState := &State{
		pos:   start,
		dir:   East,
		score: 0,
		path:  []rune{},
	}
	heap.Push(&pq, initialState)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*State)

		// Path success condition
		if current.pos == end {
			return current.score
		}

		// Generate a unique state key
		stateKey := fmt.Sprintf("%d,%d,%d", current.pos.x, current.pos.y, current.dir)

		// Check if we've found a better path to this state
		if bestScore, exists := visited[stateKey]; exists && current.score >= bestScore {
			continue
		}
		visited[stateKey] = current.score

		// Try moving forward
		forward := Coordinate{
			x: current.pos.x + dirDeltas[current.dir].x,
			y: current.pos.y + dirDeltas[current.dir].y,
		}

		if isValidMove(maze, forward) {
			newPath := make([]rune, len(current.path))
			copy(newPath, current.path)
			newPath = append(newPath, '>')

			forwardState := &State{
				pos:   forward,
				dir:   current.dir,
				score: current.score + 1,
				path:  newPath,
			}
			heap.Push(&pq, forwardState)
		}

		// Try rotating clockwise
		clockwiseState := &State{
			pos:   current.pos,
			dir:   (current.dir + 1) % 4,
			score: current.score + 1000,
			path:  append(current.path, '^'),
		}
		heap.Push(&pq, clockwiseState)

		// Try rotating counterclockwise
		counterClockwiseDir := (current.dir - 1 + 4) % 4
		counterClockwiseState := &State{
			pos:   current.pos,
			dir:   counterClockwiseDir,
			score: current.score + 1000,
			path:  append(current.path, 'v'),
		}
		heap.Push(&pq, counterClockwiseState)
	}

	return -1 // No path found
}

func isValidMove(maze [][]rune, pos Coordinate) bool {
	rows, cols := len(maze), len(maze[0])

	// Out of bounds check
	if pos.x < 0 || pos.x >= cols || pos.y < 0 || pos.y >= rows {
		return false
	}

	// Wall check
	return maze[pos.y][pos.x] != '#'
}

func readMaze(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var maze [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return maze
}

func main() {
	maze := readMaze("inputdata.txt")
	fmt.Println("Lowest score:", solveMaze(maze))
}
