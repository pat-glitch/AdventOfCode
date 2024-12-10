#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Disk structure representing the current state
typedef struct {
    int *blocks; // -1 represents free space, non-negative numbers represent file IDs
    int size;    // Total size of the disk
} Disk;

// Function to parse the disk map from the input string
void parseDiskMap(const char *input, int **files, int *filesCount, int **spaces, int *spacesCount) {
    int len = strlen(input);
    *filesCount = (len + 1) / 2;
    *spacesCount = len / 2;

    *files = (int *)malloc((*filesCount) * sizeof(int));
    *spaces = (int *)malloc((*spacesCount) * sizeof(int));

    for (int i = 0; i < len; i++) {
        int size = input[i] - '0';
        if (i % 2 == 0) {
            (*files)[i / 2] = size;
        } else {
            (*spaces)[i / 2] = size;
        }
    }
}

// Function to create the initial disk state
Disk createInitialDisk(int *files, int filesCount, int *spaces, int spacesCount) {
    int maxSize = 0;
    for (int i = 0; i < filesCount; i++) {
        maxSize += files[i];
        if (i < spacesCount) {
            maxSize += spaces[i];
        }
    }

    Disk disk;
    disk.blocks = (int *)malloc(maxSize * sizeof(int));
    disk.size = maxSize;

    int index = 0, fileID = 0;
    for (int i = 0; i < filesCount; i++) {
        // add file blocks
        for (int j = 0; j < files[i]; j++) {
            disk.blocks[index++] = fileID;
        }
        fileID++;
        if (i < spacesCount) {
            // add free space
            for (int j = 0; j < spaces[i]; j++) {
                disk.blocks[index++] = -1;
            }
        }
    }
    return disk;
}

// Function to find the first free space in the disk
int findFirstFreeSpace(Disk *disk) {
    for (int i = 0; i < disk->size; i++) {
        if (disk->blocks[i] == -1) {
            return i;
        }
    }
    return -1;
}

// Function to find the last file block in the disk
int findLastFileBlock(Disk *disk) {
    for (int i = disk->size - 1; i >= 0; i--) {
        if (disk->blocks[i] != -1) {
            return i;
        }
    }
    return -1;
}

// Function to move one block from one index to another
void moveOneBlock(Disk *disk, int fromIndex, int toIndex) {
    int fileID = disk->blocks[fromIndex];
    disk->blocks[fromIndex] = -1;
    disk->blocks[toIndex] = fileID;
}

// Function to calculate the checksum of the disk
int calculateChecksum(Disk *disk) {
    int checksum = 0;
    for (int i = 0; i < disk->size; i++) {
        if (disk->blocks[i] != -1) {
            checksum += i * disk->blocks[i];
        }
    }
    return checksum;
}

// Function to find the start of a file given a position in the disk
int findFileStart(Disk *disk, int pos) {
    int fileID = disk->blocks[pos];
    while (pos >= 0 && disk->blocks[pos] == fileID) {
        pos--;
    }
    return pos + 1;
}

// Function to find the size of a file given its start position
int findFileSize(Disk *disk, int start) {
    int fileID = disk->blocks[start];
    int size = 0;
    while (start < disk->size && disk->blocks[start] == fileID) {
        size++;
        start++;
    }
    return size;
}

// Function to find the size of free space starting at a position
int findFreeSpaceSize(Disk *disk, int pos) {
    int size = 0;
    while (pos < disk->size && disk->blocks[pos] == -1) {
        size++;
        pos++;
    }
    return size;
}

// Function to move an entire file to a new location
void moveWholeFile(Disk *disk, int fromIndex, int toIndex, int size) {
    int fileID = disk->blocks[fromIndex];

    // Clear the old location
    for (int i = 0; i < size; i++) {
        disk->blocks[fromIndex + i] = -1;
    }

    // Write the file to the new location
    for (int i = 0; i < size; i++) {
        disk->blocks[toIndex + i] = fileID;
    }
}

// Disk compaction function
int compactDisk(const char *diskMap, int part) {
    int *files, *spaces, filesCount, spacesCount;
    parseDiskMap(diskMap, &files, &filesCount, &spaces, &spacesCount);

    Disk disk = createInitialDisk(files, filesCount, spaces, spacesCount);

    if (part == 1) {
        // Block-by-block movement
        while (1) {
            int freeSpace = findFirstFreeSpace(&disk);
            if (freeSpace == -1) break;

            int lastBlock = findLastFileBlock(&disk);
            if (lastBlock == -1 || lastBlock <= freeSpace) break;

            moveOneBlock(&disk, lastBlock, freeSpace);
        }
    } else {
        // Move whole files
        for (int fileID = filesCount - 1; fileID >= 0; fileID--) {
            for (int i = 0; i < disk.size; i++) {
                if (disk.blocks[i] == fileID) {
                    int fileStart = findFileStart(&disk, i);
                    int fileSize = findFileSize(&disk, fileStart);

                    int bestFreeSpace = -1;
                    for (int j = 0; j < fileStart; j++) {
                        if (disk.blocks[j] == -1) {
                            int freeSize = findFreeSpaceSize(&disk, j);
                            if (freeSize >= fileSize) {
                                bestFreeSpace = j;
                                break;
                            }
                            j += freeSize - 1;
                        }
                    }

                    if (bestFreeSpace != -1) {
                        moveWholeFile(&disk, fileStart, bestFreeSpace, fileSize);
                    }
                    break;
                }
            }
        }
    }

    int checksum = calculateChecksum(&disk);

    // Cleanup
    free(files);
    free(spaces);
    free(disk.blocks);

    return checksum;
}

// Function to read the input file and return its contents as a string
char* readInputFile(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        perror("Failed to open file");
        exit(1);
    }

    // Get the file size
    fseek(file, 0, SEEK_END);
    long fileSize = ftell(file);
    fseek(file, 0, SEEK_SET);

    // Allocate memory for the content and read the file
    char *content = (char *)malloc(fileSize + 1);
    fread(content, 1, fileSize, file);
    content[fileSize] = '\0'; // Null-terminate the string

    fclose(file);
    return content;
}

// Main function
int main() {
    // Read input from the file "inputdata.txt"
    char *input = readInputFile("inputdata.txt");

    // Choose which part of the problem to run (1 or 2)
    int part = 1; // Change to 2 for part 2

    // Calculate the checksum and print it
    int checksum = compactDisk(input, part);
    printf("Filesystem Checksum: %d\n", checksum);

    // Cleanup
    free(input);

    return 0;
}
