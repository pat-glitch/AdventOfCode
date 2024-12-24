#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_NODES 1000
#define MAX_NAME_LEN 10

// Adjacency list representation
typedef struct Node {
    char name[MAX_NAME_LEN];
    struct Node *next;
} Node;

typedef struct {
    char name[MAX_NAME_LEN];
    Node *head;
} AdjList;

AdjList graph[MAX_NODES];
int nodeCount = 0;

// Function to find index of a node, or add it if not found
int getNodeIndex(char *name) {
    for (int i = 0; i < nodeCount; i++) {
        if (strcmp(graph[i].name, name) == 0)
            return i;
    }
    strcpy(graph[nodeCount].name, name);
    graph[nodeCount].head = NULL;
    return nodeCount++;
}

// Function to add an edge
void addEdge(char *src, char *dest) {
    int srcIndex = getNodeIndex(src);
    int destIndex = getNodeIndex(dest);

    // Add dest to src's adjacency list
    Node *newNode = (Node *)malloc(sizeof(Node));
    strcpy(newNode->name, dest);
    newNode->next = graph[srcIndex].head;
    graph[srcIndex].head = newNode;

    // Add src to dest's adjacency list
    newNode = (Node *)malloc(sizeof(Node));
    strcpy(newNode->name, src);
    newNode->next = graph[destIndex].head;
    graph[destIndex].head = newNode;
}

// Check if two nodes are connected
int isConnected(char *a, char *b) {
    int index = getNodeIndex(a);
    Node *temp = graph[index].head;
    while (temp) {
        if (strcmp(temp->name, b) == 0)
            return 1;
        temp = temp->next;
    }
    return 0;
}

// Check if a name starts with 't'
int startsWithT(char *name) {
    return name[0] == 't';
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        printf("Error opening file.\n");
        return 1;
    }

    char src[MAX_NAME_LEN], dest[MAX_NAME_LEN];
    while (fscanf(file, "%[^-]-%s\n", src, dest) != EOF) {
        addEdge(src, dest);
    }
    fclose(file);

    int triangleCount = 0;

    // Find triangles
    for (int i = 0; i < nodeCount; i++) {
        char *a = graph[i].name;

        Node *node1 = graph[i].head;
        while (node1) {
            char *b = node1->name;
            if (strcmp(a, b) < 0) {  // Enforce a < b
                Node *node2 = node1->next;
                while (node2) {
                    char *c = node2->name;
                    if (strcmp(b, c) < 0 && strcmp(a, c) < 0) {  // Enforce b < c and a < c
                        // Check if these three form a triangle
                        if (isConnected(b, c)) {
                            // Check if one of the nodes starts with 't'
                            if (startsWithT(a) || startsWithT(b) || startsWithT(c)) {
                                printf("Triangle: %s, %s, %s\n", a, b, c);
                                triangleCount++;
                            }
                        }
                    }
                    node2 = node2->next;
                }
            }
            node1 = node1->next;
        }
    }

    printf("Number of triangles with at least one 't': %d\n", triangleCount);
    return 0;
}
