#include <stdio.h>
#include "goclib.h"

void manyCalls(int count, int argument)
{
    printf("Calling ... ");
    for(int i = 0; i < count; ++i) {
        printf("%d ", i + 1);
        fflush(stdout);
        Fibonacci(argument);
    }
    printf(". Done.\n");
}

int main(int argc, char **argv) 
{
    printf("Hello from C\n");
    int n = 32;
    printf("fib(%d)=%lld\n", n, Fibonacci(n));
    manyCalls(1000, n);
    return 0;
}
