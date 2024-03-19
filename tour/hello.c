#include <stdio.h>
#include <stdbool.h>

int main (int argc, char **argv) {
    int sum, j = 0;
//    for (int i = 0; i < 10; i++) {
//        sum += i;
//    }

    for (; j < 10;) {
        sum += j;
        j++;
    }
    printf("sum is %d\n", sum);

//    while (true) {} -> running forever
}
