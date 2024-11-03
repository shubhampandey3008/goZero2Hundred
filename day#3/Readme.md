# Day 3: Functions with Receivers, Slices, and Control Flow

## 1. Functions with Receivers (Methods)
In Go, a receiver is a special parameter that allows you to associate a function with a type, effectively creating a method. This is how Go implements behavior similar to object-oriented programming.

```go
// Defining a custom type
type Rectangle struct {
    width  float64
    height float64
}

// Method with a receiver
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// Method with a pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    fmt.Printf("Area: %.2f\n", rect.Area())
    rect.Scale(2)
    fmt.Printf("Area after scaling: %.2f\n", rect.Area())
}
```

## 2. Slice Splitting
Go provides powerful built-in features for manipulating slices, including splitting them into sub-slices.

```go
func main() {
    // Creating and splitting slices
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Slice syntax: slice[startIndex:endIndex]
    firstHalf := numbers[:5]    // [1, 2, 3, 4, 5]
    secondHalf := numbers[5:]   // [6, 7, 8, 9, 10]
    middle := numbers[2:7]      // [3, 4, 5, 6, 7]
    
    fmt.Println("First half:", firstHalf)
    fmt.Println("Second half:", secondHalf)
    fmt.Println("Middle:", middle)
}
```

## 3. Handling Multiple Return Values
Go functions can return multiple values, which is particularly useful for error handling and returning related values.

```go
// Function returning multiple values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    // Using multiple return values
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Result: %.2f\n", result)
    
    // Using blank identifier (_) to ignore error
    result, _ = divide(20, 4)
    fmt.Printf("Result: %.2f\n", result)
}
```

## 4. Nested For Loops
Nested loops are useful for working with multi-dimensional data or when you need to process data in a matrix-like fashion.

```go
func main() {
    // Example of nested loops - Multiplication table
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("%d x %d = %d\t", i, j, i*j)
        }
        fmt.Println() // New line after each row
    }
    
    // Using nested loops with slices
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
}
```

## Key Takeaways
- Receivers allow you to create methods associated with custom types
- Slice splitting provides flexible ways to work with portions of slices
- Multiple return values make error handling more idiomatic in Go
- Nested loops are powerful for working with multi-dimensional data structures

## Practice Exercises
1. Create a custom type with both value and pointer receivers
2. Write a program that splits a slice in different ways and observes the behavior
3. Implement a function that returns multiple values including error handling
4. Create a pattern using nested loops

## Resources
- [Go Documentation - Methods](https://golang.org/doc/effective_go.html#methods)
- [Go Documentation - Slices](https://blog.golang.org/slices-intro)
- [Go by Example - Multiple Return Values](https://gobyexample.com/multiple-return-values)