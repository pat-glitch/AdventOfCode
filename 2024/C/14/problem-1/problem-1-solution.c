#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <math.h>

#define ROOM_WIDTH 101
#define ROOM_HEIGHT 103
#define TIME_ELAPSED 100

typedef struct {
    int x, y;   // Position
    int vx, vy; // Velocity
} Robot;

void update_position(Robot *robot, int time) {
    robot->x = (robot->x + robot->vx * time) % ROOM_WIDTH;
    if (robot->x < 0) robot->x += ROOM_WIDTH;

    robot->y = (robot->y + robot->vy * time) % ROOM_HEIGHT;
    if (robot->y < 0) robot->y += ROOM_HEIGHT;
}

void forward(Robot *robots, int count, int time, Robot *updatedRobots) {
    for (int i = 0; i < count; i++) {
        updatedRobots[i] = robots[i];
        update_position(&updatedRobots[i], time);
    }
}

void calculate_quadrants(Robot *robots, int count, int *quadrants) {
    memset(quadrants, 0, 4 * sizeof(int));
    int u = ROOM_WIDTH / 2;
    int v = ROOM_HEIGHT / 2;

    for (int i = 0; i < count; i++) {
        int x = robots[i].x;
        int y = robots[i].y;

        if (x < u && y < v) {
            quadrants[0]++;
        } else if (x > u && y < v) {
            quadrants[1]++;
        } else if (x < u && y > v) {
            quadrants[2]++;
        } else if (x > u && y > v) {
            quadrants[3]++;
        }
    }
}

int part1(Robot *robots, int count) {
    Robot updatedRobots[count];
    forward(robots, count, TIME_ELAPSED, updatedRobots);

    int quadrants[4];
    calculate_quadrants(updatedRobots, count, quadrants);

    int safetyFactor = 1;
    for (int i = 0; i < 4; i++) {
        safetyFactor *= quadrants[i];
    }
    return safetyFactor;
}

int part2(Robot *robots, int count) {
    Robot updatedRobots[count];
    bool diverged = false;

    for (int t = 0; ; t++) {
        forward(robots, count, t, updatedRobots);

        // Check uniqueness of positions
        diverged = true;
        for (int i = 0; i < count; i++) {
            for (int j = i + 1; j < count; j++) {
                if (updatedRobots[i].x == updatedRobots[j].x && updatedRobots[i].y == updatedRobots[j].y) {
                    diverged = false;
                    break;
                }
            }
            if (!diverged) break;
        }

        if (diverged) return t;
    }
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        perror("Error opening input file");
        return EXIT_FAILURE;
    }

    char line[256];
    Robot robots[1000];
    int count = 0;

    while (fgets(line, sizeof(line), file)) {
        int x, y, vx, vy;
        sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy);
        robots[count++] = (Robot){x, y, vx, vy};
    }

    fclose(file);

    printf("Part 1 Safety Factor: %d\n", part1(robots, count));
    printf("Part 2 Time Until All Robots Diverge: %d\n", part2(robots, count));

    return 0;
}
