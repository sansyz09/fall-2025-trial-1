# 

# Go Programming Homework Instructions

## 📋 Overview

This homework consists of 10 Go programming exercises designed to teach fundamental programming concepts. Each exercise includes starter code and automated tests to verify your solutions.

---

## 🔧 Step 1: Install Go

### Windows

1. **Download Go**:
    - Visit https://golang.org/dl/
    - Download the Windows installer (`.msi` file) for 64-bit systems
    - Example: `go1.22.0.windows-amd64.msi`
2. **Install Go**:
    - Double-click the downloaded `.msi` file
    - Follow the installation wizard (use default settings)
    - Go will be installed to `C:\Program Files\Go`
3. **Verify Installation**:
    - Press `Win + R`, type `cmd`, press Enter
    - Type: `go version`
    - Expected output: `go version go1.22.0 windows/amd64`

### macOS

1. **Using the Official Installer**:
    - Visit https://golang.org/dl/
    - Download the macOS installer (`.pkg` file)
    - Double-click and follow the installation wizard
2. **Using Homebrew** (if installed):
    
    ```bash
    brew install go
    
    ```
    
3. **Verify Installation**:
    - Press `Cmd + Space`, type "Terminal", press Enter
    - Type: `go version`
    - Expected output: `go version go1.22.0 darwin/amd64`

### Linux (Ubuntu/Debian)

1. **Using Package Manager** (easiest):
    
    ```bash
    sudo apt update
    sudo apt install golang-go
    
    ```
    
2. **Using Official Binary** (latest version):
    
    ```bash
    # Download latest version (check golang.org for current version)
    wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
    
    # Remove any previous installation and extract
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
    
    # Add Go to your PATH
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    source ~/.bashrc
    
    ```
    
3. **Verify Installation**:
    
    ```bash
    go version
    
    ```
    

---

---

## 🧪 How to Run Tests

### Testing Individual Exercises

```bash
# Test only exercise 1
go test exercise1.go exercise1_test.go

# Test only exercise 2
go test exercise2.go exercise2_test.go

```

### Testing All Exercises

```bash
# Test all exercises in the directory
go test

# Or test all with subdirectories
go test ./...

```

### Verbose Testing (Recommended)

```bash
# See detailed output for all tests
go test -v

# Verbose output for specific exercise
go test -v exercise1.go exercise1_test.go

```

---

## 📊 Understanding Test Results

### ✅ **Passing Test**:

```
=== RUN   TestGreet
--- PASS: TestGreet (0.00s)
PASS
ok      go-homework    0.002s

```

**Meaning**: Your solution is correct!

### ❌ **Failing Test**:

```
=== RUN   TestGreet
--- FAIL: TestGreet (0.00s)
    exercise1_test.go:8: Expected Hello, Alice!, got
FAIL
exit status 1
FAIL    go-homework    0.002s

```

**Meaning**: Your solution needs work. The error message tells you what was expected vs. what you returned.

### 🔍 **Reading Error Messages**:

- `Expected Hello, Alice!, got`  - Your function returned an empty string instead of "Hello, Alice!"
- Line numbers (e.g., `exercise1_test.go:8`) show exactly where the test failed

---

## 🎯 Step 5: Your Homework Tasks

Complete these 10 exercises in order:

1. **Hello World Function** - Return a greeting message
2. **Sum Two Numbers** - Add two integers
3. **Even or Odd** - Check if a number is even
4. **Find Maximum** - Return the larger of two numbers
5. **String Length** - Count characters in a string
6. **Factorial** - Calculate n! (n factorial)
7. **Array Sum** - Sum all numbers in a slice

---

## 🛠️ Essential Commands Reference

| Command | Purpose |
| --- | --- |
| `go version` | Check if Go is installed |
| `go mod init project-name` | Start a new Go project |
| `go test` | Run all tests |
| `go test -v` | Run tests with detailed output |
| `go test file.go file_test.go` | Test specific files |
| `go run main.go` | Run a Go program |

---

## 🚨 Troubleshooting

### Problem: "go: command not found"

**Solution**: Go is not installed or not in your PATH

- **Windows**: Reinstall using the `.msi` installer
- **macOS/Linux**: Add to PATH or reinstall

### Problem: "no Go files in directory"

**Solution**:

- Make sure you're in the correct directory (`go-homework`)
- Ensure you've created the `.go` files

### Problem: "package main is not a main package"

**Solution**:

- Both `.go` and `_test.go` files must start with `package main`
- Check for typos in the package declaration

### Problem: Tests don't run

**Solution**:

- Test files must end with `_test.go`
- Test functions must start with `Test` (capital T)
- Test functions must accept `testing.T` parameter

---

## 💡 Tips for Success

1. **Start Simple**: Begin with Exercise 1 and work sequentially
2. **Read Error Messages**: They tell you exactly what's wrong
3. **Test Frequently**: Run tests after each change
4. **Use `fmt.Println()`** for debugging if needed
5. **Ask for Help**: If stuck, don't hesitate to ask questions
6. **Practice**: The more you code, the easier it becomes

---

## 📚 Additional Resources

- **Go Documentation**: https://golang.org/doc/
- **Go Tour** (Interactive Tutorial): https://tour.golang.org/
- **Go by Example**: https://gobyexample.com/

Good luck with your homework! 🚀