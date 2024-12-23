#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <regex.h>

#define MAX_LINE_LENGTH 4096

// Function to calculate the product of mul(x,y) pairs in the string
int process_mul(const char *line, regex_t *regex) {
    regmatch_t matches[3];  // stores positions of the matched groups
    int total_sum = 0;
    const char *ptr = line;
    
    while (regexec(regex, ptr, 3, matches, 0) == 0) {
        // Extract the matched numbers from the string
        char num1[4] = {0}, num2[4] = {0};
        
        // Get the first number (x)
        int length1 = matches[1].rm_eo - matches[1].rm_so;
        strncpy(num1, ptr + matches[1].rm_so, length1);
        int x = atoi(num1);
        
        // Get the second number (y)
        int length2 = matches[2].rm_eo - matches[2].rm_so;
        strncpy(num2, ptr + matches[2].rm_so, length2);
        int y = atoi(num2);
        
        // Add the product to the total sum
        total_sum += x * y;

        // Move the pointer past the current match
        ptr += matches[0].rm_eo;
    }

    return total_sum;
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        perror("Error opening file");
        return 1;
    }

    // Define the regex pattern for mul(x,y)
    regex_t regex;
    if (regcomp(&regex, "mul\\((\\d{1,3}),(\\d{1,3})\\)", REG_EXTENDED)) {
        fprintf(stderr, "Could not compile regex\n");
        return 1;
    }

    char line[MAX_LINE_LENGTH];
    int total_sum = 0;

    // Process each line from the file
    while (fgets(line, sizeof(line), file)) {
        total_sum += process_mul(line, &regex);
    }

    fclose(file);

    printf("Total Sum of all valid mul(x,y): %d\n", total_sum);

    // Free the compiled regex
    regfree(&regex);
    return 0;
}
