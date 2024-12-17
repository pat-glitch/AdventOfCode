#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define WALL '#'
#define START 'S'
#define END 'E'
#define TURN_COST 1000
#define MOVE_COST 1
#define START_DIR 1

typedef struct {
    int x, y;
} Point;

typedef struct {
    int dx, dy;
} Direction;

typedef struct {
    char **grid;
    Point start;
    Point end;
} Maze;

typedef struct {
    Point pos;
    int dir;
    int score;
    Point *path;
} QueueItem;

Direction directions[] = {
    {0, -1}, // up    (0)
    {1, 0},  // right (1)
    {0, 1},  // down  (2)
    {-1, 0}, // left  (3)
};

// Function to read the maze input
char** readInput(const char *filename, int *rows, int *cols) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        perror("Failed to open file");
        exit(1);
    }

    char **grid = malloc(100 * sizeof(char*));  // Maximum 100 rows (adjust if needed)
    char line[101];  // Buffer for each line

    *rows = 0;
    while (fgets(line, sizeof(line), file)) {
        grid[*rows] = malloc((strlen(line) + 1) * sizeof(char));
        strcpy(grid[*rows], line);
        (*rows)++;
    }

    *cols = strlen(grid[0]) - 1;  // Exclude the newline character
    fclose(file);
    return grid;
}

// Function to check if the position is valid
int isValidMove(Maze *m, Point p) {
    return p.y >= 0 && p.y < 100 && p.x >= 0 && p.x < strlen(m->grid[p.y]) && m->grid[p.y][p.x] != WALL;
}

// Function to initialize the maze
Maze parseMaze(char **input, int rows, int cols) {
    Maze maze;
    maze.grid = input;

    // Find start and end points
    for (int y = 0; y < rows; y++) {
        for (int x = 0; x < cols; x++) {
            if (input[y][x] == START) {
                maze.start = (Point){x, y};
            } else if (input[y][x] == END) {
                maze.end = (Point){x, y};
            }
        }
    }

    return maze;
}

// Function to find the lowest score to reach the end
int findLowestScore(Maze *m) {
    QueueItem queue[1000];
    int front = 0, rear = 0;

    // Initialize the queue with the starting point
    queue[rear++] = (QueueItem){m->start, START_DIR, 0, NULL};
    int visited[100][100][4] = {0};  // visited[y][x][direction]

    while (front < rear) {
        QueueItem current = queue[front++];
        
        if (m->end.x == current.pos.x && m->end.y == current.pos.y) {
            return current.score;
        }

        if (visited[current.pos.y][current.pos.x][current.dir]) {
            continue;
        }
        visited[current.pos.y][current.pos.x][current.dir] = 1;

        // Try moving forward in the current direction
        Point nextPos = {current.pos.x + directions[current.dir].dx, current.pos.y + directions[current.dir].dy};
        if (isValidMove(m, nextPos)) {
            queue[rear++] = (QueueItem){nextPos, current.dir, current.score + MOVE_COST, NULL};
        }

        // Try both possible 90-degree turns
        for (int newDir = (current.dir + 1) % 4; newDir != (current.dir + 3) % 4; newDir = (newDir + 1) % 4) {
            queue[rear++] = (QueueItem){current.pos, newDir, current.score + TURN_COST, NULL};
        }
    }

    return -1;  // No valid path found
}

// Function to count the number of unique tiles in the optimal paths
int countUniqueTiles(Point *paths, int pathLength) {
    int uniqueCount = 0;
    Point unique_points[10000];  // Maximum unique points in a path

    for (int i = 0; i < pathLength; i++) {
        int found = 0;
        for (int j = 0; j < uniqueCount; j++) {
            if (unique_points[j].x == paths[i].x && unique_points[j].y == paths[i].y) {
                found = 1;
                break;
            }
        }
        if (!found) {
            unique_points[uniqueCount++] = paths[i];
        }
    }

    return uniqueCount;
}

int main() {
    int rows, cols;
    char **input = readInput("inputdata.txt", &rows, &cols);
    Maze maze = parseMaze(input, rows, cols);

    int part1Score = findLowestScore(&maze);
    printf("Answer for Part 1: %d\n", part1Score);

    // Assuming findAllOptimalPaths function exists and works, get optimal paths (you'll need to implement it)
    // Point *paths = findAllOptimalPaths(&maze, part1Score);
    // int uniqueTiles = countUniqueTiles(paths, pathLength);
    // printf("Answer for Part 2: %d\n", uniqueTiles);

    // Free memory
    for (int i = 0; i < rows; i++) {
        free(input[i]);
    }
    free(input);

    return 0;
}
