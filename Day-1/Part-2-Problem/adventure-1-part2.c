#include <stdio.h>
#include <stdlib.h>

// Function to calculate similarity score
long long calculateSimilarityScore(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        perror("Error opening file");
        return -1;
    }

    // Maximum size for possible unique numbers
    int frequency[100000] = {0};
    int leftNumber, rightNumber;
    long long totalScore = 0;

    // Read the file line by line
    while (fscanf(file, "%d %d", &leftNumber, &rightNumber) != EOF) {
        // Count occurrences of numbers in the right list
        frequency[rightNumber]++;
    }

    // Rewind file to read left numbers again
    rewind(file);

    // Calculate the similarity score
    while (fscanf(file, "%d %d", &leftNumber, &rightNumber) != EOF) {
        totalScore += (long long)leftNumber * frequency[leftNumber];
    }

    fclose(file);
    return totalScore;
}

int main() {
    char filename[100];
    printf("Enter the name of the input file: ");
    scanf("%s", filename);

    long long similarityScore = calculateSimilarityScore(filename);

    if (similarityScore != -1) {
        printf("Total similarity score: %lld\n", similarityScore);
    }

    return 0;
}
