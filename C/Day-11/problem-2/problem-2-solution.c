#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

#define SIZE 37

// Linked list node definition
typedef struct Node {
    uint64_t value;
    uint64_t count;
    struct Node *next;
} Node;

// Function to count the number of space-separated integers in the input string
int count(char *s) {
    int count = 1;
    while (*s++) {
        if (*s == ' ') count++;
    }
    return count;
}

// Function to convert the input string to a linked list of nodes
Node *parse(char *s) {
    Node *head = NULL;
    Node *tail = NULL;

    while (*s) {
        Node *new_node = malloc(sizeof(Node));
        new_node->value = strtoul(s, &s, 10);
        new_node->count = 1;
        new_node->next = NULL;

        if (head == NULL) {
            head = new_node;
            tail = new_node;
        } else {
            tail->next = new_node;
            tail = new_node;
        }
    }

    return head;
}

// Function to count the number of digits in a number
uint64_t count_digits(uint64_t n) {
    int i = 0;
    while (n) {
        n /= 10;
        i++;
    }
    return i;
}

// Function to find a node by its value in the linked list
Node *find(Node *head, uint64_t value) {
    while (head) {
        if (head->value == value) return head;
        head = head->next;
    }
    return NULL;
}

// Function to add a new node to the end of the linked list
void add(Node *head, uint64_t value, uint64_t count) {
    Node *new_node = malloc(sizeof(Node));
    new_node->value = value;
    new_node->count = count;
    new_node->next = NULL;

    while (head->next) {
        head = head->next;
    }
    head->next = new_node;
}

// Function to copy a linked list
Node *copy(Node *head) {
    if (!head) return NULL;
    Node *new_head = malloc(sizeof(Node));
    new_head->value = head->value;
    new_head->count = head->count;
    new_head->next = copy(head->next);
    return new_head;
}

// Function to delete the linked list and free memory
void delete(Node *head) {
    if (head) {
        delete(head->next);
        free(head);
    }
}

// Function to perform a blink (transformation) on the linked list
Node *blink(Node *head) {
    Node *new_head = copy(head);

    // Process each stone in the original list
    for (Node *orig = head, *copy_node = new_head; orig; orig = orig->next, copy_node = copy_node->next) {
        if (orig->count == 0) continue; // Skip if there are no stones left for this value

        int digits = count_digits(orig->value);
        if (orig->value == 0) {
            // Zero becomes 1
            Node *node = find(new_head, 1);
            if (node) node->count += orig->count;
            else add(new_head, 1, orig->count);
        } else if (digits % 2 == 0) {
            // Split the stone value with an even number of digits
            int pow = 1;
            for (int i = 0; i < digits / 2; i++) pow *= 10;

            uint64_t left_val = orig->value / pow;
            uint64_t right_val = orig->value % pow;
            Node *left_node = find(new_head, left_val);
            Node *right_node = find(new_head, right_val);

            if (left_node) left_node->count += orig->count;
            else add(new_head, left_val, orig->count);

            if (right_node) right_node->count += orig->count;
            else add(new_head, right_val, orig->count);
        } else {
            // Multiply the stone value with odd digits by 2024
            uint64_t new_value = orig->value * 2024;
            Node *node = find(new_head, new_value);
            if (node) node->count += orig->count;
            else add(new_head, new_value, orig->count);
        }

        copy_node->count -= orig->count;
    }

    delete(head); // Free the old list
    return new_head; // Return the new list after transformation
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

    // Parse the input into a linked list
    Node *head = parse(content);

    // Perform 75 blinks
    for (int i = 0; i < 75; i++) {
        head = blink(head);
    }

    // Calculate the total number of stones after 75 blinks
    uint64_t total_count = 0;
    for (Node *n = head; n; n = n->next) {
        total_count += n->count;
    }

    printf("Total stone count after 75 blinks: %llu\n", total_count);

    // Free the linked list memory
    delete(head);

    return 0;
}
