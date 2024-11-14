
# Go Web Server

A simple web server built with Go. This project is organized into modular packages for easy management.

## Project Structure

```
/go-web-server
    └── pkg/
        ├── articles.go    # Handles article operations
        ├── consts.go      # Contains constants
        ├── renderers.go   # Renders responses
        ├── router.go      # Defines HTTP routes
        └── server.go      # Sets up and runs the server
```

## Getting Started

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/go-web-server.git
   ```

2. **Run the server**:
   ```bash
   go run ./pkg/server.go
   ```

3. **Access the server**:
   Visit `http://localhost:8080` in your browser.

## License

This project is under the MIT License.
