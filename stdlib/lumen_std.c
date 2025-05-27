#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <time.h>
#include "lumen_std.h"

// Basic I/O
void print(const char* msg) {
    printf("%s", msg);
}

void println(const char* msg) {
    printf("%s\n", msg);
}

int read_int() {
    int val;
    scanf("%d", &val);
    return val;
}

double read_float() {
    double val;
    scanf("%lf", &val);
    return val;
}

char* read_string() {
    char* buffer = malloc(1024);
    scanf("%1023s", buffer);
    return buffer;
}

// Arithmetic
int add(int a, int b) { return a + b; }
int sub(int a, int b) { return a - b; }
int mul(int a, int b) { return a * b; }
int div(int a, int b) { return a / b; }
int mod(int a, int b) { return a % b; }

int powi(int base, int exp) {
    int result = 1;
    for (int i = 0; i < exp; i++) {
        result *= base;
    }
    return result;
}

double powf(double base, double exp) {
    return pow(base, exp);
}

// Math
double sin(double x) { return sin(x); }
double cos(double x) { return cos(x); }
double tan(double x) { return tan(x); }
double sqrt(double x) { return sqrt(x); }
double abs(double x) { return fabs(x); }
int absi(int x) { return abs(x); }

// String operations
int str_len(const char* str) {
    return strlen(str);
}

char* str_concat(const char* a, const char* b) {
    int len_a = strlen(a);
    int len_b = strlen(b);
    char* result = malloc(len_a + len_b + 1);
    strcpy(result, a);
    strcat(result, b);
    return result;
}

int str_cmp(const char* a, const char* b) {
    return strcmp(a, b);
}

char* str_substring(const char* str, int start, int length) {
    char* result = malloc(length + 1);
    strncpy(result, str + start, length);
    result[length] = '\0';
    return result;
}

// Time and random
double clock_now() {
    return (double)clock() / CLOCKS_PER_SEC;
}

int random_int(int min, int max) {
    return min + rand() % (max - min + 1);
}

double random_float(double min, double max) {
    return min + ((double)rand() / RAND_MAX) * (max - min);
}

// Initialize random seed
__attribute__((constructor))
static void init_random() {
    srand(time(NULL));
}
