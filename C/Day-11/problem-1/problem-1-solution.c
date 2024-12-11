#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#define SIZE 37

// Function to count the number of space-separated integers in the input string
int count(char *s) {
    int count = 1;
    while (*s++) {
        if (*s == ' ') count++;
    }
    return count;
}

// Function to convert the input string to an array of uint64_t values
uint64_t *to_arr(char *s, int count) {
    uint64_t *arr = malloc(sizeof(uint64_t) * count);
    for (int i = 0; *s; i++) {
        arr[i] = strtoul(s, &s, 10);
    }
    return arr;
}

// Function to count the number of digits in a number
uint64_t count_digits(uint64_t n) {
    int i = 0;
    for (; n; n /= 10, i++);
    return i;
}

// Function to perform a single blink on the stone array
void blink(uint64_t **arr, int *count) {
    int new_count = 0;

    // Calculate the new size for the array
    for (int i = 0; i < *count; i++) {
        if ((*arr)[i] == 0) {
            new_count++;  // zero becomes 1
        } else if (count_digits((*arr)[i]) % 2 == 0) {
            new_count += 2;  // even number of digits is split
        } else {
            new_count++;  // odd number of digits is multiplied
        }
    }

    // Allocate the new array with the computed size
    uint64_t *new_arr = malloc(sizeof(uint64_t) * new_count);
    int j = 0;

    // Perform the transformations based on stone values
    for (int i = 0; i < *count; i++, j++) {
        if ((*arr)[i] == 0) {
            new_arr[j] = 1;  // zero becomes 1
        } else if (count_digits((*arr)[i]) % 2 == 0) {
            // Split the number into two parts if it has an even number of digits
            int digits = count_digits((*arr)[i]);
            int pow = 1;
            for (int k = 0; k < digits / 2; k++) pow *= 10;

            new_arr[j++] = (*arr)[i] / pow;
            new_arr[j] = (*arr)[i] % pow;
        } else {
            // Multiply the stone by 2024 if the number of digits is odd
            new_arr[j] = (*arr)[i] * 2024;
        }
    }

    // Free the old array and update the pointer to the new array
    free(*arr);
    *arr = new_arr;
    *count = new_count;
}

int main() {
    // Open the file for reading
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        printf("Error opening file.\n");
        return 1;
    }

    // Read the file content into a buffer
    char content[SIZE];
    fread(content, sizeof(char), SIZE, file);
    content[SIZE-1] = 0;
    fclose(file);

    // Count the number of stones and convert the input string to an array
    int c = count(content);
    uint64_t *arr = to_arr(content, c);

    // Perform 25 blinks on the stones
    for (int i = 0; i < 25; i++) {
        blink(&arr, &c);
    }

    // Print the number of stones after 25 blinks
    printf("stone count: %i\n", c);

    // Free the allocated memory for the array
    free(arr);

    return 0;
}
