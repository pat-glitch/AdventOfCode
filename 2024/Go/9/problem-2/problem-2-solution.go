package main

import (
	"bufio"
	"fmt"
	"os"
)

type Info struct {
	i, len int
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file content
	scanner := bufio.NewScanner(file)
	var diskMap []byte
	for scanner.Scan() {
		diskMap = append(diskMap, scanner.Bytes()...)
	}

	// Initialize variables
	var fileID int
	var decoded []int
	isFile := true

	// Initialize slices to hold file and space information
	files := []Info{}
	spaces := []Info{}

	// Process the input diskMap
	for _, fileOrSpace := range diskMap {
		fsLen := fileOrSpace - '0' // converting character to integer
		if isFile {
			isFile = false
			files = append(files, Info{len(decoded), int(fsLen)})
			for i := 0; i < int(fsLen); i++ { // Corrected range logic here
				decoded = append(decoded, fileID)
			}
			fileID++
		} else {
			isFile = true
			if fsLen > 0 {
				spaces = append(spaces, Info{len(decoded), int(fsLen)})
			}
			for i := 0; i < int(fsLen); i++ { // Corrected range logic here
				decoded = append(decoded, -1) // -1 represents empty space
			}
		}
	}

	// Process and move files into available spaces
	for fi := len(files) - 1; fi >= 0; fi-- {
		f := files[fi]
		x := decoded[f.i]
		for si := 0; si < len(spaces); si++ {
			s := spaces[si]
			// Check if space is large enough to hold the file and comes before the file's start position
			if s.len >= f.len && s.i < f.i {
				// Move the file into the space
				for i := 0; i < f.len; i++ { // Corrected range logic here
					decoded[s.i] = x
					decoded[f.i] = -1
					s.i++
					f.i++
				}
				s.len -= f.len
				spaces[si] = s // Update the remaining space
				break
			}
		}
	}

	// Calculate the checksum based on the final positions
	var checkSum int
	for i, e := range decoded {
		if e > 0 { // Only count non-negative values (files)
			checkSum += i * e
		}
	}

	// Output the result
	fmt.Println(checkSum)
}
