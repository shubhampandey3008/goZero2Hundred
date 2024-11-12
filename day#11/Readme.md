# Go Learning - Day 11

## Topics Covered
- Working with Interfaces
- Reader and Writer Interfaces in Go
- Using io.Copy() for efficient data transfer
- Command Line Arguments
- Introduction to Goroutines

## Interfaces in Go
Interfaces define a contract of behavior through methods. They enable loose coupling and make code more flexible and testable.

### Example: Basic Interface Implementation
```go
// Define an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implement interface for Rectangle
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}

// Usage
func main() {
    var s Shape = Rectangle{width: 5, height: 3}
    fmt.Printf("Area: %.2f\n", s.Area())
}
```

### Pros and Cons of Interfaces

Pros:
- Enables polymorphic behavior
- Facilitates unit testing through mock implementations
- Promotes loose coupling between components
- Makes code more modular and maintainable

Cons:
- Can add complexity if overused
- May lead to interface pollution if not designed carefully
- Potential performance overhead (minimal in most cases)

## Reader and Writer Interfaces
Go's `io.Reader` and `io.Writer` interfaces are fundamental to I/O operations in Go.

```go
// Reader interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer interface
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Example: Reading from a file
func readFile() error {
    file, err := os.Open("example.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    buffer := make([]byte, 1024)
    _, err = file.Read(buffer)
    return err
}
```

## Using io.Copy()
`io.Copy()` efficiently copies data from a Reader to a Writer.

```go
func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, sourceFile)
    return err
}
```

## Command Line Arguments
Go provides easy access to command line arguments through the `os.Args` slice.

```go
func main() {
    // os.Args[0] is the program name
    // os.Args[1:] contains the arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: program [arguments]")
        return
    }

    for i, arg := range os.Args[1:] {
        fmt.Printf("Argument %d: %s\n", i+1, arg)
    }
}
```

## Introduction to Goroutines
Goroutines are lightweight threads managed by the Go runtime. They enable concurrent execution.

```go
func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func printLetters() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main() {
    // Run concurrently using goroutines
    go printNumbers()
    go printLetters()

    // Wait to see the results
    time.Sleep(2 * time.Second)
}
```

### Key Points About Goroutines:
- Extremely lightweight (a few KB in stack size)
- Managed by Go runtime scheduler
- Communication through channels (not covered today)
- Much more efficient than OS threads

## Next Steps
- Explore channel communication between goroutines
- Learn about mutexes and other synchronization primitives
- Practice more complex interface implementations
- Study error handling patterns with interfaces