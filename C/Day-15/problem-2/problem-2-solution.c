#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define GRID_SIZE 100
#define MAX_MOVES 1000

typedef struct {
    int dy;
    int dx;
} Direction;

typedef struct {
    Direction dirs[128];
} Solution;

Solution* new_solution() {
    Solution* s = (Solution*)malloc(sizeof(Solution));
    s->dirs['^'] = (Direction){-1, 0};
    s->dirs['v'] = (Direction){1, 0};
    s->dirs['<'] = (Direction){0, -1};
    s->dirs['>'] = (Direction){0, 1};
    return s;
}

void get_robot_pos(char grid[GRID_SIZE][GRID_SIZE], int rows, int cols, int* posY, int* posX) {
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            if (grid[i][j] == '@') {
                *posY = i;
                *posX = j;
                return;
            }
        }
    }
    *posY = -1;
    *posX = -1;
}

void parse_input(const char* filename, char grid[GRID_SIZE][GRID_SIZE], int* rows, int* cols, char* moves) {
    FILE* file = fopen(filename, "r");
    if (!file) {
        perror("Error opening file");
        exit(1);
    }

    char line[GRID_SIZE];
    int grid_rows = 0;
    int grid_cols = 0;
    int reading_grid = 1;

    while (fgets(line, sizeof(line), file)) {
        if (strcmp(line, "\n") == 0) {
            reading_grid = 0;
            continue;
        }

        if (reading_grid) {
            strcpy(grid[grid_rows], line);
            grid_cols = strlen(line) - 1;  // Exclude newline character
            grid_rows++;
        } else {
            strcat(moves, line);
        }
    }

    fclose(file);
    *rows = grid_rows;
    *cols = grid_cols;
}

void move_robot(Solution* s, char grid[GRID_SIZE][GRID_SIZE], int* posY, int* posX, char* moves, int rows, int cols) {
    for (int i = 0; moves[i] != '\0'; i++) {
        char move = moves[i];
        int ny = *posY + s->dirs[move].dy;
        int nx = *posX + s->dirs[move].dx;

        if (ny >= 0 && ny < rows && nx >= 0 && nx < cols && grid[ny][nx] == '.') {
            *posY = ny;
            *posX = nx;
        }
    }
}

int calculate_sum(char grid[GRID_SIZE][GRID_SIZE], int rows, int cols, char box_type) {
    int sum = 0;
    for (int y = 0; y < rows; y++) {
        for (int x = 0; x < cols; x++) {
            if (grid[y][x] == box_type) {
                sum += 100 * y + x;
            }
        }
    }
    return sum;
}

int part1(Solution* s, char grid[GRID_SIZE][GRID_SIZE], int rows, int cols, char* moves) {
    int posY, posX;
    get_robot_pos(grid, rows, cols, &posY, &posX);
    grid[posY][posX] = '.';  // Clear the robot's initial position
    move_robot(s, grid, &posY, &posX, moves, rows, cols);
    return calculate_sum(grid, rows, cols, 'O');
}

int part2(Solution* s, char grid[GRID_SIZE][GRID_SIZE], int rows, int cols, char* moves) {
    // Implement resizing and other part2 logic as required
    // This is a placeholder to show extensibility
    return 0;
}

int main() {
    char grid[GRID_SIZE][GRID_SIZE] = {0};
    char moves[MAX_MOVES] = {0};
    int rows, cols;

    parse_input("inputdata.txt", grid, &rows, &cols, moves);

    Solution* solution = new_solution();

    printf("Part 1: %d\n", part1(solution, grid, rows, cols, moves));
    printf("Part 2: %d\n", part2(solution, grid, rows, cols, moves));

    free(solution);
    return 0;
}
