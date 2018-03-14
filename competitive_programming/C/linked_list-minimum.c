// This code is based on http://capm-network.com/?tag=C%E8%A8%80%E8%AA%9E%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-%E9%80%A3%E7%B5%90%E3%83%AA%E3%82%B9%E3%83%88

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

typedef int32_t data_t;

typedef struct cell {
  data_t data;
  struct cell *next;
} cell_t;

cell_t *create_cell (data_t data) {
  cell_t *new_cell = NULL;
  new_cell = (cell_t *)malloc(sizeof(cell_t));
  new_cell -> next = NULL;
  new_cell -> data = data;
  return new_cell;
}

void append_cell (cell_t *head_cell, cell_t *new_cell) {
  cell_t *next_cell = NULL;
  cell_t *previous_cell = head_cell;

  next_cell = new_cell;
  while(previous_cell -> next != NULL) {
    previous_cell = previous_cell -> next;
  }
  previous_cell -> next = next_cell;
}

cell_t *nth_cell(cell_t *head_cell, int32_t n)
{
  cell_t *cell = head_cell;
  for (int32_t i = 1; cell != NULL; i++, cell = cell->next)
    if (i == n) break;
  return cell;
}

void insert_nth (cell_t *head_cell, cell_t *new_cell, int32_t n) {
  cell_t *cell_n_minus_1 = nth_cell(head_cell, n-1);
  cell_t *cell_n = nth_cell(head_cell, n);
  cell_n_minus_1 -> next = new_cell;
  new_cell -> next = cell_n;
}

static void list_print(cell_t *head_cell) {
    cell_t *p = head_cell;

    printf("list{ ");
    while(p != NULL){
        printf("%d ", p->data);
        p = p->next;
    }
    printf("}\n");
}

int main (void) {
  cell_t *head_cell = create_cell(0);
  for(int32_t i = 1; i < 100; i += 12) {
    cell_t *new_cell = create_cell(i);
    append_cell(head_cell, new_cell);
  }
  list_print(head_cell);
  printf("%d\n", nth_cell(head_cell, 8)->data);
  cell_t *new_cell = create_cell(1000);
  insert_nth(head_cell, new_cell, 3);
  list_print(head_cell);
  return 0;
}
