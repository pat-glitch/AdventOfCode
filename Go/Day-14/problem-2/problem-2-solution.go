package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type posT struct {
	x, y int
}

// Function to handle modulo with positive results
func mod(a, b posT) posT {
	return posT{(a.x%b.x + b.x) % b.x, (a.y%b.y + b.y) % b.y}
}

// The Robot struct holds position and velocity for each robot
type robotT struct {
	p, v posT
}

// Load the input from the given file
func loadInput(filename string) []robotT {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
	var robots []robotT
	for _, line := range lines {
		var r robotT
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.p.x, &r.p.y, &r.v.x, &r.v.y)
		if err != nil {
			continue
		}
		robots = append(robots, r)
	}
	return robots
}

// Part 1 solution function
func part1(robots []robotT, size posT) int {
	quads := [4]int{0, 0, 0, 0}

	for i := range robots {
		robots[i].p = mod(posT{robots[i].p.x + robots[i].v.x*100, robots[i].p.y + robots[i].v.y*100}, size)
		if robots[i].p.x != size.x/2 && robots[i].p.y != size.y/2 {
			ix := 0
			if robots[i].p.y < size.y/2 {
				ix = 1
			}
			iy := 0
			if robots[i].p.x < size.x/2 {
				iy = 1
			}
			quads[iy*2+ix]++
		}
	}

	return quads[0] * quads[1] * quads[2] * quads[3]
}

// Depth-first search to find clusters of robots
func dfs(curr posT, points map[posT]bool, visited map[posT]bool) int {
	if visited[curr] {
		return 0
	}
	visited[curr] = true
	clusterSize := 1

	directions := []posT{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, d := range directions {
		newPos := posT{curr.x + d.x, curr.y + d.y}
		if points[newPos] && !visited[newPos] {
			clusterSize += dfs(newPos, points, visited)
		}
	}
	return clusterSize
}

// Find a cluster of robots greater than a threshold
func findCluster(bots []robotT, threshold int) bool {
	points := make(map[posT]bool)
	for _, bot := range bots {
		points[bot.p] = true
	}

	visited := make(map[posT]bool)
	for bot := range points {
		clusterSize := dfs(bot, points, visited)
		if clusterSize >= threshold {
			return true
		}
	}
	return false
}

// Part 2 solution function
func part2(bots []robotT, size posT) int {
	seconds := 0

	for !findCluster(bots, 40) {
		for i := range bots {
			bots[i].p = mod(posT{bots[i].p.x + bots[i].v.x, bots[i].p.y + bots[i].v.y}, size)
		}
		seconds++
	}

	return seconds
}

func main() {
	// Load input
	actualValues := loadInput("inputdata.txt")

	fmt.Println("part1:", part1(actualValues, posT{101, 103}))

	// Part 2
	fmt.Println("part2:", part2(actualValues, posT{101, 103}))
}
