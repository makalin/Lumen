# 🎲 Lumen Fake Data Generator

A library for generating realistic fake data in Lumen. Perfect for testing, development, and creating sample data.

## Features

- Generate random first names
- Generate random integers within a range
- Generate random boolean values
- Generate random dates in YYYY-MM-DD format
- Memory-safe string handling
- Thread-safe random number generation

## Usage

First, initialize the random number generator:

```llvm
call void @init_random()
```

### Generate a Random Name

```llvm
%name = call i8* @generate_first_name()
; Don't forget to free the memory when done
call void @free(i8* %name)
```

### Generate a Random Integer

```llvm
; Generate a number between 1 and 100
%number = call i32 @generate_int(i32 1, i32 100)
```

### Generate a Random Boolean

```llvm
%bool_val = call i1 @generate_bool()
```

### Generate a Random Date

```llvm
%date = call i8* @generate_date()
; Don't forget to free the memory when done
call void @free(i8* %date)
```

## Building

To use this library in your Lumen project:

1. Copy the `fake_data.ll` file to your project
2. Link it with your main LLVM IR file:

```bash
llc fake_data.ll -filetype=obj -o fake_data.o
llc your_program.ll -filetype=obj -o your_program.o
clang your_program.o fake_data.o -o your_program
```

## Memory Management

The library uses dynamic memory allocation for strings. Remember to free any strings returned by the library functions when you're done with them using the `free` function.

## Future Improvements

- Add more data types (last names, addresses, emails, etc.)
- Add support for custom data patterns
- Add support for generating arrays of fake data
- Add support for generating structured data (JSON-like objects) 