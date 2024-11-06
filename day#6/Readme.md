# Day 6: Structs in Go

## 1. Defining Structs
In Go, you can create custom data types using `struct`. A `struct` is a collection of fields, similar to a class in object-oriented programming.

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```

## 2. Zero Values
When you create a new `struct` instance, the fields are initialized with their respective zero values. For example:

```go
var pandeyJi Person
fmt.Printf("%+v\n", pandeyJi) // Output: {FirstName: LastName: Age:0}
```

The zero values for the fields are:
- `string`: empty string `""`
- `int`, `float`: `0`
- `bool`: `false`

## 3. Accessing Struct Fields
You can access the fields of a `struct` using the dot (`.`) notation.

```go
pandeyJi.FirstName = "Hari"
pandeyJi.LastName = "Pandey"
pandeyJi.Age = 35
fmt.Printf("%+v\n", pandeyJi) // Output: {FirstName:Hari LastName:Pandey Age:35}
```

## 4. Structs with Receiver Functions
You can define functions that have a `struct` as a receiver. These functions are called "methods" and can provide behavior for your custom data types.

```go
func (p *Person) Introduce() {
    fmt.Printf("Hello, my name is %s %s and I'm %d years old.\n", p.FirstName, p.LastName, p.Age)
}

func main() {
    pandeyJi := Person{FirstName: "Hari", LastName: "Pandey", Age: 35}
    pandeyJi.Introduce() // Output: Hello, my name is Hari Pandey and I'm 35 years old.
}
```

In the above example, `Introduce()` is a method that has a `*Person` receiver. This allows the method to access and manipulate the fields of the `Person` struct.

## Key Takeaways
- Structs are custom data types that group related fields together
- Zero values are assigned to uninitialized struct fields
- You can access struct fields using the dot (`.`) notation
- Receiver functions (methods) can be defined on structs to add behavior

## Practice Exercises
1. Create a `Rectangle` struct with `Width` and `Height` fields, and a `Area()` method to calculate the area.
2. Implement a `Circle` struct with a `Radius` field, and methods to calculate the `Area()` and `Circumference()`.
3. Write a program that creates multiple `Person` instances and calls their `Introduce()` method.

## Resources
- [Go Documentation - Structs](https://golang.org/doc/effective_go.html#struct_literals)
- [Go by Example - Structs](https://gobyexample.com/structs)
- [Effective Go - Embedding and Promotion](https://golang.org/doc/effective_go.html#embedding)