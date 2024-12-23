#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SIZE 50

// Structure to represent grid coordinates
typedef struct {
    int x, y;
} Point;

// Function to check if a point is within bounds
int isValid(Point p) {
    return p.x >= 0 && p.x < SIZE && p.y >= 0 && p.y < SIZE;
}

// Node structure for linked list
typedef struct Node {
    Point point;
    struct Node* next;
} Node;

// Add a point to a linked list
Node* addPoint(Node* head, Point p) {
    Node* newNode = (Node*)malloc(sizeof(Node));
    newNode->point = p;
    newNode->next = head;
    return newNode;
}

// Check if a point exists in the unique points list
int containsPoint(Node* head, Point p) {
    while (head) {
        if (head->point.x == p.x && head->point.y == p.y) {
            return 1;
        }
        head = head->next;
    }
    return 0;
}

// Parse input from "inputdata.txt" to build a frequency map
Node** parseInput() {
    FILE* file = fopen("inputdata.txt", "r");
    if (!file) {
        fprintf(stderr, "Error: Unable to open input file.\n");
        exit(1);
    }

    Node** freqs = (Node**)calloc(256, sizeof(Node*)); // ASCII characters

    char line[SIZE + 2]; // Accommodate newline and null-terminator
    int i = 0;
    while (fgets(line, sizeof(line), file)) {
        for (int j = 0; line[j] != '\0' && line[j] != '\n'; j++) {
            if (line[j] == '.') continue;
            char ch = line[j];
            Point p = {i, j};
            freqs[ch] = addPoint(freqs[ch], p);
        }
        i++;
    }

    fclose(file);
    return freqs;
}

int main() {
    // Parse the input file
    Node** freqs = parseInput();

    // Linked list to store unique antinodes
    Node* antiNodes = NULL;

    // Process each frequency group
    for (int ch = 0; ch < 256; ch++) {
        if (!freqs[ch]) continue;

        Node* locs = freqs[ch];
        for (Node* a = locs; a != NULL; a = a->next) {
            for (Node* b = a->next; b != NULL; b = b->next) {
                Point delta = {b->point.x - a->point.x, b->point.y - a->point.y};
                int outOfBound = 0;
                for (int period = 0; outOfBound < 2; period++) {
                    outOfBound = 0;

                    // Calculate antinode 1
                    Point anti1 = {a->point.x - period * delta.x, a->point.y - period * delta.y};
                    if (isValid(anti1)) {
                        if (!containsPoint(antiNodes, anti1)) {
                            antiNodes = addPoint(antiNodes, anti1);
                        }
                    } else {
                        outOfBound++;
                    }

                    // Calculate antinode 2
                    Point anti2 = {b->point.x + period * delta.x, b->point.y + period * delta.y};
                    if (isValid(anti2)) {
                        if (!containsPoint(antiNodes, anti2)) {
                            antiNodes = addPoint(antiNodes, anti2);
                        }
                    } else {
                        outOfBound++;
                    }
                }
            }
        }
    }

    // Count unique antinodes
    int count = 0;
    Node* temp = antiNodes;
    while (temp) {
        count++;
        temp = temp->next;
    }

    printf("%d\n", count);

    // Free memory
    for (int ch = 0; ch < 256; ch++) {
        Node* temp = freqs[ch];
        while (temp) {
            Node* next = temp->next;
            free(temp);
            temp = next;
        }
    }
    free(freqs);

    while (antiNodes) {
        Node* next = antiNodes->next;
        free(antiNodes);
        antiNodes = next;
    }

    return 0;
}
