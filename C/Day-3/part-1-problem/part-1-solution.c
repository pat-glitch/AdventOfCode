#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

// Function to validate and process a "mul(x,y)" substring
int process_mul_substring(const char *start, int *product) {
    int x, y;
    // Parse the substring for the "mul(x,y)" format
    if (sscanf(start, "mul(%d,%d)", &x, &y) == 2) {
        // Check that both x and y are 1-3 digit integers
        if (x >= 0 && x <= 999 && y >= 0 && y <= 999) {
            *product = x * y;
            return 1;
        }
    }
    return 0;
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        perror("Error opening file");
        return 1;
    }

    char line[4096];
    int total_sum = 0;

    while (fgets(line, sizeof(line), file)) {
        char *ptr = line;
        while (*ptr) {
            // Look for "mul(" in the current substring
            if (strncmp(ptr, "mul(", 4) == 0) {
                int product = 0;
                if (process_mul_substring(ptr, &product)) {
                    total_sum += product;
                }
            }
            ptr++; // Move to the next character
        }
    }

    fclose(file);

    printf("Total Sum of all valid mul(x,y): %d\n", total_sum);
    return 0;
}
