# Lumen Examples

This directory contains various examples demonstrating Lumen's capabilities.

## Basic Examples

- `hello.lumen` - Simple "Hello, World!" program
- `arithmetic_example.lumen` - Basic arithmetic operations
- `factorial.lumen` - Recursive factorial calculation
- `fib.lumen` - Fibonacci sequence
- `gcd.lumen` - Greatest Common Divisor
- `pi_approx.lumen` - Pi approximation using Leibniz series

## Advanced Examples

### Ray Tracer (`raytracer/`)
A complete ray tracer implementation demonstrating:
- Complex data structures (Vector3, Sphere, Ray)
- Mathematical algorithms
- Cross-platform compilation (native + WebAssembly)
- Web integration with JavaScript

**Files:**
- `raytracer.lumen` - Lumen source code version
- `raytracer.ll` - LLVM IR version
- `raytracer.html` - Web interface
- `raytracer.js` - JavaScript integration

**Build:**
```bash
# Native compilation
lumen raytracer.lumen -o raytracer
./raytracer

# WebAssembly
lumen raytracer.lumen -target wasm -o raytracer.wasm
# Serve with HTTP server and open raytracer.html
```

### Virtual DOM (`virtual_dom/`)
A Virtual DOM implementation for frontend development:
- Declarative UI components
- Efficient diffing algorithm
- HTML generation
- Component system

**Files:**
- `vdom.lumen` - Virtual DOM implementation

**Features:**
- VNode structure for representing DOM elements
- Diff algorithm for efficient updates
- Component-based architecture
- String rendering for HTML output

### Language Playground (`playground/`)
An interactive web-based playground for writing and running Lumen code:
- Syntax-highlighted code editor
- Real-time compilation simulation
- Built-in examples
- Modern web interface

**Files:**
- `playground.html` - Main playground interface
- `playground.js` - Playground functionality

**Features:**
- CodeMirror-based editor with syntax highlighting
- Multiple example programs
- Simulated compilation and execution
- Error handling and output display
- Keyboard shortcuts (Ctrl+Enter to run)

**Usage:**
1. Open `playground.html` in a web browser
2. Write Lumen code in the editor
3. Click "Run" or press Ctrl+Enter to execute
4. View output in the right panel
5. Try different examples from the dropdown

## HTTP Server Example

See `http_server/` directory for a high-performance HTTP server implementation.

## Building and Running

### Prerequisites
- LLVM toolchain
- Clang compiler
- Emscripten (for WebAssembly)

### Native Compilation
```bash
lumen example.lumen -o example
./example
```

### WebAssembly Compilation
```bash
lumen example.lumen -target wasm -o example.wasm
```

### Using the HTTP Server
```bash
cd http_server
llc server.ll -filetype=obj -o server.o
clang server.o -o http_server
./http_server
```

## Contributing

When adding new examples:
1. Include both `.lumen` and `.ll` versions if applicable
2. Add comprehensive documentation
3. Include build instructions
4. Test both native and WebAssembly targets
5. Update this README with new examples

## Future Examples

Planned examples include:
- [ ] WebSocket client/server
- [ ] Database integration
- [ ] Graphics and animation
- [ ] Machine learning algorithms
- [ ] Network protocols
- [ ] Game development