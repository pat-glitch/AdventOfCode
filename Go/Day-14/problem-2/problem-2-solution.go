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

func (r *Robot) UpdatePosition(time int) {
	r.x = discreteWrap(r.x+r.vx*time, RoomWidth)
	r.y = discreteWrap(r.y+r.vy*time, RoomHeight)
}

func discreteWrap(pos, size int) int {
	for pos < 0 {
		pos += size
	}
	return pos % size
}

func determineQuadrant(x, y int) int {
	if x == RoomWidth/2 || y == RoomHeight/2 {
		return -1 // Not in any quadrant
	}
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

func forward(robots []Robot, time int) []Robot {
	updatedRobots := make([]Robot, len(robots))
	for i, robot := range robots {
		robot.UpdatePosition(time)
		updatedRobots[i] = robot
	}
	return updatedRobots
}

func calculateQuadrants(robots []Robot) []int {
	u := RoomWidth / 2
	v := RoomHeight / 2
	quadrants := make([]int, 4)
	for _, robot := range robots {
		x, y := robot.x, robot.y
		if x < u && y < v {
			quadrants[0]++
		} else if x > u && y < v {
			quadrants[1]++
		} else if x < u && y > v {
			quadrants[2]++
		} else if x > u && y > v {
			quadrants[3]++
		}
	}
	return quadrants
}

func part1(robots []Robot) int {
	quadrants := calculateQuadrants(forward(robots, TimeElapsed))
	safetyFactor := 1
	for _, count := range quadrants {
		safetyFactor *= count
	}
	return safetyFactor
}

func part2(robots []Robot) int {
	for t := 0; ; t++ {
		updatedPositions := forward(robots, t)
		positionSet := make(map[string]struct{})
		for _, robot := range updatedPositions {
			key := fmt.Sprintf("%d,%d", robot.x, robot.y)
			positionSet[key] = struct{}{}
		}
		if len(positionSet) == len(robots) {
			return t
		}
	}
}

func main() {
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
		pParts := strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		x, _ := strconv.Atoi(pParts[0])
		y, _ := strconv.Atoi(pParts[1])
		vParts := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		vx, _ := strconv.Atoi(vParts[0])
		vy, _ := strconv.Atoi(vParts[1])
		robots = append(robots, Robot{x: x, y: y, vx: vx, vy: vy})
	}

	fmt.Println("Part 1 Safety Factor:", part1(robots))
	fmt.Println("Part 2 Time Until All Robots Diverge:", part2(robots))
}
