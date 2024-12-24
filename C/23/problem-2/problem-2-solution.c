#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define MAX_NODES 1000
#define MAX_NAME_LEN 10

// Structure for adjacency list
typedef struct {
    char name[MAX_NAME_LEN];
    bool neighbors[MAX_NODES];
} Node;

Node nodes[MAX_NODES];
int nodeCount = 0;

// Function to find index of a node, or add it if not found
int getNodeIndex(char *name) {
    for (int i = 0; i < nodeCount; i++) {
        if (strcmp(nodes[i].name, name) == 0)
            return i;
    }
    strcpy(nodes[nodeCount].name, name);
    memset(nodes[nodeCount].neighbors, 0, sizeof(nodes[nodeCount].neighbors));
    return nodeCount++;
}

// Add an undirected edge between two nodes
void addEdge(char *a, char *b) {
    int indexA = getNodeIndex(a);
    int indexB = getNodeIndex(b);
    nodes[indexA].neighbors[indexB] = true;
    nodes[indexB].neighbors[indexA] = true;
}

// Helper functions for dynamic arrays
char **copyArray(char **src, int size) {
    char **copy = (char **)malloc(size * sizeof(char *));
    for (int i = 0; i < size; i++) {
        copy[i] = strdup(src[i]);
    }
    return copy;
}

void freeArray(char **arr, int size) {
    for (int i = 0; i < size; i++) {
        free(arr[i]);
    }
    free(arr);
}

// Bron-Kerbosch algorithm to find maximal cliques
void bronKerbosch(char **r, int rSize, char **p, int pSize, char **x, int xSize, bool adjMatrix[MAX_NODES][MAX_NODES], char ****cliques, int *cliqueCount) {
    if (pSize == 0 && xSize == 0) {
        // Found a maximal clique
        (*cliques) = realloc(*cliques, (*cliqueCount + 1) * sizeof(char **));
        (*cliques)[*cliqueCount] = copyArray(r, rSize);
        (*cliqueCount)++;
        return;
    }

    for (int i = 0; i < pSize; i++) {
        char *node = p[i];
        char **newR = copyArray(r, rSize + 1);
        newR[rSize] = strdup(node);

        char **newP = NULL;
        char **newX = NULL;
        int newPSize = 0, newXSize = 0;

        // Build newP and newX
        for (int j = 0; j < pSize; j++) {
            if (adjMatrix[getNodeIndex(node)][getNodeIndex(p[j])]) {
                newP = realloc(newP, (newPSize + 1) * sizeof(char *));
                newP[newPSize++] = strdup(p[j]);
            }
        }
        for (int j = 0; j < xSize; j++) {
            if (adjMatrix[getNodeIndex(node)][getNodeIndex(x[j])]) {
                newX = realloc(newX, (newXSize + 1) * sizeof(char *));
                newX[newXSize++] = strdup(x[j]);
            }
        }

        bronKerbosch(newR, rSize + 1, newP, newPSize, newX, newXSize, adjMatrix, cliques, cliqueCount);

        // Free allocated memory
        freeArray(newR, rSize + 1);
        freeArray(newP, newPSize);
        freeArray(newX, newXSize);

        // Update P and X
        memmove(&p[i], &p[i + 1], (pSize - i - 1) * sizeof(char *));
        pSize--;
        x = realloc(x, (xSize + 1) * sizeof(char *));
        x[xSize++] = strdup(node);
        i--;
    }
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        printf("Error opening file.\n");
        return 1;
    }

    char line[256];
    bool adjMatrix[MAX_NODES][MAX_NODES] = {0};

    // Read input and build the adjacency matrix
    while (fgets(line, sizeof(line), file)) {
        char *a = strtok(line, "-");
        char *b = strtok(NULL, "\n");
        if (a && b) {
            addEdge(a, b);
            adjMatrix[getNodeIndex(a)][getNodeIndex(b)] = true;
            adjMatrix[getNodeIndex(b)][getNodeIndex(a)] = true;
        }
    }
    fclose(file);

    // Get all nodes
    char **allNodes = (char **)malloc(nodeCount * sizeof(char *));
    for (int i = 0; i < nodeCount; i++) {
        allNodes[i] = strdup(nodes[i].name);
    }

    // Find all maximal cliques
    char ***cliques = NULL;
    int cliqueCount = 0;
    bronKerbosch(NULL, 0, allNodes, nodeCount, NULL, 0, adjMatrix, &cliques, &cliqueCount);

    // Find the largest clique
    char **largestClique = NULL;
    int largestSize = 0;
    for (int i = 0; i < cliqueCount; i++) {
        int size = 0;
        while (cliques[i][size]) size++;
        if (size > largestSize) {
            largestSize = size;
            largestClique = cliques[i];
        }
    }

    // Sort the largest clique alphabetically
    if (largestClique) {
        qsort(largestClique, largestSize, sizeof(char *), (int (*)(const void *, const void *))strcmp);
    }

    // Generate the password
    printf("Password to get into the LAN party: ");
    for (int i = 0; i < largestSize; i++) {
        printf("%s%s", largestClique[i], (i == largestSize - 1) ? "\n" : ",");
    }

    // Free allocated memory
    for (int i = 0; i < cliqueCount; i++) {
        freeArray(cliques[i], largestSize);
    }
    free(cliques);
    freeArray(allNodes, nodeCount);

    return 0;
}
