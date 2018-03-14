#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
    int n;
    int k;
    scanf("%d %d",&n,&k);
    int *a = malloc(sizeof(int) * n);
    for(int a_i = 0; a_i < n; a_i++){
       scanf("%d",&a[a_i]);
    }

    int* newArray = malloc(sizeof(int) * n);
    int c = 0;
    for(int i = k; i < n; i++) {
      newArray[c] = a[i];
      c++;
    }
    for(int i = 0; i < k; i++) {
      newArray[c] = a[i];
      c++;
    }
    for(int i = 0; i < n; i++){
      printf("%d ", newArray[i]);
    }
    return 0;
}
