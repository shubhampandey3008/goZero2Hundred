# Golang Learning - Day 13

## Topics Covered
- Function Literals (Anonymous Functions)
- Goroutines and Concurrency
- Channels
- Practical Project: Website Status Checker

## Function Literals
Function literals in Go are anonymous functions that can be assigned to variables or passed as arguments to other functions.

```go
// Example of a function literal
func main() {
    greet := func(name string) string {
        return "Hello, " + name
    }
    
    fmt.Println(greet("Gopher")) // Output: Hello, Gopher
}
```

## Goroutines and Variable Scope
Important learning: Variables should not be accessed from different goroutines without proper synchronization.

```go
// Bad Practice
func main() {
    message := "Hello"
    go func() {
        // Accessing variable from different goroutine - potential race condition
        fmt.Println(message)
    }()
    message = "Goodbye"
}

// Good Practice
func main() {
    message := "Hello"
    go func(msg string) {
        // Passing variable as parameter ensures proper value capture
        fmt.Println(msg)
    }(message)
    message = "Goodbye"
}
```

## Goroutine Execution
When launching a goroutine, we need to ensure the main program doesn't exit before the goroutine completes.

```go
func main() {
    // Without wait - goroutine might not execute
    go printMessage("Hello")
    // Program ends immediately

    // With wait - ensures goroutine execution
    go printMessage("Hello")
    time.Sleep(time.Second)
    // Or better, use sync.WaitGroup
}

func printMessage(msg string) {
    fmt.Println(msg)
}
```

## Channels
Channels in Go work like teleportation doors - both sender and receiver must be ready for the communication to happen.

```go
func main() {
    ch := make(chan string)

    // This will deadlock if no receiver
    go func() {
        ch <- "Hello"  // Sender
    }()

    message := <-ch    // Receiver
    fmt.Println(message)

    // Deadlock example
    ch <- "Hello"  // Will deadlock if no receiver is ready
}
```

## Project: Website Status Checker
Created a concurrent website checker that monitors multiple websites simultaneously using goroutines and channels.

```go
type WebsiteStatus struct {
    URL    string
    Status bool
}

func main() {
    websites := []string{
        "http://google.com",
        "http://facebook.com",
        "http://golang.org",
    }

    statusChannel := make(chan WebsiteStatus)

    for _, site := range websites {
        go checkWebsite(site, statusChannel)
    }

    // Receive results
    for i := 0; i < len(websites); i++ {
        result := <-statusChannel
        fmt.Printf("Website: %s - Status: %v\n", result.URL, result.Status)
    }
}

func checkWebsite(url string, c chan WebsiteStatus) {
    _, err := http.Get(url)
    c <- WebsiteStatus{
        URL:    url,
        Status: err == nil,
    }
}
```

## Key Takeaways
1. Function literals provide a way to create anonymous functions
2. Always be careful with variable scope in goroutines
3. Ensure proper synchronization when using goroutines
4. Channels require both sender and receiver to be ready
5. Practical applications of goroutines and channels can create efficient concurrent programs

## Next Steps
- Explore more complex channel patterns
- Learn about buffered channels
- Study mutex and other synchronization primitives
- Practice creating more concurrent applications