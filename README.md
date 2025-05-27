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
├── compiler/
│   └── src/
│       ├── lexer.go        # Go lexer for Lumen
│       ├── parser.go       # Go parser for Lumen
│       ├── ast.go          # AST node definitions
│       └── tokens.go       # Token types and struct
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
| `powi(b, e)`   | Power of integers                |
| `clock_now()`  | Time in seconds (double)         |

---

## ✨ Example LLVM IR

```llvm
declare i32 @add(i32, i32)
declare void @println(i8*)

@str_result = constant [8 x i8] c"Result:\0A\00"

define i32 @main() {
entry:
  %sum = call i32 @add(i32 3, i32 4)
  %msg_ptr = getelementptr [8 x i8], [8 x i8]* @str_result, i32 0, i32 0
  call void @println(i8* %msg_ptr)
  ret i32 0
}
```

---

## 💡 Roadmap

- [x] LLVM-based native backend
- [x] WASM backend support
- [x] Minimal standard library
- [x] Custom parser + compiler in Go
- [ ] Virtual DOM and frontend DSL
- [ ] Language playground in browser

---

## 📄 License

MIT License
