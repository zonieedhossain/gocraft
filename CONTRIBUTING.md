# Contributing to gocraft

Thank you for your interest in contributing to **gocraft**! We welcome contributions from the community to help make this tool better.

## 🛠️ Development Setup

1. **Fork the repository** on GitHub.
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/gocraft.git
   cd gocraft
   ```
3. **Install dependencies**:
   ```bash
   go mod download
   ```

## 🏗️ Building and Running

You can build the CLI and install it locally to test your changes.

```bash
# Build the binary
go build -o gocraft main.go

# Install globally (ensure $GOPATH/bin is in your PATH)
go install .
```

To run the tool directly from source:
```bash
go run main.go [command] [flags]
```

## 🧪 Testing

Please ensure all tests pass before submitting a Pull Request.

```bash
go test ./...
```

If you are adding a new feature, please include relevant tests.

## 📝 Code Style

- Follow standard Go conventions (run `go fmt ./...`).
- Keep code modular and clean.
- Add comments for exported functions and complex logic.

## 📬 Submitting a Pull Request

1. Create a new branch for your feature or fix: `git checkout -b feature/amazing-feature`.
2. Commit your changes with clear messages.
3. Push to your fork: `git push origin feature/amazing-feature`.
4. Open a Pull Request on the main repository describing your changes.

---

Happy Coding! 🚀
