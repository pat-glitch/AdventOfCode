#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Function to evaluate the expression left-to-right (ignoring operator precedence)
long long evaluateExpression(char* expression) {
    long long result = 0;
    long long current = 0;
    char lastOperator = '+';

    for (int i = 0; expression[i]; i++) {
        if (expression[i] >= '0' && expression[i] <= '9') {
            current = current * 10 + (expression[i] - '0');
        } else if (expression[i] == '+' || expression[i] == '*') {
            if (lastOperator == '+') {
                result += current;
            } else if (lastOperator == '*') {
                result *= current;
            }
            current = 0;
            lastOperator = expression[i];
        }
    }

    // Final operation
    if (lastOperator == '+') {
        result += current;
    } else if (lastOperator == '*') {
        result *= current;
    }

    return result;
}

// Function to check all possible operator combinations
int checkEquation(long long result, int* numbers, int count) {
    int combinations = 1 << (count - 1); // Total combinations of '+' and '*'
    char expression[256];

    for (int i = 0; i < combinations; i++) {
        int pos = 0;
        for (int j = 0; j < count; j++) {
            pos += sprintf(expression + pos, "%d", numbers[j]);
            if (j < count - 1) {
                expression[pos++] = (i & (1 << j)) ? '*' : '+';
            }
        }
        expression[pos] = '\0';

        if (evaluateExpression(expression) == result) {
            printf("Valid: %s = %lld\n", expression, result);
            return 1; // Valid equation found
        }
    }

    return 0; // No valid equation found
}

int main() {
    FILE* file = fopen("inputdata.txt", "r");
    if (!file) {
        printf("Error opening file.\n");
        return 1;
    }

    long long totalSum = 0;
    char line[1024];

    while (fgets(line, sizeof(line), file)) {
        long long result;
        char* rightSide;
        int numbers[100];
        int count = 0;

        // Parse the line
        char* colon = strchr(line, ':');
        if (!colon) continue;

        *colon = '\0';
        result = atoll(line);
        rightSide = colon + 2; // Skip ": "

        // Parse numbers
        char* token = strtok(rightSide, " ");
        while (token) {
            numbers[count++] = atoi(token);
            token = strtok(NULL, " ");
        }

        // Check the equation
        if (checkEquation(result, numbers, count)) {
            totalSum += result; // Add the result from the left-hand side
        }
    }

    fclose(file);

    printf("Total sum of proven true results: %lld\n", totalSum);

    return 0;
}
