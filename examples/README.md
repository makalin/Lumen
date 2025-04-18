## Examples

Below are a handful of illustrative Lumen programs you can find in the `examples/` directory. They demonstrate core language features, standard‑library functions, and common programming patterns.

### hello.lumen
A minimal “Hello, world!” program that shows how to define `main()`, invoke `println()`, and return an exit code.  
```lumen
func main() -> int {
    println("Hello, Lumen!")
    return 0
}
```

---

### arithmetic_example.lumen
Demonstrates basic integer arithmetic and timing via the stdlib’s `add()`, `mul()`, `powi()` and `clock_now()`:  
- Declares and initializes variables with explicit types.  
- Prints results of addition, multiplication, and exponentiation.  
- Measures the runtime of a simple loop summing numbers.

---

### factorial.lumen
Implements a recursive `factorial(n)` using `mul()`, then prints the factorial of 10. Shows:  
- User‑defined functions with typed parameters  
- Recursion and conditional returns  
```lumen
func factorial(n: int) -> int { … }
func main() -> int { … }
```

---

### fib.lumen
Prints the first 10 Fibonacci numbers using a naive recursive `fib(n)` built on top of `add()`. Highlights:  
- Recursion with multiple branches  
- Looping with a `for i in 0..10` range  
- String interpolation via `to_string()`

---

### pi_approx.lumen
Approximates π using the Leibniz series (`π = 4 · Σ ((–1)^i / (2i + 1))`) over 500 000 terms, and times the computation with `clock_now()`. Illustrates:  
- Floating‑point arithmetic (`double`)  
- Alternating signs via `%` and conditional expressions  
- Performance measurement

---

### gcd.lumen
Calculates the greatest common divisor of two integers (`48` and `18`) using Euclid’s algorithm with a `while` loop and the modulus operator `%`. Demonstrates:  
- Mutable variables (`var`)  
- Looping until a condition is met  
- Simple imperative control flow

---

Feel free to run any of these with the hypothetical compiler toolchain:

```bash
lumenc examples/<name>.lumen -o <name>.ll
llc <name>.ll -filetype=obj -o <name>.o
clang <name>.o lumen_std.o -o <name> && ./<name>
```

Each example is self‑contained and showcases a different aspect of Lumen’s syntax and standard library.