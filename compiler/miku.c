#include <stdio.h>
#include <stdbool.h>
#include <string.h>

char usrin[30];
int i;

int start() { // This handles the message you see at the start and user input.
    printf("miku-c\n");
    printf("A Compiler built for turning .frog files into valid python configs\n");
    printf("===============\n");
    scanf("%s", usrin);
    return 0;
}

int read() { // This reads user input and handles it.
    for (i = 0; if (i < usrinlen) {i++}) { // This pos seems to want to not work!!!
        printf("%c\n", usrin[i]);
    }
}

int main() {
    int usrinlen = strlen(usrin);
    start();
    printf("input: %s", usrin);
    return 0;
}