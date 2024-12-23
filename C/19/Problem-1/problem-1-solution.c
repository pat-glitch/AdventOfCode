#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Represents consuming no char for a state
enum { START_STATE = '\0', OR_STATE = '\r' };

typedef struct {
    char match;
    int next_state;
    int or_next_state; // Only used for OR states
} state;

int add_string(state *state_arr, int state_count, int start_state, char *str) {
    if (state_count == 1) {
        state_arr[start_state].next_state = state_count;
    } else {
        state_arr[state_count].match = OR_STATE;
        state_arr[state_count].or_next_state = state_arr[start_state].next_state;
        state_arr[state_count].next_state = state_count + 1;

        state_arr[start_state].next_state = state_count;

        ++state_count;
    }

    for (int i = 0; str[i] != '\0'; ++i) {
        state_arr[state_count].match = str[i];
        state_arr[state_count].next_state = state_count + 1;

        ++state_count;
    }

    state_arr[state_count - 1].next_state = start_state;

    return state_count;
}

int match_char(char cur_char, state *state_arr, int *cur_char_states, int *next_char_states, int *next_char_state_count, int cur_char_state_count) {
    int start_added = 0;
    *next_char_state_count = 0;

    for (int i = 0; i < cur_char_state_count; ++i) {
        if (state_arr[cur_char_states[i]].match == START_STATE) {
            cur_char_states[cur_char_state_count] = state_arr[cur_char_states[i]].next_state;
            ++cur_char_state_count;
        } else if (state_arr[cur_char_states[i]].match == OR_STATE) {
            cur_char_states[cur_char_state_count] = state_arr[cur_char_states[i]].next_state;
            ++cur_char_state_count;

            cur_char_states[cur_char_state_count] = state_arr[cur_char_states[i]].or_next_state;
            ++cur_char_state_count;
        } else if (cur_char == state_arr[cur_char_states[i]].match) {
            if (state_arr[cur_char_states[i]].next_state == 0) {
                if (start_added == 0) {
                    start_added = 1;
                    next_char_states[*next_char_state_count] = state_arr[cur_char_states[i]].next_state;
                    *next_char_state_count += 1;
                }
            } else {
                next_char_states[*next_char_state_count] = state_arr[cur_char_states[i]].next_state;
                *next_char_state_count += 1;
            }
        }
    }

    return start_added;
}

char *read_line(FILE *file) {
    size_t size = 128;
    size_t len = 0;
    char *buffer = malloc(size);

    if (!buffer) {
        perror("Unable to allocate buffer");
        return NULL;
    }

    int ch;
    while ((ch = fgetc(file)) != EOF && ch != '\n') {
        if (len + 1 >= size) {
            size *= 2;
            char *new_buffer = realloc(buffer, size);
            if (!new_buffer) {
                perror("Unable to reallocate buffer");
                free(buffer);
                return NULL;
            }
            buffer = new_buffer;
        }
        buffer[len++] = ch;
    }
    if (len == 0 && ch == EOF) {
        free(buffer);
        return NULL; // No data read, EOF reached
    }

    buffer[len] = '\0';
    return buffer;
}

int main(void) {
    enum { MATCHES, INPUTS };
    enum { START = 0 };

    int cur_part = MATCHES;
    char *input_line = NULL;

    state state_arr[3000];
    int state_count = 0;

    state_arr[START].match = START_STATE;
    state_arr[START].next_state = START;
    ++state_count;

    char *temp_match = NULL;
    char delim[] = ", \n";

    int cur_char_states[3000];
    int next_char_states[3000];
    int next_char_state_count = 0;
    int cur_char_state_count = 0;

    int valid_pat_count = 0;

    // Open the input file
    FILE *file = fopen("inputdata.txt", "r");
    if (!file) {
        perror("Error opening file");
        return EXIT_FAILURE;
    }

    while ((input_line = read_line(file)) != NULL) {
        switch (cur_part) {
        case MATCHES:
            temp_match = strtok(input_line, delim);

            while (temp_match != NULL) {
                state_count = add_string(state_arr, state_count, START, temp_match);

                temp_match = strtok(NULL, delim);
            }

            cur_part = INPUTS;
            break;
        case INPUTS:
            int valid_match = 0;
            int *arr_ptr = cur_char_states;
            int *arr_ptr_2 = next_char_states;
            int *temp_ptr = NULL;

            cur_char_state_count = 0;
            next_char_state_count = 0;

            arr_ptr[cur_char_state_count] = START;

            ++cur_char_state_count;

            for (int i = 0; input_line[i] != '\0'; ++i) {
                valid_match = match_char(input_line[i], state_arr, arr_ptr, arr_ptr_2, &next_char_state_count, cur_char_state_count);

                cur_char_state_count = next_char_state_count;

                temp_ptr = arr_ptr;
                arr_ptr = arr_ptr_2;
                arr_ptr_2 = temp_ptr;
            }

            if (valid_match == 1) {
                ++valid_pat_count;
            }

            break;
        }

        free(input_line);
    }

    printf("%d\n", valid_pat_count);

    fclose(file);
    return 0;
}
