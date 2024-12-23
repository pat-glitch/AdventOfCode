#include <stdio.h>
#include <stdbool.h>
#include <string.h>

#define MAX_ROWS 200
#define MAX_COLS 200

int rows, cols;
int grid[MAX_ROWS][MAX_COLS];
bool visited[MAX_ROWS][MAX_COLS];

// Direction vectors for up, down, left, right
int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

// Function to check if a move is valid
bool is_valid_move(int x, int y, int prevHeight) {
    return x >= 0 && x < rows && y >= 0 && y < cols && !visited[x][y] && grid[x][y] == prevHeight + 1;
}

// DFS function to collect reachable 9s
void dfs_collect_9s(int x, int y, int height, bool reachable9s[MAX_ROWS][MAX_COLS]) {
    if (grid[x][y] == 9) {
        reachable9s[x][y] = true;
    }

    visited[x][y] = true;

    for (int i = 0; i < 4; i++) {
        int nx = x + dx[i];
        int ny = y + dy[i];
        if (is_valid_move(nx, ny, height)) {
            dfs_collect_9s(nx, ny, height + 1, reachable9s);
        }
    }

    visited[x][y] = false;  // Backtrack
}

// Function to calculate the total score with distinct 9s
int calculate_total_score_distinct() {
    int totalScore = 0;

    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            if (grid[i][j] == 0) {  // Identify trailhead
                // Reset visited matrix for each trailhead
                for (int r = 0; r < rows; r++) {
                    for (int c = 0; c < cols; c++) {
                        visited[r][c] = false;
                    }
                }

                // Track distinct reachable 9s
                bool reachable9s[MAX_ROWS][MAX_COLS] = {false};

                // Perform DFS from this trailhead
                dfs_collect_9s(i, j, 0, reachable9s);

                // Count distinct 9s
                int distinctCount = 0;
                for (int r = 0; r < rows; r++) {
                    for (int c = 0; c < cols; c++) {
                        if (reachable9s[r][c]) {
                            distinctCount++;
                        }
                    }
                }

                totalScore += distinctCount;  // Add distinct count for this trailhead
            }
        }
    }

    return totalScore;
}

int main() {
    // Open input file
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        printf("Error opening file\n");
        return 1;
    }

    // Read the input map
    rows = 0;
    char line[MAX_COLS + 1];  // Buffer for each line in the map

    while (fgets(line, sizeof(line), file) != NULL) {
        cols = strlen(line) - 1;  // Subtract 1 to remove the newline character
        for (int j = 0; j < cols; j++) {
            grid[rows][j] = line[j] - '0';  // Convert each character to integer
        }
        rows++;
    }

    fclose(file);

    // Calculate and print the total score
    int finalTotalScoreDistinct = calculate_total_score_distinct();
    printf("%d\n", finalTotalScoreDistinct);

    return 0;
}
