# Day 7: Understanding Pointers in Go

## 1. Pass-by-Value in Go
Go is a pass-by-value language, which means when you pass a variable to a function, you're actually passing a copy of the value, not the original variable.

```go
func main() {
    x := 5
    changeValue(x)
    fmt.Println(x) // Output: 5 (unchanged)
}

func changeValue(y int) {
    y = 10
}
```

## 2. Pointer Operations
To work around the pass-by-value behavior, you can use pointers. Pointers hold the memory address of a variable, allowing you to modify the original value.

![Pointer Operations](image.png)

```go
func main() {
    x := 5
    changeValueWithPointer(&x)
    fmt.Println(x) // Output: 10 (updated)
}

func changeValueWithPointer(y *int) {
    *y = 10
}
```

In the above example, `&x` gives us the memory address of `x`, and `*y` dereferences the pointer to access the value it points to.

## 3. Pointer Shortcuts
Go provides some shortcuts for working with pointers:

![Pointer Shortcuts](image-1.png)

```go
// Creating a pointer using the "new" keyword
p := new(int)
*p = 42

// Pointer to a struct
type Person struct {
    Name string
}
john := &Person{Name: "John"}
```

## 4. Slices, Maps, Channels, and Functions
Slices, maps, channels, and functions in Go are all passed by reference, so you don't need to use the `&` operator when passing them to functions.

![Pass by Reference](image-2.png)

```go
func main() {
    numbers := []int{1, 2, 3}
    modifySlice(numbers)
    fmt.Println(numbers) // Output: [10 2 3]
}

func modifySlice(nums []int) {
    nums[0] = 10
}
```

## Key Takeaways
- Go is a pass-by-value language, but you can use pointers to modify values
- Pointers hold the memory address of a variable, accessed with the `&` operator
- Dereferencing a pointer with the `*` operator allows you to access the value it points to
- Slices, maps, channels, and functions are passed by reference in Go

## Practice Exercises
1. Write a function that swaps the values of two integers using pointers.
2. Implement a function that doubles the values in a slice of integers.
3. Create a struct with pointer fields and demonstrate accessing and modifying the values.

## Resources
- [Go Documentation - Pointers](https://golang.org/doc/effective_go.html#pointers_and_structs)
- [Go by Example - Pointers](https://gobyexample.com/pointers)
- [Effective Go - Initialization](https://golang.org/doc/effective_go.html#initialization)