#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 100

// ButtonMove represents the movement of a button
typedef struct {
  char name;
  int tokenCost;
  int xMove;
  int yMove;
} ButtonMove;

// Machine represents a single claw machine
typedef struct {
  ButtonMove buttonA;
  ButtonMove buttonB;
  int prizeX;
  int prizeY;
} Machine;

// Function prototypes
int gcd(int a, int b);
int calculateMinTokens(Machine machine);
Machine parseMachineFromFile(FILE* file);

int main() {
  FILE* inputFile = fopen("inputdata.txt", "r");
  if (inputFile == NULL) {
    printf("Error opening input file!\n");
    return 1;
  }

  Machine machine;
  int totalTokens = 0;
  while ((machine = parseMachineFromFile(inputFile)).prizeX != 0) { // Check if prizeX is 0 (indicating end of file)
    int tokens = calculateMinTokens(machine);
    printf("Machine - Minimum tokens: %d\n", tokens);
    totalTokens += tokens;
  }

  printf("\nTotal minimum tokens for all machines: %d\n", totalTokens);

  fclose(inputFile);
  return 0;
}

// gcd calculates the Greatest Common Divisor using Euclidean algorithm
int gcd(int a, int b) {
  while (b != 0) {
    int t = b;
    b = a % b;
    a = t;
  }
  return a;
}

// calculateMinTokens finds the minimum tokens to reach the prize
int calculateMinTokens(Machine machine) {
  int minTokens = 2147483647; // Maximum value for an integer (32-bit signed)

  // Use GCD to optimize the search space
  int gcdX = gcd(machine.buttonA.xMove, machine.buttonB.xMove);
  int gcdY = gcd(machine.buttonA.yMove, machine.buttonB.yMove);

  // Determine maximum reasonable search iterations using GCD
  int maxIterations = (machine.prizeX / gcdX + machine.prizeY / gcdY) * 2;

  for (int a = 0; a <= maxIterations; a++) {
    for (int b = 0; b <= maxIterations; b++) {
      // Calculate total X and Y movements
      int totalX = a * machine.buttonA.xMove + b * machine.buttonB.xMove;
      int totalY = a * machine.buttonA.yMove + b * machine.buttonB.yMove;

      // Check if we've reached the prize exactly
      if (totalX == machine.prizeX && totalY == machine.prizeY) {
        // Calculate total tokens spent
        int tokens = a * machine.buttonA.tokenCost + b * machine.buttonB.tokenCost;

        // Update minimum tokens if found
        minTokens = tokens < minTokens ? tokens : minTokens;
      }
    }
  }

  // If no solution found
  if (minTokens == 2147483647) {
    printf("No solution found for machine with prize X=%d, Y=%d\n", machine.prizeX, machine.prizeY);
    return 0;
  }

  return minTokens;
}

// parseMachineFromFile reads a single machine configuration from the file
Machine parseMachineFromFile(FILE* file) {
  Machine machine = {0}; // Initialize all fields to 0

  char line[MAX_LINE_LENGTH];
  int numLinesRead = 0;

  while (fgets(line, MAX_LINE_LENGTH, file) != NULL && numLinesRead < 3) {
    if (line[0] == '\n' || line[0] == '\r') {
      continue; // Skip empty lines
    }

    // Parse Button A
    if (strstr(line, "Button A: X+") != NULL) {
      sscanf(line, "Button A: X+%d, Y+%d", &machine.buttonA.xMove, &machine.buttonA.yMove);
      machine.buttonA.name = 'A';
      machine.buttonA.tokenCost = 3;
    }

    // Parse Button B
    if (strstr(line, "Button B: X+") != NULL) {
      sscanf(line, "Button B: X+%d, Y+%d", &machine.buttonB.xMove, &machine.buttonB.yMove);
      machine.buttonB.name = 'B';
      machine.buttonB.tokenCost = 1;
    }

    // Parse Prize
    if (strstr(line, "Prize: X=") != NULL) {
      sscanf(line, "Prize: X=%d, Y=%d", &machine.prizeX, &machine.prizeY);
    }

    numLinesRead++;
  }

  // If all three lines for a machine were read, return the machine
  if (numLinesRead == 3) {
    return machine;
  } else {
    // If less than 3 lines were read, it means we reached the end of the file
    machine.prizeX = 0; // Set prizeX to 0 to signal end of file
    return machine; 
  }
}