#pragma once
#include <stdio.h>

#define type_s typedef struct
#define type typedef

type char* String;

type_s {} Lexer;
type_s {} Tokenizer;
type_s {} Parser;

char peek(Tokenizer t);

void printHello();

int fibonacci(int x);