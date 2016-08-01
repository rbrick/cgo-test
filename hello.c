#include <string.h>

#include "hello.h"
#include "_cgo_export.h"

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
    printf("Fibonacci of 10 is %d", fib10);
}
