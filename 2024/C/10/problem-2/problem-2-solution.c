#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>

#define MAX_ROWS 200
#define MAX_COLS 200
#define MAX_TRAILS 1000000  // Size for hash map, adjust as necessary

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

// Function to generate a path signature as a string
char* generate_path_signature(int path[][2], int path_len) {
    char* sig = (char*)malloc(path_len * 12 * sizeof(char));  // Enough space for path coordinates
    sig[0] = '\0';  // Initialize empty string
    
    for (int i = 0; i < path_len; i++) {
        char coord[20];
        sprintf(coord, "%d,%d|", path[i][0], path[i][1]);  // Create coordinate signature
        strcat(sig, coord);  // Append to the signature string
    }

    return sig;
}

// DFS function to collect unique trails
void dfs_collect_trails(int x, int y, int height, bool trails[], int path[][2], int path_len) {
    if (visited[x][y]) return;

    // Add current position to the path
    path[path_len][0] = x;
    path[path_len][1] = y;
    path_len++;

    // Check if this is a destination (9)
    if (grid[x][y] == 9) {
        char* path_sig = generate_path_signature(path, path_len);
        int hash = 0;
        for (int i = 0; path_sig[i] != '\0'; i++) {
            hash = (hash * 31 + path_sig[i]) % MAX_TRAILS;  // Generate a unique hash
        }
        trails[hash] = true;  // Store the trail path signature in the trails map
        free(path_sig);  // Free the generated path signature string
    }

    visited[x][y] = true;

    for (int i = 0; i < 4; i++) {
        int nx = x + dx[i];
        int ny = y + dy[i];
        if (is_valid_move(nx, ny, height)) {
            dfs_collect_trails(nx, ny, height + 1, trails, path, path_len);
        }
    }

    visited[x][y] = false;  // Backtrack
}

// Function to calculate the total score with unique trails
int calculate_trailhead_ratings() {
    int totalRatings = 0;

    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            if (grid[i][j] == 0) {  // Trailhead
                // Reset visited matrix
                for (int r = 0; r < rows; r++) {
                    for (int c = 0; c < cols; c++) {
                        visited[r][c] = false;
                    }
                }

                // Track unique trails using an array (hashing)
                bool trails[MAX_TRAILS] = {false};  // Array for trail uniqueness

                // Start DFS with initial empty path
                int path[MAX_ROWS * MAX_COLS][2];  // Path to store coordinates
                int path_len = 0;

                dfs_collect_trails(i, j, 0, trails, path, path_len);

                // Count unique trails
                int trailRating = 0;
                for (int k = 0; k < MAX_TRAILS; k++) {
                    if (trails[k]) trailRating++;
                }

                totalRatings += trailRating;
                printf("Trailhead at (%d, %d) has rating: %d\n", i, j, trailRating);
            }
        }
    }

    return totalRatings;
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

    // Calculate and print total trailhead ratings
    int totalTrailheadRatings = calculate_trailhead_ratings();
    printf("Total Trailhead Ratings: %d\n", totalTrailheadRatings);

    return 0;
}
