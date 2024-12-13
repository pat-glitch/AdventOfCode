#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

#define OFFSET 10000000000000LL

// Struct to represent coordinates
typedef struct {
    int64_t x;
    int64_t y;
} Coord;

// Function to parse input file
int parseInputFile(const char *filename, int64_t ****data, size_t *groupCount) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        perror("Error opening file");
        return -1;
    }

    char line[256];
    int64_t ***tempData = NULL;
    size_t tempGroupCount = 0;
    int64_t **currentGroup = NULL;
    size_t currentGroupSize = 0;

    while (fgets(line, sizeof(line), file)) {
        // Remove newline character
        line[strcspn(line, "\n")] = '\0';

        // Skip empty lines
        if (strlen(line) == 0) {
            continue;
        }

        // Extract numbers from the line
        int64_t *numbers = NULL;
        size_t matchCount = 0;
        char *token = strtok(line, " 	");

        while (token) {
            int64_t num;
            if (sscanf(token, "%lld", &num) == 1) {
                numbers = realloc(numbers, (matchCount + 1) * sizeof(int64_t));
                numbers[matchCount++] = num;
            }
            token = strtok(NULL, " 	");
        }

        // Add numbers to current group
        currentGroup = realloc(currentGroup, (currentGroupSize + 1) * sizeof(int64_t *));
        currentGroup[currentGroupSize++] = numbers;

        // If the group has 3 lines, finalize it
        if (currentGroupSize == 3) {
            tempData = realloc(tempData, (tempGroupCount + 1) * sizeof(int64_t **));
            tempData[tempGroupCount++] = currentGroup;

            currentGroup = NULL;
            currentGroupSize = 0;
        }
    }

    fclose(file);

    // Assign results
    *data = tempData;
    *groupCount = tempGroupCount;
    return 0;
}

int main() {
    int64_t ***data;
    size_t groupCount;

    if (parseInputFile("inputdata.txt", &data, &groupCount) != 0) {
        fprintf(stderr, "Error parsing input file\n");
        return 1;
    }

    int64_t total = 0;

    for (size_t i = 0; i < groupCount; ++i) {
        Coord a = {data[i][0][0], data[i][0][1]};
        Coord b = {data[i][1][0], data[i][1][1]};
        Coord c = {data[i][2][0] + OFFSET, data[i][2][1] + OFFSET};

        int64_t a1 = a.x, b1 = b.x, c1 = -c.x;
        int64_t a2 = a.y, b2 = b.y, c2 = -c.y;

        int64_t x = b1 * c2 - c1 * b2;
        int64_t y = c1 * a2 - a1 * c2;
        int64_t z = a1 * b2 - b1 * a2;

        if (z == 0 || x % z != 0 || y % z != 0) {
            continue;
        }

        x /= z;
        y /= z;

        if (x >= 0 && y >= 0) {
            total += x * 3 + y;
        }
    }

    printf("%lld\n", total);

    // Free allocated memory
    for (size_t i = 0; i < groupCount; ++i) {
        for (size_t j = 0; j < 3; ++j) {
            free(data[i][j]);
        }
        free(data[i]);
    }
    free(data);

    return 0;
}
