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

// Function to check if an update is in the correct order
int is_update_valid(int line[], int count, Rule rules[], int rule_count) {
    for (int j = 0; j < rule_count; j++) {
        int a_pos = -1, b_pos = 1000;

        // Find positions of both pages in the update
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

// Function to perform topological sorting for a single update
void topological_sort(int line[], int count, Rule rules[], int rule_count) {
    // Repeatedly apply the rules to sort the line
    for (int i = 0; i < rule_count; i++) {
        for (int j = 0; j < count - 1; j++) {
            for (int k = j + 1; k < count; k++) {
                // If the rule a|b is violated (a comes after b), swap
                for (int r = 0; r < rule_count; r++) {
                    if (line[j] == rules[r].b && line[k] == rules[r].a) {
                        int temp = line[j];
                        line[j] = line[k];
                        line[k] = temp;
                    }
                }
            }
        }
    }
}

// Function to find the middle element of a valid update
int get_middle_element(int line[], int count) {
    return line[count / 2]; // Use integer division for consistent results
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

        // Check if the update is valid, and sort if it's not
        if (!is_update_valid(line, count, rules, NUM_RULES)) {
            // Sort the update if it's invalid
            topological_sort(line, count, rules, NUM_RULES);
            // Find the middle element after sorting
            int midpoint = get_middle_element(line, count);
            sum += midpoint;
        }
    }

    printf("Sum of middle elements of incorrectly ordered updates after sorting: %d\n", sum);

    fclose(f);
    return 0;
}
