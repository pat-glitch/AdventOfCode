package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Gate represents a logic gate with inputs, operation, and output
type Gate struct {
	Input1    string
	Input2    string
	Operation string
	Output    string
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Initialize data structures
	values := make(map[string]int) // Wire values
	gates := []Gate{}              // List of gates

	// Regular expressions to parse input
	valueRegex := regexp.MustCompile(`^(\w+): (\d+)$`)
	gateRegex := regexp.MustCompile(`^(\w+) (AND|OR|XOR) (\w+) -> (\w+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if match := valueRegex.FindStringSubmatch(line); match != nil {
			// Parse initial wire values
			wire := match[1]
			value, _ := strconv.Atoi(match[2])
			values[wire] = value
		} else if match := gateRegex.FindStringSubmatch(line); match != nil {
			// Parse gates
			gates = append(gates, Gate{
				Input1:    match[1],
				Operation: match[2],
				Input2:    match[3],
				Output:    match[4],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Simulate the gates
	processed := make(map[string]bool) // Track processed gates
	for len(gates) > 0 {
		remainingGates := []Gate{}
		for _, gate := range gates {
			// Ensure input wires are initialized
			value1, ok1 := values[gate.Input1]
			if !ok1 {
				value1 = 0
			}
			value2, ok2 := values[gate.Input2]
			if !ok2 {
				value2 = 0
			}

			if ok1 && ok2 {
				var result int
				switch gate.Operation {
				case "AND":
					result = value1 & value2
				case "OR":
					result = value1 | value2
				case "XOR":
					result = value1 ^ value2
				default:
					panic(fmt.Sprintf("Unknown operation: %s", gate.Operation))
				}

				// Correctly store the result of the operation
				values[gate.Output] = result
				processed[gate.Output] = true
			} else {
				remainingGates = append(remainingGates, gate)
			}
		}
		gates = remainingGates
	}

	// Collect values for wires starting with "z"
	outputBits := []string{}
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if value, exists := values[wire]; exists {
			outputBits = append(outputBits, strconv.Itoa(value))
		} else {
			break
		}
	}

	// Reverse the collected bits to ensure correct binary order
	reverse(outputBits)

	// Convert binary to decimal
	binaryOutput := strings.Join(outputBits, "")
	if len(binaryOutput) == 0 {
		fmt.Println("Error: No binary output produced.")
		return
	}
	decimalOutput, err := strconv.ParseInt(binaryOutput, 2, 64)
	if err != nil {
		fmt.Printf("Error converting binary to decimal: %v\n", err)
		return
	}

	// Print the result
	fmt.Println("Output:", decimalOutput)
}

// reverse reverses the order of a slice of strings
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
