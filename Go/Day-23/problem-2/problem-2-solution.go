package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Function to find all maximal cliques using Bron-Kerbosch algorithm
func bronKerbosch(r, p, x []string, adjList map[string]map[string]bool, cliques *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		// Found a maximal clique
		*cliques = append(*cliques, append([]string{}, r...))
		return
	}
	for i := 0; i < len(p); i++ {
		node := p[i]
		newR := append(r, node)
		newP := []string{}
		newX := []string{}
		for _, neighbor := range p {
			if adjList[node][neighbor] {
				newP = append(newP, neighbor)
			}
		}
		for _, neighbor := range x {
			if adjList[node][neighbor] {
				newX = append(newX, neighbor)
			}
		}
		bronKerbosch(newR, newP, newX, adjList, cliques)
		p = append(p[:i], p[i+1:]...) // Remove the node from P
		x = append(x, node)           // Add the node to X
		i--
	}
}

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

	// Find all maximal cliques
	allNodes := []string{}
	for node := range adjList {
		allNodes = append(allNodes, node)
	}
	cliques := [][]string{}
	bronKerbosch([]string{}, allNodes, []string{}, adjList, &cliques)

	// Find the largest clique
	largestClique := []string{}
	for _, clique := range cliques {
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}

	// Sort the largest clique alphabetically and generate the password
	sort.Strings(largestClique)
	password := strings.Join(largestClique, ",")

	// Output the password
	fmt.Println("Password to get into the LAN party:", password)
}
