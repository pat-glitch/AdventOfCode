#include <stdio.h>
#include <string.h>

#define MAX_ROWS 1000
#define MAX_COLS 1000

char grid[MAX_ROWS][MAX_COLS];
int rows = 0, cols = 0;

// Function to load the grid from a file
void load_grid_from_file(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        printf("Error: Could not open file.\n");
        return;
    }

    while (fgets(grid[rows], MAX_COLS, file)) {
        grid[rows][strcspn(grid[rows], "\n")] = '\0'; // Remove newline
        cols = strlen(grid[rows]);
        rows++;
    }
    fclose(file);
}

// Function to check if an X-MAS pattern exists centered at (r, c)
int is_x_mas(int r, int c) {
    // Ensure center is 'A'
    if (grid[r][c] != 'A') return 0;

    // Offsets for diagonals
    int dr1[] = {-1, 1};
    int dc1[] = {-1, 1};

    int dr2[] = {-1, 1};
    int dc2[] = {1, -1};

    // Check two diagonals for "MAS" patterns
    for (int d1 = 0; d1 < 2; d1++) {
        for (int d2 = 0; d2 < 2; d2++) {
            int r1 = r + dr1[d1], c1 = c + dc1[d1];
            int r2 = r + dr1[1 - d1], c2 = c + dc1[1 - d1];

            int r3 = r + dr2[d2], c3 = c + dc2[d2];
            int r4 = r + dr2[1 - d2], c4 = c + dc2[1 - d2];

            // Check bounds
            if (r1 < 0 || r1 >= rows || c1 < 0 || c1 >= cols ||
                r2 < 0 || r2 >= rows || c2 < 0 || c2 >= cols ||
                r3 < 0 || r3 >= rows || c3 < 0 || c3 >= cols ||
                r4 < 0 || r4 >= rows || c4 < 0 || c4 >= cols)
                continue;

            // Check for "MAS" in two diagonals
            if (((grid[r1][c1] == 'M' && grid[r2][c2] == 'S') || 
                 (grid[r1][c1] == 'S' && grid[r2][c2] == 'M')) &&
                ((grid[r3][c3] == 'M' && grid[r4][c4] == 'S') ||
                 (grid[r3][c3] == 'S' && grid[r4][c4] == 'M')))
                return 1;
        }
    }

    return 0;
}

// Function to count all X-MAS patterns in the grid
int count_x_mas() {
    int total_count = 0;

    // Traverse each cell, treating it as the center of a potential X-MAS
    for (int r = 1; r < rows - 1; r++) {
        for (int c = 1; c < cols - 1; c++) {
            total_count += is_x_mas(r, c);
        }
    }

    return total_count;
}

int main() {
    const char *filename = "./inputpuzzle.txt";

    // Load grid from the file
    load_grid_from_file(filename);
    if (rows == 0 || cols == 0) {
        printf("Grid is empty or file not loaded correctly.\n");
        return 1;
    }

    // Count X-MAS patterns
    int total_count = count_x_mas();
    printf("Total occurrences of X-MAS: %d\n", total_count);

    return 0;
}