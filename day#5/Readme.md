# Day 5: Testing in Go

## 1. Introduction to Testing in Go
Go has a built-in testing framework that makes it easy to write and run tests for your code. Here are the key things to know:

- Test files should have the naming convention `<filename>_test.go`
- The `go.mod` file is used to manage dependencies for your tests
- Tests are written using the `testing` package, which provides the `testing.T` struct

## 2. Deciding What to Test
When writing tests, it's important to decide what to test. Generally, you should focus on:
- Testing the functionality of your core business logic
- Ensuring edge cases and error handling work as expected
- Verifying the behavior of your public APIs

## 3. Naming Test Functions
Test function names should follow the convention `TestXxx(t *testing.T)`, where `Xxx` is a descriptive name for the test case.

```go
func TestAddNumbers(t *testing.T) {
    // Test the AddNumbers function
}
```

## 4. Assertion and Error Handling
Go's testing framework provides various assertion functions, such as `t.Errorf()` and `t.Fatal()`, to help you write clear and informative tests.

```go
func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Errorf("Divide(10, 0) should have returned an error")
    }
}
```

## 5. Running Tests
To run your tests, use the `go test` command in your terminal:

```
$ go test ./...
```

This will run all the tests in your project.

## Key Takeaways
- Test files use the naming convention `<filename>_test.go`
- The `go.mod` file manages dependencies for your tests
- Focus on testing core functionality, edge cases, and public APIs
- Use descriptive names for test functions
- Leverage `t.Errorf()` and `t.Fatal()` for clear error reporting

## Practice Exercises
1. Write tests for a simple function that calculates the area of a rectangle.
2. Implement tests that cover both valid and invalid inputs for a function.
3. Create a test suite for a more complex module in your Golang project.

## Resources
- [Go Documentation - Testing](https://golang.org/pkg/testing/)
- [Go by Example - Testing](https://gobyexample.com/testing)
- [Effective Go - Testing](https://golang.org/doc/effective_go.html#testing)