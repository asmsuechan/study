#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>
#include <stdint.h>

void add() {
  int32_t a = 0;
  int32_t b = 1;
  int32_t c = a + b;
  printf("%d", c);
}

int main(){
  add();
  return 0;
}
