# Lumen HTTP Server Example

This is a high-performance HTTP server implementation in Lumen, demonstrating the language's capabilities for system-level programming. The server uses epoll (Linux) or kqueue (macOS) for efficient I/O multiplexing.

## Features

- ⚡ High-performance event-driven architecture
- 🔄 Non-blocking I/O with epoll/kqueue
- 🚀 Single-threaded event loop
- 📦 Simple HTTP/1.1 response handling
- 🔒 Basic error handling and resource cleanup

## Building

### Linux

```bash
# Compile the LLVM IR to object file
llc server.ll -filetype=obj -o server.o

# Link with system libraries
clang server.o -o http_server

# Run the server
./http_server
```

### macOS

```bash
# Compile the LLVM IR to object file
llc server.ll -filetype=obj -o server.o

# Link with system libraries
clang server.o -o http_server

# Run the server
./http_server
```

## Testing

Once the server is running, you can test it using curl:

```bash
curl http://localhost:8080
```

You should receive a "Hello, World!" response.

## Implementation Details

The server implements a basic HTTP/1.1 server with the following components:

1. Socket creation and configuration
2. Event loop using epoll/kqueue
3. Non-blocking I/O operations
4. Basic HTTP response handling

The server listens on port 8080 by default and can handle multiple concurrent connections efficiently.

## Performance

This implementation is designed for high performance:
- Uses event-driven architecture
- Implements non-blocking I/O
- Efficient memory usage with fixed-size buffers
- Minimal system calls in the main event loop

## Future Improvements

- [ ] Add HTTP request parsing
- [ ] Implement proper HTTP response headers
- [ ] Add support for different HTTP methods
- [ ] Implement keep-alive connections
- [ ] Add basic routing capabilities 