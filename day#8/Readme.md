# Day 8 - Maps and Interfaces in Go

## Maps in Go
Maps are Go's built-in hash table implementation that store key-value pairs. Unlike structs, maps can grow dynamically and are ideal when the number of key-value pairs is unknown at compile time.

### Defining Maps
There are multiple ways to define a map:

```go
// Method 1: Using make
studentScores := make(map[string]int)

// Method 2: Map literal
studentScores := map[string]int{
    "Alice": 92,
    "Bob":   85,
    "Charlie": 88,
}

// Method 3: Empty map literal
studentScores := map[string]int{}

// Declaring a nil map (not initialized)
var studentScores map[string]int
```

### Basic Operations

#### Insertion and Update
```go
// Insertion
studentScores["David"] = 95

// Update existing value
studentScores["Alice"] = 94
```

#### Deletion
```go
// Delete an entry
delete(studentScores, "Bob")
```

#### Accessing Values
```go
// Simple access
score := studentScores["Alice"]

// Safe access with existence check
score, exists := studentScores["Alice"]
if exists {
    fmt.Printf("Alice's score is: %d\n", score)
} else {
    fmt.Println("Alice not found")
}
```

### Maps vs Structs
```go
// Struct example - Fixed fields known at compile time
type Student struct {
    Name  string
    Score int
    Grade string
}

// Map example - Dynamic key-value pairs
studentDetails := map[string]interface{}{
    "name":  "Alice",
    "score": 92,
    // Can add more fields later
}
```

Key differences:
- Structs: Fixed fields, compile-time type checking, better performance
- Maps: Dynamic key-value pairs, runtime flexibility, slightly higher overhead

## Interfaces in Go

Interfaces define a set of method signatures. They provide abstraction and help in writing more flexible and reusable code.

### Basic Interface Example
```go
// Define an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implement interface for Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Implement interface for Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Function that works with any Shape
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %0.2f\n", s.Area())
    fmt.Printf("Perimeter: %0.2f\n", s.Perimeter())
}
```

### Why Interfaces?
1. Code Reusability: Write functions that work with any type implementing the interface
2. Loose Coupling: Dependencies on interfaces rather than concrete implementations
3. Easy Testing: Mock implementations for testing
4. Flexibility: Add new types without changing existing code

## Important Notes
- Go will complain if you declare variables but don't use them (compile-time error)
- Interface implementation is implicit in Go (no explicit declaration needed)
- Use interfaces to define behavior rather than data storage