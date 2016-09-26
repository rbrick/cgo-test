#include <string.h>
#include <stdlib.h>

#include "hello.h"
#include "_cgo_export.h"

typedef unsigned long long int rbrick;

void printHello() {
    // This is normal C code
    printf("Hello!\n");
    
    SayHello("Hello Again!"); 
    
    struct Test_return s = Test();
    
    printf("%d\n", s.r1);
    
    char sUrl[] = "http://rcw.io/";
    
    GoString url = { sUrl, strlen(sUrl) };
    
    char* result = MakeHTTPRequestC(url);
    printf("%s\n", result);

    int fib10 = Fib(10);
    printf("Fibonacci of 10 is %d\n", fib10);

    printf("Fibonacci (Written in C) of 11 is %d\n", fibonacci(11));

    String js = "test";
    printf("String = \"%s\"\n", js);
    int four = add(2, 2);
    printf("Result: %d\n", four);


    int x = 11, y = 10;
    int z = add(x, y);
    printf("Result: %d\n", z);
}

inline int fibonacci(int x) {
    if (x == 1 || x == 0) {
        return 1;
    } else {
        return fibonacci(x - 1) + fibonacci(x - 2);
    }
}


inline 
int add(int x, int y) 
{
    return x + y;    
}
