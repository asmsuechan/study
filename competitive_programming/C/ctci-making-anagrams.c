#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>
#include <stdint.h>

int main(){
    char* a = (char *)malloc(512000 * sizeof(char));
    scanf("%s",a);
    char* b = (char *)malloc(512000 * sizeof(char));
    scanf("%s",b);

    printf("%i %i %i ", a[0]-'a', a[1], a[2]);
    int alphabet[26] = {0};
    for(int i = 0; i < 4; i++) {
      alphabet[a[i] - 'a']++;
      alphabet[b[i] - 'a']--;
    }
    int sum = 0;
    for(int i = 0; i < 26; i++) {
      sum += alphabet[i];
    }
    printf("%d", sum);
    return 0;
}
