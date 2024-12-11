#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

#define SIZE 37

typedef struct Node {
    uint64_t value;
    uint64_t count;
    struct Node *next;
} Node;

Node *parse(char *s) {
    Node *next = NULL;
    for (int i = 0; *s; i++) {
        Node *new = malloc(sizeof(Node));
        new->value = strtoul(s, &s, 10);
        new->count = 1;
        new->next = next;
        next = new;
    }
    return next;
}

uint64_t count_digits(uint64_t n) {
    int i = 0;
    for (; n; n /= 10, i++);
    return i;
}

Node *find(Node *head, uint64_t value) {
    if (!head) return NULL;
    if (head->value == value) return head;
    return find(head->next, value);
}

void add(Node *head, uint64_t value, uint64_t count) {
    if (head->next) {
        add(head->next, value, count);
    } else {
        Node *new = malloc(sizeof(Node));
        new->value = value;
        new->count = count;
        new->next = NULL;
        head->next = new;
    }
}

Node *copy(Node *head) {
    if (head == NULL) return NULL;

    Node *next = malloc(sizeof(Node));
    next->value = head->value;
    next->count = head->count;
    next->next = copy(head->next);

    return next;
}

void delete(Node *head) {
    if (head->next) delete(head->next);
    free(head);
}

Node *blink(Node *head) {
    Node *new = copy(head);
    for (Node *n1 = head, *n2 = new; n1; n1 = n1->next, n2 = n2->next) {
        if (n1->count == 0) continue;

        int digits = count_digits(n1->value);
        if (n1->value == 0) {
            Node *node = find(new, 1);
            if (node) node->count += n1->count;
            else add(new, 1, n1->count);
        } else if (digits % 2 == 0) {
            int pow = 1;
            for (int k = 0; k < digits/2; k++) pow *= 10;

            uint64_t lval = n1->value / pow;
            uint64_t rval = n1->value % pow;

            Node *lnode = find(new, lval), *rnode = find(new, rval);

            if (lnode) lnode->count += n1->count;
            else add(new, lval, n1->count);

            if (rnode) rnode->count += n1->count;
            else add(new, rval, n1->count);
        } else {
            uint64_t value = n1->value * 2024;
            Node *node = find(new, value);
            if (node) node->count += n1->count;
            else add(new, value, n1->count);
        }
        n2->count -= n1->count;
    }
    delete(head);
    return new;
}

int main() {
    FILE *file = fopen("inputdata.txt", "r");
    char content[SIZE];
    fread(content, sizeof(char), SIZE, file);
    content[SIZE-1] = 0;

    Node *head = parse(content);
    for (int i = 0; i < 75; i++) {
        head = blink(head);

        continue;
        // debug
        printf("blink %i: ", i);
        for (Node *h = head; h; h = h->next) {
            if (h->count) printf("(%llu, %llu), ", h->value, h->count);
        }
        printf("\b\b \n\n");
    }

    uint64_t count = 0;
    for (Node *h = head; h; h = h->next) {
        count += h->count;
    }
    printf("count: %llu\n", count);
}