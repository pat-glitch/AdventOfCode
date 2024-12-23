#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

// The input program (instructions)
int program[] = {2, 4, 1, 1, 7, 5, 1, 5, 4, 0, 5, 5, 0, 3, 3, 0};
int program_length = sizeof(program) / sizeof(program[0]);

// Combo function equivalent
int combo(int* abc, int n) {
    if (n < 4) {
        return n;
    } else {
        return abc[n-4];
    }
}

// Run function executes the program with given starting value of 'a'
int* run(int a, int* output_length) {
    int ip = 0;
    int abc[3] = {a, 0, 0}; // Registers a, b, c
    int* p1 = malloc(1024 * sizeof(int)); // Dynamic array to store outputs
    *output_length = 0;

    while (ip < program_length) {
        int opcode = program[ip];
        int operand = (ip + 1 < program_length) ? program[ip+1] : 0;

        switch (opcode) {
            case 0:
                abc[0] = abc[0] >> combo(abc, operand);
                break;
            case 1:
                abc[1] ^= operand;
                break;
            case 2:
                abc[1] = combo(abc, operand) & 7;
                break;
            case 3:
                if (abc[0] != 0) {
                    ip = operand;
                    continue;
                }
                break;
            case 4:
                abc[1] ^= abc[2];
                break;
            case 5:
                p1[(*output_length)++] = combo(abc, operand) & 7;
                break;
            case 6:
                abc[1] = abc[0] >> combo(abc, operand);
                break;
            case 7:
                abc[2] = abc[0] >> combo(abc, operand);
                break;
        }
        ip += 2;
    }

    return p1;
}

// Recombine function to reconstruct 'a' from steps
int recombine(int* steps, int steps_length) {
    int i = steps[0];
    int d = 10;
    for (int j = 1; j < steps_length; j++) {
        i += (steps[j] >> 7) << d;
        d += 3;
    }
    return i;
}

// Compare arrays for validation
int compare_arrays(int* arr1, int len1, int* arr2, int len2) {
    if (len1 != len2) return 0;
    for (int i = 0; i < len1; i++) {
        if (arr1[i] != arr2[i]) return 0;
    }
    return 1;
}

int main() {
    // Generate first outputs for possible 'a' values
    int steps[1024];
    int output_lengths[1024];
    for (int a = 0; a < 1024; a++) {
        int output_length;
        int* result = run(a, &output_length);
        
        steps[a] = (output_length > 0) ? result[0] : -1;
        output_lengths[a] = output_length;
        
        free(result);
    }

    // Build possible sequences
    int** candidates = malloc(1024 * sizeof(int*));
    int* candidate_lengths = malloc(1024 * sizeof(int));
    int candidates_count = 0;

    for (int i = 0; i < 1024; i++) {
        if (steps[i] == program[0]) {
            candidates[candidates_count] = malloc(sizeof(int));
            candidates[candidates_count][0] = i;
            candidate_lengths[candidates_count] = 1;
            candidates_count++;
        }
    }

    // Extend candidates based on program sequence
    for (int k_index = 1; k_index < program_length; k_index++) {
        int k = program[k_index];
        
        int** new_candidates = malloc(1024 * sizeof(int*));
        int* new_candidate_lengths = malloc(1024 * sizeof(int));
        int new_candidates_count = 0;

        for (int l = 0; l < candidates_count; l++) {
            int current = candidates[l][candidate_lengths[l]-1] >> 3;
            for (int i = 0; i < 8; i++) {
                int new_step = (i << 7) + current;
                if (steps[new_step] == k) {
                    new_candidates[new_candidates_count] = malloc((candidate_lengths[l] + 1) * sizeof(int));
                    memcpy(new_candidates[new_candidates_count], candidates[l], candidate_lengths[l] * sizeof(int));
                    new_candidates[new_candidates_count][candidate_lengths[l]] = new_step;
                    new_candidate_lengths[new_candidates_count] = candidate_lengths[l] + 1;
                    new_candidates_count++;
                }
            }
        }

        // Free old candidates
        for (int i = 0; i < candidates_count; i++) {
            free(candidates[i]);
        }
        free(candidates);
        free(candidate_lengths);

        // Update candidates
        candidates = new_candidates;
        candidate_lengths = new_candidate_lengths;
        candidates_count = new_candidates_count;

        if (candidates_count == 0) break;
    }

    // Find smallest valid 'a'
    int ans = INT_MAX;
    for (int l = 0; l < candidates_count; l++) {
        int a = recombine(candidates[l], candidate_lengths[l]);
        int output_length;
        int* result = run(a, &output_length);

        if (compare_arrays(result, output_length, program, program_length)) {
            ans = (a < ans) ? a : ans;
        }

        free(result);
    }

    // Free remaining memory
    for (int i = 0; i < candidates_count; i++) {
        free(candidates[i]);
    }
    free(candidates);
    free(candidate_lengths);

    // Final output
    printf("Part 2: %d\n", ans);

    return 0;
}