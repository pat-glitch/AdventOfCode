#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define INPUT

#ifdef TEST
#define FILE_NAME "test.txt"
#define NUM_RULES 21
#define NUM_LINES 6
#endif

#ifdef INPUT
#define FILE_NAME "inputdata.txt"
#define NUM_RULES 1176
#define NUM_LINES 202
#endif

typedef struct {
    int a;
    int b;
} Rule;

// Function to check if the update is valid according to the rules
int is_update_valid(int line[], int count, Rule rules[], int rule_count) {
    for (int j = 0; j < rule_count; j++) {
        int a_pos = -1, b_pos = 1000;
        
        for (int k = 0; k < count; k++) {
            if (line[k] == rules[j].a) a_pos = k;
            if (line[k] == rules[j].b) b_pos = k;
        }

        // If both a and b are present and a comes after b, it's invalid
        if (a_pos != -1 && b_pos != -1 && a_pos > b_pos) {
            return 0; // Invalid update
        }
    }
    return 1; // Valid update
}

int main() {
    FILE* f = fopen(FILE_NAME, "r");
    if (!f) {
        perror("Error opening file");
        return 1;
    }

    Rule rules[NUM_RULES];
    for (int i = 0; i < NUM_RULES; i++) {
        int a, b;
        fscanf(f, "%d|%d ", &a, &b);
        rules[i].a = a;
        rules[i].b = b;
    }

    int sum = 0;
    for (int i = 0; i < NUM_LINES; i++) {
        int count = 0;
        int line[100] = {0};

        char delim;
        do {
            fscanf(f, "%d%c", &line[count], &delim);
            count++;
        } while (delim == ',' && !feof(f));

        // Find the middle element
        int midpoint = line[count / 2];

        // Check if the update is valid
        if (is_update_valid(line, count, rules, NUM_RULES)) {
            sum += midpoint; // Add the middle element if valid
        }
    }

    printf("Sum of middle elements of valid updates: %d\n", sum);

    fclose(f);
    return 0;
}
