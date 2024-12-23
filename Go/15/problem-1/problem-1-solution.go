package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Solution struct {
	dirs map[rune][2]int
}

func NewSolution() *Solution {
	return &Solution{
		dirs: map[rune][2]int{
			'^': {-1, 0},
			'v': {1, 0},
			'<': {0, -1},
			'>': {0, 1},
		},
	}
}

func (s *Solution) getRobotPos(grid [][]rune) (int, int) {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func (s *Solution) moving(grid [][]rune, pos [2]int, moves []rune, part int) [][]rune {
	for _, move := range moves {
		ny := pos[0] + s.dirs[move][0]
		nx := pos[1] + s.dirs[move][1]

		if grid[ny][nx] == '.' {
			pos = [2]int{ny, nx}
		} else if grid[ny][nx] == '#' {
			continue
		} else {
			edges, adjs := s.getAdjsAndEdges(grid, pos, move, part)
			blocked := 0
			dy, dx := s.dirs[move][0], s.dirs[move][1]
			for _, box := range edges {
				ny, nx := box[0]+dy, box[1]+dx
				if grid[ny][nx] == '#' {
					blocked++
				}
			}
			if blocked == 0 {
				grid = s.updateGrid(grid, adjs, move)
				pos = [2]int{pos[0] + dy, pos[1] + dx}
			}
		}
	}
	return grid
}

func (s *Solution) getAdjsAndEdges(grid [][]rune, pos [2]int, move rune, part int) ([][2]int, map[[2]int]struct{}) {
	y, x := pos[0], pos[1]
	dy, dx := s.dirs[move][0], s.dirs[move][1]

	adjs := make(map[[2]int]struct{})
	if part == 1 || move == '<' || move == '>' {
		for {
			ny, nx := y+dy, x+dx
			if grid[ny][nx] == '.' || grid[ny][nx] == '#' {
				return [][2]int{{ny - dy, nx - dx}}, adjs
			}
			y, x = ny, nx
			adjs[[2]int{y, x}] = struct{}{}
		}
	} else {
		edges := make([][2]int, 0)
		queue := [][2]int{{y, x}}
		for len(queue) > 0 {
			y, x = queue[0][0], queue[0][1]
			queue = queue[1:]
			if _, exists := adjs[[2]int{y, x}]; exists {
				continue
			}
			adjs[[2]int{y, x}] = struct{}{}
			ny, nx := y+dy, x+dx
			if grid[ny][nx] == '.' || grid[ny][nx] == '#' {
				edges = append(edges, [2]int{y, x})
			} else if grid[ny][nx] == '[' {
				queue = append(queue, [2]int{ny, nx}, [2]int{ny, nx + 1})
			} else if grid[ny][nx] == ']' {
				queue = append(queue, [2]int{ny, nx}, [2]int{ny, nx - 1})
			}
		}
		return edges, adjs
	}
	return nil, nil
}

func (s *Solution) updateGrid(grid [][]rune, adjs map[[2]int]struct{}, move rune) [][]rune {
	sortedCoords := make([][2]int, 0, len(adjs))
	for coord := range adjs {
		sortedCoords = append(sortedCoords, coord)
	}

	switch move {
	case '^':
		sortedCoords = sortCoords(sortedCoords, func(a, b [2]int) bool { return a[0] < b[0] })
	case 'v':
		sortedCoords = sortCoords(sortedCoords, func(a, b [2]int) bool { return a[0] > b[0] })
	case '<':
		sortedCoords = sortCoords(sortedCoords, func(a, b [2]int) bool { return a[1] < b[1] })
	case '>':
		sortedCoords = sortCoords(sortedCoords, func(a, b [2]int) bool { return a[1] > b[1] })
	}
	dy, dx := s.dirs[move][0], s.dirs[move][1]
	for _, coord := range sortedCoords {
		y, x := coord[0], coord[1]
		ny, nx := y+dy, x+dx
		grid[ny][nx] = grid[y][x]
		grid[y][x] = '.'
	}
	return grid
}

func sortCoords(coords [][2]int, less func(a, b [2]int) bool) [][2]int {
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			if less(coords[j], coords[i]) {
				coords[i], coords[j] = coords[j], coords[i]
			}
		}
	}
	return coords
}

func (s *Solution) part1(data []string) int {
	gridData, movesData := parseInput(data)
	grid := parseGrid(gridData)
	moves := []rune(strings.Join(movesData, ""))

	posY, posX := s.getRobotPos(grid)
	grid[posY][posX] = '.'

	grid = s.moving(grid, [2]int{posY, posX}, moves, 1)
	return s.getCoordsSum(grid, 1)
}

func (s *Solution) part2(data []string) int {
	gridData, movesData := parseInput(data)
	grid := parseGrid(gridData)
	moves := []rune(strings.Join(movesData, ""))

	grid = s.resizeGrid(grid)

	posY, posX := s.getRobotPos(grid)
	grid[posY][posX] = '.'

	grid = s.moving(grid, [2]int{posY, posX}, moves, 2)
	return s.getCoordsSum(grid, 2)
}

func (s *Solution) resizeGrid(grid [][]rune) [][]rune {
	mappings := map[rune]string{
		'#': "##",
		'O': "[]",
		'.': "..",
		'@': "@.",
	}
	newGrid := make([][]rune, len(grid))
	for i, row := range grid {
		newRow := make([]rune, 0, len(row)*2)
		for _, cell := range row {
			newRow = append(newRow, []rune(mappings[cell])...)
		}
		newGrid[i] = newRow
	}
	return newGrid
}

func (s *Solution) getCoordsSum(grid [][]rune, part int) int {
	box := 'O'
	if part == 2 {
		box = '['
	}
	sum := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == box {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func parseInput(data []string) ([]string, []string) {
	parts := strings.Split(strings.Join(data, "\n"), "\n\n")
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}

func parseGrid(data []string) [][]rune {
	grid := make([][]rune, len(data))
	for i, row := range data {
		grid[i] = []rune(row)
	}
	return grid
}

func main() {
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	solution := NewSolution()
	fmt.Println("Part 1:", solution.part1(data))
	fmt.Println("Part 2:", solution.part2(data))
}
