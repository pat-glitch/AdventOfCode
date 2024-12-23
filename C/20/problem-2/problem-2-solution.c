#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#define FILE_NAME "inputdata.txt" // Input file
#define SIZE 141                  // Map size (rows/columns)
#define MAX_CHEAT 20              // Maximum cheat duration in taxicab distance
#define CUTOFF 100                // Minimum picoseconds saved for a valid cheat

int getCheatsFromPoint(int r, int c);

char map[SIZE][SIZE];
int dist[SIZE][SIZE];

int main() {
    FILE* f = fopen(FILE_NAME, "r");
    if (!f) {
        perror("Failed to open input file");
        return 1;
    }

    // Start (S) and end (E) positions
    int sr, sc, er, ec;

    // Read map and initialize distances
    for (int r = 0; r < SIZE; r++) {
        for (int c = 0; c < SIZE; c++) {
            fscanf(f, "%c ", &map[r][c]);
            dist[r][c] = -1; // Unvisited cells
            if (map[r][c] == 'S') {
                sr = r;
                sc = c; // Starting position
            }
            if (map[r][c] == 'E') {
                er = r;
                ec = c; // Ending position
            }
        }
    }
    fclose(f);

    // Calculate distances along the direct path
    int r = sr, c = sc, d = 0;
    while (r != er || c != ec) {
        dist[r][c] = d;
        if (r - 1 >= 0 && map[r - 1][c] != '#' && dist[r - 1][c] == -1) {
            r--;
        } else if (c + 1 < SIZE && map[r][c + 1] != '#' && dist[r][c + 1] == -1) {
            c++;
        } else if (r + 1 < SIZE && map[r + 1][c] != '#' && dist[r + 1][c] == -1) {
            r++;
        } else if (c - 1 >= 0 && map[r][c - 1] != '#' && dist[r][c - 1] == -1) {
            c--;
        }
        d++;
    }
    dist[er][ec] = d;

    int numCheats = 0;

    // Check for cheats
    for (int r = 0; r < SIZE; r++) {
        for (int c = 0; c < SIZE; c++) {
            if (map[r][c] != '#') {
                numCheats += getCheatsFromPoint(r, c);
            }
        }
    }

    // Output the total number of cheats
    printf("%d\n", numCheats);

    return 0;
}

// Calculate cheats from a specific point
int getCheatsFromPoint(int r, int c) {
    int result = 0;

    // Iterate over all points within MAX_CHEAT taxicab distance
    for (int i = -MAX_CHEAT; i <= MAX_CHEAT; i++) {
        for (int j = -MAX_CHEAT; j <= MAX_CHEAT; j++) {
            int cheatLen = abs(i) + abs(j); // Taxicab distance
            if (cheatLen <= MAX_CHEAT && r + i < SIZE && r + i >= 0 && c + j < SIZE && c + j >= 0 && map[r + i][c + j] != '#') {
                int d1 = dist[r][c];
                int d2 = dist[r + i][c + j];
                int saved = 0;
                if (d1 + cheatLen < d2) {
                    saved = d2 - (d1 + cheatLen);
                }
                if (saved >= CUTOFF) {
                    result++;
                }
            }
        }
    }
    return result;
}
