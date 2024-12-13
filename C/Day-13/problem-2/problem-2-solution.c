#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

#define OFFSET 10000000000000LL

// Struct for the linear equations
typedef struct {
    int64_t adx, ady;
    int64_t bdx, bdy;
    int64_t prizex, prizey;
} Equation;

// Function to parse input file
int parseInputFile(const char *filename, Equation **equations, size_t *count) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        perror("Error opening file");
        return -1;
    }

    char line[256];
    Equation *tempEquations = NULL;
    size_t tempCount = 0;
    Equation currentEquation;
    int lineType = 0; // 0: A, 1: B, 2: Prize

    while (fgets(line, sizeof(line), file)) {
        // Remove newline character
        line[strcspn(line, "\n")] = '\0';

        // Skip empty lines
        if (strlen(line) == 0) {
            continue;
        }

        // Parse Button A
        if (lineType == 0 && strstr(line, "Button A") != NULL) {
            sscanf(line, "Button A: X+%lld, Y+%lld", &currentEquation.adx, &currentEquation.ady);
        }
        // Parse Button B
        else if (lineType == 1 && strstr(line, "Button B") != NULL) {
            sscanf(line, "Button B: X+%lld, Y+%lld", &currentEquation.bdx, &currentEquation.bdy);
        }
        // Parse Prize
        else if (lineType == 2 && strstr(line, "Prize") != NULL) {
            sscanf(line, "Prize: X=%lld, Y=%lld", &currentEquation.prizex, &currentEquation.prizey);
            currentEquation.prizex += OFFSET;
            currentEquation.prizey += OFFSET;

            // Store the equation
            tempEquations = realloc(tempEquations, (tempCount + 1) * sizeof(Equation));
            tempEquations[tempCount++] = currentEquation;
        }

        lineType = (lineType + 1) % 3;
    }

    fclose(file);

    *equations = tempEquations;
    *count = tempCount;
    return 0;
}

// Solve the linear equations using integer arithmetic
int solveEquation(const Equation *eq, int64_t *a_count, int64_t *b_count) {
    int64_t a1 = eq->adx, b1 = eq->bdx, c1 = -eq->prizex;
    int64_t a2 = eq->ady, b2 = eq->bdy, c2 = -eq->prizey;

    int64_t determinant = a1 * b2 - b1 * a2;

    if (determinant == 0) {
        return 0; // No unique solution
    }

    int64_t x_numerator = b1 * c2 - c1 * b2;
    int64_t y_numerator = c1 * a2 - a1 * c2;

    if (x_numerator % determinant != 0 || y_numerator % determinant != 0) {
        return 0; // No integer solution
    }

    *a_count = x_numerator / determinant;
    *b_count = y_numerator / determinant;

    return (*a_count >= 0 && *b_count >= 0);
}

int main() {
    Equation *equations;
    size_t equationCount;

    if (parseInputFile("inputdata.txt", &equations, &equationCount) != 0) {
        fprintf(stderr, "Error parsing input file\n");
        return 1;
    }

    int64_t total = 0;

    for (size_t i = 0; i < equationCount; ++i) {
        int64_t a_count, b_count;

        if (solveEquation(&equations[i], &a_count, &b_count)) {
            total += 3 * a_count + b_count;
        }
    }

    printf("%lld\n", total);

    free(equations);
    return 0;
}
