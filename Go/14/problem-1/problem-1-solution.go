package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	RoomWidth   = 101
	RoomHeight  = 103
	TimeElapsed = 100
)

type Robot struct {
	x, y   int // Position as grid coordinates
	vx, vy int // Velocity in tiles per second
}

func (r *Robot) UpdatePosition() {
	// Move discretely for 100 seconds
	r.x = discreteWrap(r.x+r.vx*TimeElapsed, RoomWidth)
	r.y = discreteWrap(r.y+r.vy*TimeElapsed, RoomHeight)
}

// Discrete wrapping function for grid-based movement
func discreteWrap(pos, size int) int {
	// Handle wrapping for both positive and negative velocities
	for pos < 0 {
		pos += size
	}
	return pos % size
}

func determineQuadrant(x, y int) int {
	// Exclude middle lines exactly
	if x == RoomWidth/2 || y == RoomHeight/2 {
		return -1 // Not in any quadrant
	}

	// Quadrant numbering:
	// 0 | 1
	// -----
	// 2 | 3
	if x < RoomWidth/2 && y < RoomHeight/2 {
		return 0
	} else if x > RoomWidth/2 && y < RoomHeight/2 {
		return 1
	} else if x < RoomWidth/2 && y > RoomHeight/2 {
		return 2
	} else {
		return 3
	}
}

func main() {
	// Read input from file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		// Parse position
		pParts := strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		x, _ := strconv.Atoi(pParts[0])
		y, _ := strconv.Atoi(pParts[1])

		// Parse velocity
		vParts := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		vx, _ := strconv.Atoi(vParts[0])
		vy, _ := strconv.Atoi(vParts[1])

		robots = append(robots, Robot{
			x:  x,
			y:  y,
			vx: vx,
			vy: vy,
		})
	}

	// Debug: Show initial robot positions and movements
	fmt.Println("Initial Robots:")
	for i, robot := range robots {
		fmt.Printf("Robot %d: Initial(x=%d,y=%d) Velocity(vx=%d,vy=%d)\n",
			i, robot.x, robot.y, robot.vx, robot.vy)
	}

	// Update robot positions
	for i := range robots {
		robots[i].UpdatePosition()
	}

	// Debug: Show final robot positions
	fmt.Println("\nFinal Robots:")
	for i, robot := range robots {
		fmt.Printf("Robot %d: Final(x=%d,y=%d)\n", i, robot.x, robot.y)
	}

	// Count robots in quadrants
	quadrantCounts := make([]int, 4)
	for _, robot := range robots {
		quadrant := determineQuadrant(robot.x, robot.y)
		if quadrant >= 0 {
			quadrantCounts[quadrant]++
		}
	}

	// Calculate safety factor
	safetyFactor := 1
	for _, count := range quadrantCounts {
		safetyFactor *= count
	}

	// Print results
	fmt.Println("\nQuadrant Counts:", quadrantCounts)
	fmt.Println("Safety Factor:", safetyFactor)
}
