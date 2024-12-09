#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int* fileID;
    int length;
} FileBlock;

// Function to parse the input and generate file blocks
FileBlock* parseInput(const char* input, int* numBlocks) {
    *numBlocks = strlen(input);
    FileBlock* blocks = malloc(*numBlocks * sizeof(FileBlock));
    int fileID = 0;

    for (int i = 0; i < *numBlocks; i++) {
        int length = input[i] - '0';
        blocks[i].length = length;

        if (i % 2 == 0) {
            // File block
            blocks[i].fileID = malloc(sizeof(int));
            *(blocks[i].fileID) = fileID++;
        } else {
            // Free space block
            blocks[i].fileID = NULL;
        }
    }

    return blocks;
}

// Function to move blocks to the leftmost free spaces
int* moveBlocks(FileBlock* blocks, int numBlocks, int* diskSize) {
    *diskSize = 0;
    // First, expand the blocks
    for (int i = 0; i < numBlocks; i++) {
        if (blocks[i].fileID != NULL) {
            // File block
            *diskSize += blocks[i].length;
        } else {
            // Free space block
            *diskSize += blocks[i].length;
        }
    }

    int* disk = malloc(*diskSize * sizeof(int));
    int diskIndex = 0;

    // Fill the disk array
    for (int i = 0; i < numBlocks; i++) {
        if (blocks[i].fileID != NULL) {
            // File block
            for (int j = 0; j < blocks[i].length; j++) {
                disk[diskIndex++] = *(blocks[i].fileID);
            }
        } else {
            // Free space block
            for (int j = 0; j < blocks[i].length; j++) {
                disk[diskIndex++] = -1;
            }
        }
    }

    // Move blocks to the leftmost free spaces
    for (int i = *diskSize - 1; i >= 0; i--) {
        if (disk[i] == -1) {
            continue;
        }

        for (int j = 0; j < i; j++) {
            if (disk[j] == -1) {
                disk[j] = disk[i];
                disk[i] = -1;
                break;
            }
        }
    }

    return disk;
}

// Function to calculate the filesystem checksum
long long calculateChecksum(int* disk, int diskSize) {
    long long total = 0;
    for (int i = 0; i < diskSize; i++) {
        if (disk[i] != -1) {
            total += (long long)i * disk[i]; // Ensure the calculation uses long long to avoid overflow
        }
    }
    return total;
}

int main() {
    // Read input from file
    FILE* file = fopen("inputdata.txt", "r");
    if (file == NULL) {
        printf("Error reading input file\n");
        return 1;
    }

    char input[10000];
    fgets(input, sizeof(input), file);
    fclose(file);

    // Trim any whitespace
    input[strcspn(input, "\n")] = 0;

    // Parse the input
    int numBlocks = 0;
    FileBlock* blocks = parseInput(input, &numBlocks);

    // Move blocks
    int diskSize = 0;
    int* disk = moveBlocks(blocks, numBlocks, &diskSize);

    // Calculate and print the checksum
    long long checksum = calculateChecksum(disk, diskSize);
    printf("Filesystem Checksum: %lld\n", checksum);

    // Free memory
    for (int i = 0; i < numBlocks; i++) {
        if (blocks[i].fileID != NULL) {
            free(blocks[i].fileID);
        }
    }
    free(blocks);
    free(disk);

    return 0;
}
