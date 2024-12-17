package main

import (
	"fmt"
	"math"
	"strings"
)

func simulateProgram(registers []int, program []int) string {
	// Unpack registers
	A := registers[0]
	B := registers[1]
	C := registers[2]
	var outputs []string
	instructionPointer := 0

	// Helper function to compute combo operand value
	comboValue := func(op int) int {
		switch op {
		case 0, 1, 2, 3:
			return op // Literal values 0-3
		case 4:
			return A
		case 5:
			return B
		case 6:
			return C
		default:
			return 0 // Operand 7 will not appear in valid programs
		}
	}

	// Execute the program
	for instructionPointer < len(program) {
		opcode := program[instructionPointer]
		operand := 0
		if instructionPointer+1 < len(program) {
			operand = program[instructionPointer+1]
		}

		switch opcode {
		case 0: // adv: A = A // (2 ** combo_value(operand))
			A /= int(math.Pow(2, float64(comboValue(operand))))
		case 1: // bxl: B = B ^ operand
			B ^= operand
		case 2: // bst: B = combo_value(operand) % 8
			B = comboValue(operand) % 8
		case 3: // jnz: if A != 0, jump to operand
			if A != 0 {
				instructionPointer = operand
				continue
			}
		case 4: // bxc: B = B ^ C
			B ^= C
		case 5: // out: output combo_value(operand) % 8
			outputs = append(outputs, fmt.Sprintf("%d", comboValue(operand)%8))
		case 6: // bdv: B = A // (2 ** combo_value(operand))
			B = A / int(math.Pow(2, float64(comboValue(operand))))
		case 7: // cdv: C = A // (2 ** combo_value(operand))
			C = A / int(math.Pow(2, float64(comboValue(operand))))
		}

		// Increment instruction pointer by 2 (opcode + operand)
		instructionPointer += 2
	}

	// Join outputs with commas
	return strings.Join(outputs, ",")
}

func main() {
	// Input for the puzzle
	registers := []int{64854237, 0, 0} // A, B, C
	program := []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 0, 5, 5, 0, 3, 3, 0}

	// Run the program
	output := simulateProgram(registers, program)
	fmt.Println("Output:", output)
}
