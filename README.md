# 🌙 Lumen Programming Language

**Lumen** is a statically typed, compiled programming language designed for performance and flexibility. It compiles to native code via LLVM and supports WebAssembly (WASM) out of the box — making it perfect for both **system-level** and **frontend** development.

> 🔥 Write once, run anywhere — fast.

---

## 🚀 Features

- ⚡ **Native compilation** (via LLVM)
- 🌐 **WebAssembly support** for frontend or embedded use
- 🧠 Simple, safe syntax with strong typing
- 📚 Cross-platform **stdlib** (works on desktop and browser)
- 🦫 **Custom compiler in Go** (lexer, parser, AST)
- 🎨 **Virtual DOM** for frontend development
- 🎮 **Interactive playground** for learning and testing
- 🖥️ **High-performance HTTP server** with event-driven architecture
- 🔧 Future plans: module system, generics, async support

---

## 📁 Folder Structure

```
lumen/
├── stdlib/
│   ├── lumen_std.c         # Standard library C code
│   ├── lumen_std.h         # Header file for native/WASM use
│   └── wasm_imports.js     # JS glue for WebAssembly
├── main.ll                 # LLVM IR example using stdlib
├── examples/
│   ├── hello.lumen         # Basic "Hello, World!" example
│   ├── arithmetic_example.lumen  # Arithmetic operations
│   ├── factorial.lumen     # Recursive factorial
│   ├── fib.lumen           # Fibonacci sequence
│   ├── gcd.lumen           # Greatest Common Divisor
│   ├── pi_approx.lumen     # Pi approximation
│   ├── raytracer/          # 3D ray tracer (native + WASM)
│   │   ├── raytracer.lumen # Lumen source version
│   │   ├── raytracer.ll    # LLVM IR version
│   │   ├── raytracer.html  # Web interface
│   │   └── raytracer.js    # JavaScript integration
│   ├── virtual_dom/        # Virtual DOM implementation
│   │   └── vdom.lumen      # Frontend framework
│   ├── playground/         # Interactive language playground
│   │   ├── playground.html # Web-based code editor
│   │   └── playground.js   # Playground functionality
│   └── http_server/        # High-performance HTTP server
│       ├── server.ll       # LLVM IR server implementation
│       └── README.md       # Server documentation
├── compiler/
│   └── src/
│       ├── lexer.go        # Go lexer for Lumen
│       ├── parser.go       # Go parser for Lumen
│       ├── ast.go          # AST node definitions
│       └── tokens.go       # Token types and struct
```

---

## 🎮 Quick Start

### Try the Playground
Open `examples/playground/playground.html` in your browser to write and run Lumen code interactively!

### Run Examples
```bash
# Basic examples
lumen examples/hello.lumen -o hello
./hello

# Advanced examples
lumen examples/raytracer/raytracer.lumen -o raytracer
./raytracer

# WebAssembly
lumen examples/raytracer/raytracer.lumen -target wasm -o raytracer.wasm
```

---

## 🔧 Build Instructions

### Native (Linux/macOS)

```bash
clang -c stdlib/lumen_std.c -o lumen_std.o
llc main.ll -filetype=obj -o main.o
clang main.o lumen_std.o -o lumen_exec
./lumen_exec
```

### WebAssembly (with Emscripten)

```bash
emcc stdlib/lumen_std.c -O3 -s SIDE_MODULE=1 -o lumen_std.wasm
llc main.ll -march=wasm32 -filetype=obj -o main.o
wasm-ld main.o lumen_std.wasm -o lumen_combined.wasm --no-entry --export-all
```

Then load `lumen_combined.wasm` in the browser with `wasm_imports.js`.

---

## 🦫 Compiler (Go)

A new Lumen compiler is being developed in Go.  
Source code is in `compiler/src/`.  
To run linter checks:

```bash
cd compiler
go vet ./...
```

> More build/run instructions for the Go compiler will be added as development progresses.

---

## 📦 Standard Library

| Function       | Description                      |
|----------------|----------------------------------|
| `print(msg)`   | Print string to console          |
| `println(msg)` | Print string with newline        |
| `add(a, b)`    | Integer addition                 |
| `mul(a, b)`    | Integer multiplication           |
| `div(a, b)`    | Division (integer/double)        |
| `sub(a, b)`    | Subtraction                      |
| `mod(a, b)`    | Modulo operation                 |
| `powi(b, e)`   | Power of integers                |
| `sqrt(x)`      | Square root                      |
| `clock_now()`  | Time in seconds (double)         |

---

## 🎨 Advanced Features

### Virtual DOM
Build declarative UI components with efficient diffing:

```lumen
struct VNode {
    tag: string
    props: Map<string, string>
    children: Array<VNode>
    text: string
}

func create_element(tag: string, props: Map<string, string>, children: Array<VNode>) -> VNode {
    return VNode{tag: tag, props: props, children: children, text: ""}
}
```

### Ray Tracer
3D graphics with cross-platform compilation:

```lumen
struct Vector3 {
    x: double
    y: double
    z: double
}

struct Sphere {
    center: Vector3
    radius: double
}

func intersect(ray: Ray, sphere: Sphere) -> double {
    // Ray-sphere intersection algorithm
}
```

### HTTP Server
High-performance event-driven server:

```bash
cd examples/http_server
llc server.ll -filetype=obj -o server.o
clang server.o -o http_server
./http_server
```

---

## ✨ Example Programs

### Hello World
```lumen
func main() -> int {
    println("Hello, Lumen!")
    return 0
}
```

### Fibonacci
```lumen
func fib(n: int) -> int {
    if n <= 1 {
        return n
    }
    return add(fib(sub(n, 1)), fib(sub(n, 2)))
}

func main() -> int {
    println("Fibonacci(10) = " + fib(10).to_string())
    return 0
}
```

### Virtual DOM Component
```lumen
struct Counter {
    count: int
}

func render_counter(counter: Counter) -> VNode {
    let button_props = Map<string, string>{"onclick": "increment()"}
    let count_text = create_text_node("Count: " + counter.count.to_string())
    return create_element("div", Map<string, string>{}, Array<VNode>{count_text})
}
```

---

## 🌐 Web Development

### Playground
- Interactive code editor with syntax highlighting
- Real-time compilation simulation
- Built-in examples and tutorials
- Error handling and output display

### WebAssembly Integration
- Compile Lumen to WASM for browser execution
- JavaScript bindings for DOM manipulation
- Canvas-based graphics and animations
- Cross-browser compatibility

---

## 💡 Roadmap

- [x] LLVM-based native backend
- [x] WASM backend support
- [x] Minimal standard library
- [x] Custom parser + compiler in Go
- [x] Virtual DOM and frontend DSL
- [x] Language playground in browser
- [x] High-performance HTTP server
- [x] 3D graphics (ray tracer)
- [ ] Module system
- [ ] Generics support
- [ ] Async/await
- [ ] Package manager
- [ ] IDE integration

---

## 🤝 Contributing

We welcome contributions! See `examples/README.md` for guidelines on adding new examples.

### Development Areas
- **Compiler**: Go-based lexer, parser, and code generation
- **Standard Library**: Cross-platform utilities and functions
- **Examples**: Demonstrations of language features
- **Documentation**: Tutorials and API references

---

## 📄 License

MIT License
