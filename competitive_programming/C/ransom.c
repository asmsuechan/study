#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(){
    int m;
    int n;
    scanf("%d %d",&m,&n);
    char* *magazine = malloc(sizeof(char*) * m);
    for(int magazine_i = 0; magazine_i < m; magazine_i++){
       magazine[magazine_i] = (char *)malloc(5120* sizeof(char));
       scanf("%s",magazine[magazine_i]);
    }
    char* *ransom = malloc(sizeof(char*) * n);
    for(int ransom_i = 0; ransom_i < n; ransom_i++){
       ransom[ransom_i] = (char *)malloc(5120* sizeof(char));
       scanf("%s",ransom[ransom_i]);
    }

    int is_valid = 1;
    for(int ransom_i = 0; ransom_i < n && is_valid; ransom_i++){
      for(int magazine_i = m-1; magazine_i >= 0; magazine_i--){
        if(strcmp(ransom[ransom_i], magazine[magazine_i]) == 0) {
          magazine[magazine_i] = "";
          break;
        }
        if(magazine_i == 0) {
          is_valid = 0;
        }
      }
    }
    if(is_valid) {
      printf("Yes");
    } else {
      printf("No");
    }
    return 0;
}

