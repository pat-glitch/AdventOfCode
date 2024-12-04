#include <stdio.h>
#include <stdbool.h>

#define FNAME "inputdata.txt"     // Path to input file
#define FSIZE (5 << 12)       // 20480 bytes (input file size limit)

// Instruction identifiers (32-bit little-endian integers)
#define MUL  0x286c756d       // *(int *)"mul("
#define DO   0x29286f64       // *(int *)"do()"
#define DON  0x276e6f64       // *(int *)"don'"
#define DONT 0x00292874       // *(int *)"t()"
#define MASK ((1 << 24) - 1)  // Mask for matching "t()" (ignores MSB)

// Input buffer
static char input[FSIZE];

// Function to parse 1-3 digit positive integers
static int num(const char **const c, const char sep) {
    if (**c < '1' || **c > '9')  // Positive numbers must start with 1-9
        return 0;
    int x = *(*c)++ & 15;        // Parse the first digit
    for (int i = 0; i < 2 && **c >= '0' && **c <= '9'; ++i) // Parse up to 2 more digits
        x = x * 10 + (*(*c)++ & 15);
    if (**c == sep) {            // Check if the next character is the separator
        ++(*c);                  // Move past the separator
        return x;
    }
    return 0;                    // Return 0 if parsing fails
}

int main(void) {
    // Open the input file
    FILE *f = fopen(FNAME, "rb");
    if (!f) {
        fputs("File not found.\n", stderr);
        return 1;
    }
    fread(input, sizeof(input), 1, f);  // Read entire file into memory
    fclose(f);

    int sum1 = 0, sum2 = 0, a, b;
    bool enabled = true;  // Multiplications are enabled at the start

    // Parse the input
    for (const char *c = input; *c; ) {
        switch (*(int *)c) {  // Match first four characters as a 32-bit integer
        case MUL:  // "mul("
            c += 4;
            if ((a = num(&c, ',')) && (b = num(&c, ')'))) {  // Parse x and y
                const int mul = a * b;
                sum1 += mul;            // Always add to sum1
                sum2 += mul * enabled; // Add to sum2 only if enabled
            }
            break;
        case DO:  // "do()"
            c += 4;
            enabled = true;  // Enable future multiplications
            break;
        case DON:  // "don'"
            c += 4;
            if ((*(int *)c & MASK) == DONT) {  // Match "t()" using the mask
                c += 3;
                enabled = false;  // Disable future multiplications
            }
            break;
        default:  // Skip unrecognized characters
            ++c;
        }
    }

    // Print results
    printf("Sum1: %d\n", sum1);  // Total sum of all mul(x,y)
    printf("Sum2: %d\n", sum2);  // Total sum of enabled mul(x,y)

    return 0;
}
