#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Function prototypes
void readInputFile(const char* filename);

int main() {
    readInputFile("inputdata.txt");
    // Add your code here
    return 0;
}

void readInputFile(const char* filename) {
    FILE* file = fopen(filename, "r");
    if (!file) {
        fprintf(stderr, "Error: Unable to open input file.\n");
        exit(1);
    }
    // Reading logic here
    fclose(file);
}