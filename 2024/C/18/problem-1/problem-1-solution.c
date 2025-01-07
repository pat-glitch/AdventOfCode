#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define GRID_SIZE 71
#define MAX_INPUT 3000
#define INF INT_MAX

// Direction vectors for moving up, down, left, or right
int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

// Function to check if a position is within bounds and not corrupted
int is_valid(int x, int y, int grid[GRID_SIZE][GRID_SIZE], int visited[GRID_SIZE][GRID_SIZE]) {
    return x >= 0 && x < GRID_SIZE && y >= 0 && y < GRID_SIZE && grid[x][y] == 0 && !visited[x][y];
}

// BFS to find the shortest path from (0, 0) to (70, 70)
int bfs(int grid[GRID_SIZE][GRID_SIZE]) {
    int visited[GRID_SIZE][GRID_SIZE] = {0};
    int queue[MAX_INPUT][3]; // queue to store x, y, and distance
    int front = 0, rear = 0;

    // Start BFS from (0, 0)
    queue[rear][0] = 0;
    queue[rear][1] = 0;
    queue[rear][2] = 0; // Distance
    rear++;
    visited[0][0] = 1;

    while (front < rear) {
        int x = queue[front][0];
        int y = queue[front][1];
        int dist = queue[front][2];
        front++;

        // Check if we reached the destination
        if (x == GRID_SIZE - 1 && y == GRID_SIZE - 1) {
            return dist;
        }

        // Explore all 4 directions
        for (int i = 0; i < 4; i++) {
            int nx = x + dx[i];
            int ny = y + dy[i];

            if (is_valid(nx, ny, grid, visited)) {
                queue[rear][0] = nx;
                queue[rear][1] = ny;
                queue[rear][2] = dist + 1;
                rear++;
                visited[nx][ny] = 1;
            }
        }
    }

    return -1; // If no path is found
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        perror("Error opening input file");
        return EXIT_FAILURE;
    }

    int grid[GRID_SIZE][GRID_SIZE] = {0};
    int x, y;

    // Read up to 1024 bytes from the input file and mark them as corrupted
    for (int i = 0; i < 1024 && fscanf(file, "%d,%d", &x, &y) == 2; i++) {
        if (x >= 0 && x < GRID_SIZE && y >= 0 && y < GRID_SIZE) {
            grid[x][y] = 1; // Mark the position as corrupted
        }
    }

    fclose(file);

    // Find the shortest path using BFS
    int result = bfs(grid);

    if (result != -1) {
        printf("The minimum number of steps needed to reach the exit is: %d\n", result);
    } else {
        printf("It is not possible to reach the exit.\n");
    }

    return 0;
}
