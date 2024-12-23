#include <stdio.h>
#include <string.h>

#define MAX_ROWS 1000
#define MAX_COLS 1000

char grid[MAX_ROWS][MAX_COLS];
int rows = 0, cols = 0;

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

int find_string_in_directions(int r, int c, const char *target) {
    int len = strlen(target);
    int count = 0;

    // Directions: horizontal, vertical, diagonal, and reverse
    int dr[] = {0, 1, 1, 1,  0, -1, -1, -1};
    int dc[] = {1, 1, 0, -1, -1, -1,  0,  1};

    for (int d = 0; d < 8; d++) {
        int k, nr = r, nc = c;
        for (k = 0; k < len; k++) {
            if (nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != target[k])
                break;
            nr += dr[d];
            nc += dc[d];
        }
        if (k == len)
            count++;
    }
    return count;
}

int count_occurrences(const char *target) {
    int total_count = 0;
    for (int r = 0; r < rows; r++) {
        for (int c = 0; c < cols; c++) {
            if (grid[r][c] == target[0]) {
                total_count += find_string_in_directions(r, c, target);
            }
        }
    }
    return total_count;
}

int main() {
    const char *filename = "./inputpuzzle.txt";
    const char *target = "XMAS";

    load_grid_from_file(filename);
    if (rows == 0 || cols == 0) {
        printf("Grid is empty or file not loaded correctly.\n");
        return 1;
    }

    int total_count = count_occurrences(target);
    printf("Total occurrences of '%s': %d\n", target, total_count);

    return 0;
}
