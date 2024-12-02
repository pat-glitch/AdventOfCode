#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h> // Required for strtok

#define MAX_LINE_LENGTH 1024

// Function to check if a row is "safe"
bool isSafeRow(int *row, int size) {
    // Check consecutive differences
    for (int i = 0; i < size - 1; i++) {
        int diff = abs(row[i + 1] - row[i]);
        if (diff < 1 || diff > 3) {
            return false; // Unsafe condition
        }
    }

    // Check monotonicity
    bool increasing = true, decreasing = true;
    for (int i = 0; i < size - 1; i++) {
        if (row[i] < row[i + 1]) {
            decreasing = false;
        } else if (row[i] > row[i + 1]) {
            increasing = false;
        }
    }

    if (!increasing && !decreasing) {
        return false; // Not monotonic, hence unsafe
    }

    return true; // Safe if all conditions pass
}

// Function to process the file and determine "safe" rows
void processFile(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        perror("Error opening file");
        return;
    }

    char line[MAX_LINE_LENGTH];
    int safeCount = 0, unsafeCount = 0;

    while (fgets(line, MAX_LINE_LENGTH, file)) {
        int row[MAX_LINE_LENGTH / 2]; // Temporary array for storing numbers
        int size = 0;

        // Parse numbers from the line
        char *token = strtok(line, " \t\n");
        while (token != NULL) {
            row[size++] = atoi(token);
            token = strtok(NULL, " \t\n");
        }

        // Check if the row is safe
        if (isSafeRow(row, size)) {
            safeCount++;
        } else {
            unsafeCount++;
        }
    }

    fclose(file);

    // Output the results
    printf("Total safe rows: %d\n", safeCount);
    printf("Total unsafe rows: %d\n", unsafeCount);
}

int main() {
    char filename[100];
    printf("Enter the name of the input file: ");
    scanf("%s", filename);

    processFile(filename);

    return 0;
}
