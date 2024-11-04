# Day 4: File I/O, Type Conversion, and Random Number Generation

## 1. Writing to a File
Go provides a simple way to write data to files using the `os` package. Here's an example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Open a file for writing
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Write data to the file
    _, err = file.WriteString("Hello, Gophers!")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Data written to file successfully.")
}
```

## 2. Type Conversion in Go
Go is a statically-typed language, so you often need to convert between types. For example, converting a `[]string` to a `[]interface{}`:

```go
package main

import "fmt"

func main() {
    deck := []string{"Ace", "Two", "Three"}
    deckInterface := []interface{}(deck)
    fmt.Println(deckInterface) // Output: [Ace Two Three]
}
```

## 3. Deprecated `ioutil` Package
The `ioutil` package has been deprecated in Go 1.16, and you should use the `os` package instead for file operations. Here's an example of reading a file using `os`:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Read the file contents
    data := make([]byte, 100)
    n, err := file.Read(data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File contents:", string(data[:n]))
}
```

## 4. Exiting with Error Codes
You can use `os.Exit()` to exit a program with a specific error code. A non-zero exit code indicates an error, while 0 means a successful exit.

```go
package main

import "os"

func main() {
    // Exit with error code 1
    os.Exit(1)
}
```

## 5. One-line Swap and Random Number Generation
Go makes it easy to swap values and generate random numbers:

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // One-line swap
    a, b := 10, 20
    a, b = b, a
    fmt.Println("a:", a, "b:", b) // Output: a: 20 b: 10

    // Generate a random number
    randomNumber := rand.Intn(100)
    fmt.Println("Random number:", randomNumber)

    // Generating truly random numbers
    rand.Seed(int64(time.Now().Nanosecond()))
    trueRandomNumber := rand.Intn(100)
    fmt.Println("True random number:", trueRandomNumber)
}
```

## Key Takeaways
- Go's `os` package provides a simple API for file I/O operations
- Type conversion is essential in a statically-typed language like Go
- The `ioutil` package is deprecated, use `os` instead
- `os.Exit()` allows you to exit a program with a specific error code
- One-line swap and `rand.Intn()` are handy for common programming tasks
- Seeding the random number generator with `time.Now().Nanosecond()` ensures true randomness

## Practice Exercises
1. Write a program that creates a file, writes some data to it, and then reads the data back.
2. Implement a function that takes a slice of strings and returns a slice of `interface{}`.
3. Write a program that exits with a non-zero error code when a specific condition is met.
4. Create a program that swaps two variables and generates a random number.

## Resources
- [Go Documentation - os package](https://golang.org/pkg/os/)
- [Go by Example - Files](https://gobyexample.com/reading-files)
- [Go Documentation - Type Conversion](https://golang.org/ref/spec#Conversions)
- [Go Documentation - Math/rand package](https://golang.org/pkg/math/rand/)