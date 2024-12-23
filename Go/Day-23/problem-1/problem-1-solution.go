package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Build the adjacency list
	adjList := make(map[string]map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			return
		}
		a, b := parts[0], parts[1]
		if adjList[a] == nil {
			adjList[a] = make(map[string]bool)
		}
		if adjList[b] == nil {
			adjList[b] = make(map[string]bool)
		}
		adjList[a][b] = true
		adjList[b][a] = true
	}

	// Find triangles
	triangles := [][]string{}
	for a := range adjList {
		for b := range adjList[a] {
			if b <= a { // Avoid duplicate triangles
				continue
			}
			for c := range adjList[b] {
				if c <= b || !adjList[a][c] { // Check triangle condition
					continue
				}
				triangles = append(triangles, []string{a, b, c})
			}
		}
	}

	// Filter triangles with at least one node starting with 't'
	count := 0
	for _, triangle := range triangles {
		for _, node := range triangle {
			if strings.HasPrefix(node, "t") {
				count++
				break
			}
		}
	}

	// Output the result
	fmt.Println("Number of triangles containing a node starting with 't':", count)
}
