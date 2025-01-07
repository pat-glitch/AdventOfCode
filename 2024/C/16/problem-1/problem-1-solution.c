#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

// Directions represent cardinal directions (East, North, West, South)
typedef enum {
    East = 0,
    North,
    West,
    South
} Direction;

// Coordinate represents a position in the maze
typedef struct {
    int x, y;
} Coordinate;

// State represents the current game state
typedef struct {
    Coordinate pos;
    Direction dir;
    int score;
    char *path;
    int index; // for heap
} State;

// PriorityQueue to manage states
typedef struct {
    State **data;
    int size;
} PriorityQueue;

void initQueue(PriorityQueue *pq) {
    pq->data = (State **)malloc(sizeof(State *));
    pq->size = 0;
}

void swap(State **a, State **b) {
    State *temp = *a;
    *a = *b;
    *b = temp;
}

void push(PriorityQueue *pq, State *state) {
    pq->size++;
    pq->data = (State **)realloc(pq->data, pq->size * sizeof(State *));
    pq->data[pq->size - 1] = state;
    int index = pq->size - 1;

    while (index > 0 && pq->data[index]->score < pq->data[(index - 1) / 2]->score) {
        swap(&pq->data[index], &pq->data[(index - 1) / 2]);
        index = (index - 1) / 2;
    }
}

State *pop(PriorityQueue *pq) {
    if (pq->size == 0) return NULL;
    
    State *top = pq->data[0];
    pq->data[0] = pq->data[pq->size - 1];
    pq->size--;
    pq->data = (State **)realloc(pq->data, pq->size * sizeof(State *));
    
    int index = 0;
    while (index * 2 + 1 < pq->size) {
        int left = index * 2 + 1;
        int right = index * 2 + 2;
        int smallest = index;
        
        if (left < pq->size && pq->data[left]->score < pq->data[smallest]->score) {
            smallest = left;
        }
        if (right < pq->size && pq->data[right]->score < pq->data[smallest]->score) {
            smallest = right;
        }
        
        if (smallest == index) break;
        
        swap(&pq->data[index], &pq->data[smallest]);
        index = smallest;
    }
    
    return top;
}

// Predefined direction movement deltas
const Coordinate dirDeltas[4] = {
    {1, 0},  // East
    {0, -1}, // North
    {-1, 0}, // West
    {0, 1}   // South
};

bool isValidMove(char **maze, int rows, int cols, Coordinate pos) {
    if (pos.x < 0 || pos.x >= cols || pos.y < 0 || pos.y >= rows) {
        return false;
    }
    return maze[pos.y][pos.x] != '#';
}

int solveMaze(char **maze, int rows, int cols) {
    Coordinate start, end;
    // Find start and end positions
    for (int y = 0; y < rows; y++) {
        for (int x = 0; x < cols; x++) {
            if (maze[y][x] == 'S') {
                start = (Coordinate){x, y};
            }
            if (maze[y][x] == 'E') {
                end = (Coordinate){x, y};
            }
        }
    }

    // Track visited states
    char visited[rows][cols][4]; // Assuming up to 4 directions
    memset(visited, 0, sizeof(visited));

    PriorityQueue pq;
    initQueue(&pq);

    // Initial state: start at 'S' facing East
    State *initialState = (State *)malloc(sizeof(State));
    initialState->pos = start;
    initialState->dir = East;
    initialState->score = 0;
    initialState->path = (char *)malloc(1 * sizeof(char)); // Empty path
    initialState->path[0] = '\0';
    initialState->index = 0;
    push(&pq, initialState);

    while (pq.size > 0) {
        State *current = pop(&pq);

        // Path success condition
        if (current->pos.x == end.x && current->pos.y == end.y) {
            int score = current->score;
            free(current->path);
            free(current);
            return score;
        }

        // Generate a unique state key
        if (visited[current->pos.y][current->pos.x][current->dir]) {
            free(current->path);
            free(current);
            continue;
        }

        visited[current->pos.y][current->pos.x][current->dir] = 1;

        // Try moving forward
        Coordinate forward = {
            current->pos.x + dirDeltas[current->dir].x,
            current->pos.y + dirDeltas[current->dir].y
        };

        if (isValidMove(maze, rows, cols, forward)) {
            State *forwardState = (State *)malloc(sizeof(State));
            forwardState->pos = forward;
            forwardState->dir = current->dir;
            forwardState->score = current->score + 1;
            forwardState->path = (char *)malloc(strlen(current->path) + 2);
            strcpy(forwardState->path, current->path);
            strcat(forwardState->path, ">");
            forwardState->index = 0;
            push(&pq, forwardState);
        }

        // Try rotating clockwise
        State *clockwiseState = (State *)malloc(sizeof(State));
        clockwiseState->pos = current->pos;
        clockwiseState->dir = (current->dir + 1) % 4;
        clockwiseState->score = current->score + 1000;
        clockwiseState->path = (char *)malloc(strlen(current->path) + 2);
        strcpy(clockwiseState->path, current->path);
        strcat(clockwiseState->path, "^");
        clockwiseState->index = 0;
        push(&pq, clockwiseState);

        // Try rotating counterclockwise
        State *counterClockwiseState = (State *)malloc(sizeof(State));
        counterClockwiseState->pos = current->pos;
        counterClockwiseState->dir = (current->dir - 1 + 4) % 4;
        counterClockwiseState->score = current->score + 1000;
        counterClockwiseState->path = (char *)malloc(strlen(current->path) + 2);
        strcpy(counterClockwiseState->path, current->path);
        strcat(counterClockwiseState->path, "v");
        counterClockwiseState->index = 0;
        push(&pq, counterClockwiseState);

        free(current->path);
        free(current);
    }

    return -1; // No path found
}

char **readMaze(const char *filename, int *rows, int *cols) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        printf("Error opening file: %s\n", filename);
        exit(1);
    }

    char **maze = NULL;
    char buffer[1000];
    *rows = 0;

    while (fgets(buffer, sizeof(buffer), file)) {
        (*rows)++;
        maze = (char **)realloc(maze, *rows * sizeof(char *));
        maze[*rows - 1] = (char *)malloc(strlen(buffer) + 1);
        strcpy(maze[*rows - 1], buffer);
    }

    fclose(file);

    *cols = strlen(maze[0]) - 1; // Exclude newline character
    return maze;
}

int main() {
    int rows, cols;
    char **maze = readMaze("inputdata.txt", &rows, &cols);
    printf("Lowest score: %d\n", solveMaze(maze, rows, cols));

    // Free memory
    for (int i = 0; i < rows; i++) {
        free(maze[i]);
    }
    free(maze);

    return 0;
}
