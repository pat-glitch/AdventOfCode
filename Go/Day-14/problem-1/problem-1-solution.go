package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Room dimensions
const (
	width  = 101
	height = 103
)

// Robot structure
type Robot struct {
	posX, posY int
	velX, velY int
}

// Parse input data and return a slice of Robots
func parseInput(filePath string) ([]Robot, error) {
	var robots []Robot
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		pos := strings.Split(parts[0][2:], ",")
		vel := strings.Split(parts[1][2:], ",")

		posX, _ := strconv.Atoi(pos[0])
		posY, _ := strconv.Atoi(pos[1])
		velX, _ := strconv.Atoi(vel[0])
		velY, _ := strconv.Atoi(vel[1])

		robots = append(robots, Robot{posX, posY, velX, velY})
	}

	return robots, scanner.Err()
}

// Compute robot's position after given seconds with wrapping
func computePosition(robot Robot, seconds int) (int, int) {
	x := (robot.posX + robot.velX*seconds) % width
	y := (robot.posY + robot.velY*seconds) % height
	if x < 0 {
		x += width
	}
	if y < 0 {
		y += height
	}
	return x, y
}

// Determine the quadrant of a position
func determineQuadrant(x, y int) int {
	if x == width/2 && y == height/2 {
		return 0 // Middle point, not counted
	}
	if x < width/2 && y < height/2 {
		return 1 // Top-left
	}
	if x >= width/2 && y < height/2 {
		return 2 // Top-right
	}
	if x < width/2 && y >= height/2 {
		return 3 // Bottom-left
	}
	return 4 // Bottom-right
}

func main() {
	// Read robots from input file
	robots, err := parseInput("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Quadrant counters
	quadrantCounts := make(map[int]int)
	for _, robot := range robots {
		x, y := computePosition(robot, 100)
		quadrant := determineQuadrant(x, y)
		if quadrant != 0 { // Exclude middle point
			quadrantCounts[quadrant]++
		}
	}

	// Calculate safety factor
	safetyFactor := 1
	for i := 1; i <= 4; i++ {
		safetyFactor *= quadrantCounts[i]
	}

	fmt.Println("Safety Factor after 100 seconds:", safetyFactor)
}
