#include <stdio.h> 
#include <setjmp.h>
 
jmp_buf buf; 

void banana(void)
{ 
    printf("in banana() \n"); 
    printf("you'll never see this,because i longjmp'd"); 
    longjmp(buf,1);  

    int x = 10;
} 

int main() 
{ 
    if(setjmp(buf)) 
        printf("back in main\n"); 
    else
    { 
        printf("first time through\n"); 
        banana(); 
    } 

    return 0;
}