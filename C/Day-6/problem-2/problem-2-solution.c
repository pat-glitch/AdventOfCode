#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#define INPUT 
#ifdef INPUT
#define FILE_NAME "inputdata.txt"
#define SIZE 130
#endif

#define NORTH 0
#define EAST 1
#define SOUTH 2
#define WEST 3

int obstacleMap[SIZE][SIZE] = {0};

void printMap(char map[SIZE][SIZE], int dirMap[SIZE][SIZE]);

int solve(int guardRow, int guardCol, int direction, int depth, char map[SIZE][SIZE], int dirMap[SIZE][SIZE]);

int main() {
    FILE* f = fopen(FILE_NAME, "r");

    int guardRow;
    int guardCol;

    int direction = NORTH;

    char map[SIZE][SIZE];
    // bitmap with each element being equal to the sum of 1 << dir for every dir that has visited
    int dirMap[SIZE][SIZE] = {0};

    for (int r = 0; r < SIZE; r++) {
        for (int c = 0; c < SIZE; c++) {
            fscanf(f, "%c ", &map[r][c]);
            if (map[r][c] == '^') {
                guardRow = r;
                guardCol = c;
            }
        }
    }

    int count = solve(guardRow, guardCol, NORTH, 0, map, dirMap);

    printMap(map, dirMap);

    printf("Part 1 (Total Obstacle Count): %d\n", count);
    fclose(f);
}

// depth == 0 if this is the initial call
// depth == 1 if this is a subsequent call, searching for loops
// will not recurse above 1
int solve(int guardRow, int guardCol, int direction, int depth, char map[SIZE][SIZE], int dirMap[SIZE][SIZE]) {
    int count = 0;

    while (1) {
        // if we are searching for a loop and we've been to a position before facing that direction, we have reached a loop
        if (depth == 1 && map[guardRow][guardCol] == 'X' && ((1 << direction) & dirMap[guardRow][guardCol]) != 0) {
            return 1;
        }

        map[guardRow][guardCol] = 'X';
        dirMap[guardRow][guardCol] |= 1 << direction;

        int nextRow = guardRow;
        int nextCol = guardCol;

        // Move the guard according to the direction
        if (direction == NORTH) {
            nextRow--;
        }
        else if (direction == EAST) {
            nextCol++;
        }
        else if (direction == SOUTH) {
            nextRow++;
        }
        else if (direction == WEST) {
            nextCol--;
        }

        // Check if the next position is within bounds
        if (nextRow >= 0 && nextRow < SIZE && nextCol >= 0 && nextCol < SIZE) {
            // If it's an obstacle, turn right (clockwise)
            if (map[nextRow][nextCol] == '#') {
                direction++;
                direction %= 4;
            }
            // If we're checking for loop formation and there's no 'X' and no obstacle
            else if (depth == 0 && map[nextRow][nextCol] != 'X' && !obstacleMap[nextRow][nextCol]) {
                char newMap[SIZE][SIZE];
                int newDirMap[SIZE][SIZE];
                memcpy(newMap, map, SIZE * SIZE * sizeof(char));
                memcpy(newDirMap, dirMap, SIZE * SIZE * sizeof(int));
                newMap[nextRow][nextCol] = '#';
                int obstacle = solve(guardRow, guardCol, (direction + 1) % 4, 1, newMap, newDirMap);
                if (obstacle) {
                    obstacleMap[nextRow][nextCol] = 1;
                    count++;
                }
                guardRow = nextRow;
                guardCol = nextCol;
            }
            else {
                guardRow = nextRow;
                guardCol = nextCol;
            }
        }
        // If the guard goes out of bounds
        else {
            break;
        }
    }
    // Return count if it's the initial recursion level (depth == 0)
    if (depth == 1) {
        return 0;
    }
    return count;
}

void printMap(char map[SIZE][SIZE], int dirMap[SIZE][SIZE]) {
    for (int r = 0; r < SIZE; r++) {
        for (int c = 0; c < SIZE; c++) {
            if (map[r][c] == 'X') {
                int n, e, s, w;
                n = (dirMap[r][c] & (1 << NORTH)) != 0;
                e = (dirMap[r][c] & (1 << EAST)) != 0;
                s = (dirMap[r][c] & (1 << SOUTH)) != 0;
                w = (dirMap[r][c] & (1 << WEST)) != 0;

                if ((n || s) && (e || w)) {
                    printf("+");
                }
                else if ((n || s)) {
                    printf("|");
                }
                else if ((e || w)) {
                    printf("-");
                }
            }
            else {
                printf("%c", map[r][c]);
            }
        }
        printf("\n");
    }
    printf("\n");
}
