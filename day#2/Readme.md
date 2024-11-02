# Go Imports , Variables , Functions , Arrays

## Import
* Import is used to import other prewritten packages
  * `"fmt"` means format, used to print (standard library)

## Variables
* Go is statically typed
* `:=` is used to infer the type very first time
* Other types includes `int`, `string`, `float64`, `bool` etc.

## Functions
* Need to declare the return type after the name of the function

## Arrays and Slices
* Arrays are fixed and slices can grow
* Use `append()` to add an element

### Example:
```go
// Arrays - when assigned, full copy is made
arr1 := [3]int{1, 2, 3}
arr2 := arr1        // Creates a complete copy
arr2[0] = 10        // Only changes arr2, not arr1

// Slices - when assigned, reference is copied
slice1 := []int{1, 2, 3}
slice2 := slice1    // Both reference same underlying array
slice2[0] = 10      // Changes reflect in both slice1 and slice2