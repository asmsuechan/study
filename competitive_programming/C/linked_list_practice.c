#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

typedef int32_t data_t;

typedef struct cell {
  data_t data;
  struct cell *next;
} cell_t;

cell_t *create_cell(data) {
  cell_t *new_cell = NULL;
  new_cell = (cell_t *)malloc(sizeof(cell_t));
  new_cell -> next = NULL;
  new_cell -> data = data;
  return new_cell;
}

void append_cell(cell_t *head_cell, cell_t *cell) {
  while(head_cell -> next != NULL) {
    head_cell = head_cell -> next;
  }
  head_cell -> next = cell;
}

cell_t *nth_cell(cell_t *head_cell, int n) {
  cell_t *current_cell = head_cell;
  for(int i = 0; i < n; i++) {
    current_cell = current_cell -> next;
  }
  return current_cell;
}

void insert_nth(cell_t *head_cell, cell_t *new_cell, int n) {
  cell_t *inserted_cell = new_cell;

  cell_t *n_minus_1_th_cell = nth_cell(head_cell, n-1);
  cell_t *cell_n = nth_cell(head_cell, n);

  inserted_cell -> next = cell_n;
  n_minus_1_th_cell -> next = inserted_cell;
}

void printCells(cell_t *head_cell) {
  printf("{");
  while(head_cell != NULL) {
    printf(" %d", head_cell -> data);
    head_cell = head_cell -> next;
  }
  printf(" }\n");
}

int main(void) {
  cell_t *head_cell = create_cell(0);
  for(int i = 1; i < 11; i++) {
    append_cell(head_cell, create_cell(i));
  }
  printCells(head_cell);
  insert_nth(head_cell, create_cell(100), 3);
  printCells(head_cell);
  return 0;
}
