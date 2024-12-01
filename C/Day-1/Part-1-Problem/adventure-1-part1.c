#include <stdio.h>
#include <stdlib.h>

// Function to sort an array in ascending order
void sortArray(int *arr, int size) {
    for (int i = 0; i < size - 1; i++) {
        for (int j = i + 1; j < size; j++) {
            if (arr[i] > arr[j]) {
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
    }
}

int main() {
    char filename[100];
    printf("Enter the name of the input file: ");
    scanf("%s", filename);

    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        perror("Error opening file");
        return 1;
    }

    // Maximum size of lists based on estimated file content
    int list1[1000], list2[1000];
    int size = 0;

    // Read numbers from the file
    while (!feof(file)) {
        int num1, num2;
        if (fscanf(file, "%d %d", &num1, &num2) == 2) {
            list1[size] = num1;
            list2[size] = num2;
            size++;
        }
    }
    fclose(file);

    // Sort both lists
    sortArray(list1, size);
    sortArray(list2, size);

    // Create the third list and calculate the sum of differences
    int diffList[size];
    int sum = 0;
    for (int i = 0; i < size; i++) {
        diffList[i] = abs(list1[i] - list2[i]);
        sum += diffList[i];
    }

    // Output the results
    printf("Sorted first list: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", list1[i]);
    }
    printf("\n");

    printf("Sorted second list: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", list2[i]);
    }
    printf("\n");

    printf("Difference list: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", diffList[i]);
    }
    printf("\n");

    printf("Sum of the difference list: %d\n", sum);

    return 0;
}
