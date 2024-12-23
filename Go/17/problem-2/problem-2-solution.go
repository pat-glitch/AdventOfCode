package main

import (
	"fmt"
)

// The input program (instructions)
var program = []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 0, 5, 5, 0, 3, 3, 0}

// Run function executes the program with the given starting value of 'a'.
func run(a int) []int {
	ip := 0
	abc := [3]int{a, 0, 0} // Registers a, b, c
	var p1 []int

	combo := func(n int) int {
		if n < 4 {
			return n
		} else {
			return abc[n-4]
		}
	}

	for ip < len(program) {
		switch program[ip] {
		case 0:
			abc[0] = abc[0] >> combo(program[ip+1])
		case 1:
			abc[1] ^= program[ip+1]
		case 2:
			abc[1] = combo(program[ip+1]) & 7
		case 3:
			if abc[0] != 0 {
				ip = program[ip+1]
				continue
			}
		case 4:
			abc[1] ^= abc[2]
		case 5:
			p1 = append(p1, combo(program[ip+1])&7)
		case 6:
			abc[1] = abc[0] >> combo(program[ip+1])
		case 7:
			abc[2] = abc[0] >> combo(program[ip+1])
		default:
			fmt.Println("Unknown opcode")
		}
		ip += 2
	}

	return p1
}

// Function to recombine the steps into a single value of 'a'
func recombine(steps []int) int {
	i := steps[0]
	d := 10
	for _, c := range steps[1:] {
		i += (c >> 7) << d
		d += 3
	}
	return i
}

func main() {
	// Step 1: Generate the first outputs for all possible 'a' values
	steps := make([]int, 1024)
	for a := 0; a < 1024; a++ {
		steps[a] = run(a)[0]
	}

	// Step 2: Build possible sequences
	var candidates [][]int
	for i, step := range steps {
		if step == program[0] {
			candidates = append(candidates, []int{i})
		}
	}

	// Step 3: Extend candidates based on the program sequence
	for _, k := range program[1:] {
		var newCandidates [][]int
		for _, l := range candidates {
			current := l[len(l)-1] >> 3
			for i := 0; i < 8; i++ {
				if steps[(i<<7)+current] == k {
					newStep := append([]int{}, l...)
					newStep = append(newStep, (i<<7)+current)
					newCandidates = append(newCandidates, newStep)
				}
			}
		}
		candidates = newCandidates
	}

	// Step 4: Find the smallest valid 'a'
	ans := int(^uint(0) >> 1) // Initialize to a large value
	for _, l := range candidates {
		a := recombine(l)
		if fmt.Sprintf("%v", run(a)) == fmt.Sprintf("%v", program) {
			if a < ans {
				ans = a
			}
		}
	}

	// Final output
	fmt.Println("Part 2:", ans)
}
