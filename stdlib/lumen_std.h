#ifndef LUMEN_STD_H
#define LUMEN_STD_H

// Basic I/O
void print(const char* msg);
void println(const char* msg);
int read_int();
double read_float();
char* read_string();

// Arithmetic
int add(int a, int b);
int sub(int a, int b);
int mul(int a, int b);
int div(int a, int b);
int mod(int a, int b);
int powi(int base, int exp);
double powf(double base, double exp);

// Math
double sin(double x);
double cos(double x);
double tan(double x);
double sqrt(double x);
double abs(double x);
int absi(int x);

// String operations
int str_len(const char* str);
char* str_concat(const char* a, const char* b);
int str_cmp(const char* a, const char* b);
char* str_substring(const char* str, int start, int length);

// Time and random
double clock_now();
int random_int(int min, int max);
double random_float(double min, double max);

// Memory management
void* malloc(size_t size);
void free(void* ptr);

#endif
