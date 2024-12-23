#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_ROWS 1000
#define MAX_COLS 1000

typedef struct {
    int x;
    int y;
} Position;

void readInputFile(const char *filename, char grid[MAX_ROWS][MAX_COLS], int *rows, int *cols);
int markPath(char grid[MAX_ROWS][MAX_COLS], int rows, int cols);

int main() {
    char grid[MAX_ROWS][MAX_COLS];
    int rows = 0, cols = 0;

    // Read the grid from the file
    readInputFile("inputdata.txt", grid, &rows, &cols);

    // Calculate the distinct points visited
    int distinctCount = markPath(grid, rows, cols);
    printf("Distinct points visited: %d\n", distinctCount);

    return 0;
}

void readInputFile(const char *filename, char grid[MAX_ROWS][MAX_COLS], int *rows, int *cols) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        perror("Error opening file");
        exit(EXIT_FAILURE);
    }

    char line[MAX_COLS + 1];
    *rows = 0;
    while (fgets(line, sizeof(line), file)) {
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0'; // Remove the newline character
            len--;
        }
        strcpy(grid[*rows], line);
        (*rows)++;
    }
    *cols = strlen(grid[0]);
    fclose(file);
}

int markPath(char grid[MAX_ROWS][MAX_COLS], int rows, int cols) {
    // Directions: Up, Right, Down, Left
    Position directions[] = {{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
    int direction = 0; // Start facing up
    int distinctCount = 0;

    // Find the starting position of '^'
    Position current = {-1, -1};
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            if (grid[i][j] == '^') {
                current.x = i;
                current.y = j;
                grid[i][j] = 'X'; // Mark the starting position
                distinctCount++;
                break;
            }
        }
        if (current.x != -1) break; // Found the starting position
    }

    if (current.x == -1) {
        printf("Error: Starting position '^' not found.\n");
        exit(EXIT_FAILURE);
    }

    // Simulate the movement
    while (1) {
        int nextX = current.x + directions[direction].x;
        int nextY = current.y + directions[direction].y;

        // Check if `^` exits the grid
        if (nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols) {
            break;
        }

        // Check for an obstacle
        if (grid[nextX][nextY] == '#') {
            direction = (direction + 1) % 4; // Rotate right
        } else {
            // Move to the next cell
            current.x = nextX;
            current.y = nextY;

            // Mark the cell if not already visited
            if (grid[nextX][nextY] != 'X') {
                grid[nextX][nextY] = 'X';
                distinctCount++;
            }
        }
    }

    return distinctCount;
}
