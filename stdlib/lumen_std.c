#include "lumen_std.h"
#include <stdio.h>
#include <math.h>
#include <time.h>

void print(const char* msg) {
    printf("%s", msg);
}

void println(const char* msg) {
    printf("%s\n", msg);
}

int add(int a, int b) {
    return a + b;
}

int mul(int a, int b) {
    return a * b;
}

int powi(int base, int exp) {
    return (int)pow(base, exp);
}

double clock_now() {
    return (double)clock() / CLOCKS_PER_SEC;
}
