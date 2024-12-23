package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type FileBlock struct {
	fileID *int
	length int
}

func parseInput(input string) []FileBlock {
	var blocks []FileBlock
	for i := 0; i < len(input); i++ {
		length, _ := strconv.Atoi(string(input[i]))

		var fileID *int
		if i%2 == 0 {
			// File block
			id := i / 2
			fileID = &id
		}

		blocks = append(blocks, FileBlock{
			fileID: fileID,
			length: length,
		})
	}
	return blocks
}

func moveBlocks(blocks []FileBlock) []int {
	// Create a full representation of the disk
	disk := make([]int, 0)

	// First, expand the blocks
	for _, block := range blocks {
		if block.fileID != nil {
			// File block
			for j := 0; j < block.length; j++ {
				disk = append(disk, *block.fileID)
			}
		} else {
			// Free space block
			for j := 0; j < block.length; j++ {
				disk = append(disk, -1)
			}
		}
	}

	// Move blocks
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			if disk[j] == -1 {
				disk[j], disk[i] = disk[i], -1
				break
			}
		}
	}

	return disk
}

func calculateChecksum(disk []int) int {
	total := 0
	for i, fileID := range disk {
		if fileID != -1 {
			total += i * fileID
		}
	}
	return total
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Trim any whitespace and convert to string
	input := strings.TrimSpace(string(content))

	// Parse the input
	blocks := parseInput(input)

	// Move blocks
	disk := moveBlocks(blocks)

	// Calculate and print the checksum
	checksum := calculateChecksum(disk)
	fmt.Println("Filesystem Checksum:", checksum)
}
