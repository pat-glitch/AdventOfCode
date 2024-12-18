#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

// Function to simulate the program's logic
char* simulate_program(int* registers, int* program, int program_length) {
    // Unpack registers
    int A = registers[0];
    int B = registers[1];
    int C = registers[2];

    // Allocate memory for outputs
    char* outputs = malloc(1000 * sizeof(char));
    outputs[0] = '\0';
    int output_index = 0;
    int instruction_pointer = 0;

    // Helper function to compute combo operand value
    int combo_value(int op) {
        switch (op) {
            case 0: case 1: case 2: case 3:
                return op; // Literal values 0-3
            case 4:
                return A;
            case 5:
                return B;
            case 6:
                return C;
            default:
                return 0; // Operand 7 will not appear in valid programs
        }
    }

    // Execute the program
    while (instruction_pointer < program_length) {
        int opcode = program[instruction_pointer];
        int operand = (instruction_pointer + 1 < program_length) ? program[instruction_pointer + 1] : 0;

        int operand = (instruction_pointer + 1 < program_length) ? program[instruction_pointer + 1] : 0;
        switch (opcode) {
            case 0: // adv: A = A // (2 ** combo_value(operand))
                A /= (int)pow(2, combo_value(operand));
                break;
            case 1: // bxl: B = B ^ operand
                B ^= operand;
                break;
            case 2: // bst: B = combo_value(operand) % 8
                B = combo_value(operand) % 8;
                break;
            case 3: // jnz: if A != 0, jump to operand
                if (A != 0) {
                    instruction_pointer = operand;
                    continue;
                }
                break;
            case 4: // bxc: B = B ^ C
                B ^= C;
                break;
            case 5: // out: output combo_value(operand) % 8
            {
                char temp[10];
                sprintf(temp, "%s%d", 
                        (output_index > 0) ? "," : "", 
                        combo_value(operand) % 8);
                strcat(outputs, temp);
                output_index++;
            }
                break;
            case 6: // bdv: B = A // (2 ** combo_value(operand))
                B = A / (int)pow(2, combo_value(operand));
                break;
            case 7: // cdv: C = A // (2 ** combo_value(operand))
                C = A / (int)pow(2, combo_value(operand));
                break;
            default:
                break;
        }

        // Increment instruction pointer by 2 (opcode + operand)
        instruction_pointer += 2;
    }

    return outputs;
}

int main() {
    // Input for the puzzle
    int registers[] = {64854237, 0, 0}; // A, B, C
    int program[] = {2, 4, 1, 1, 7, 5, 1, 5, 4, 0, 5, 5, 0, 3, 3, 0};
    int program_length = sizeof(program) / sizeof(program[0]);

    // Run the program
    char* output = simulate_program(registers, program, program_length);
    printf("Output: %s\n", output);

    // Free dynamically allocated memory
    free(output);

    return 0;
}