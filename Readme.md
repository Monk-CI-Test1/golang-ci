# Go Basics: Pointers and Structs

This document explains two core Go concepts with examples:

1. Pointers
2. Structs

## 1. Pointers

By default, Go uses pass-by-value. That means when you pass a variable to a function, Go creates a copy.

A pointer is a variable that stores the memory address of another variable. Pointers help you:

- avoid unnecessary copying
- update original data from functions
- work efficiently with large values

### Pointer Example

```go
package main

import "fmt"

func main() {
    x := 10
    p := &x // p points to x

    fmt.Printf("Value of x: %d\n", x)
    fmt.Printf("Address of x: %p\n", &x)
    fmt.Printf("Value of p (address): %p\n", p)
    fmt.Printf("Value at p: %d\n", *p)

    *p = 20 // update x through pointer

    fmt.Printf("New value of x: %d\n", x)
}
```

### Pointer Operators

- `&` (address-of): gets the memory address of a variable
- `*` (dereference): gets or updates the value at a memory address

### Pass by Value vs Pointer

```go
type User struct {
    Username string
}

func UpdateNameValue(u User) {
    u.Username = "NewName" // only changes local copy
}

func UpdateNamePointer(u *User) {
    u.Username = "NewName" // changes original value
}
```

## 2. Structs

Structs are typed collections of fields. They help model real-world entities by grouping related data.

Note about visibility:

- capitalized names (for example, `User`) are exported (public)
- lowercase names (for example, `user`) are unexported (package-private)

### Basic Struct Example

```go
package main

import "fmt"

type User struct {
    ID       int
    Username string
    Email    string
    IsActive bool
}

func main() {
    u := User{ID: 1, Username: "gopher123", Email: "go@example.com", IsActive: true}
    fmt.Println(u.Username)
}
```

### Struct Embedding (Composition)

Go does not use class **inheritance**. Instead, it uses **embedding** to promote fields and methods.

```go
package main

import "fmt"

type User struct {
    Username string
}

type Admin struct {
    User  // embedded struct
    Level int
}

func main() {
    a := Admin{
        User:  User{Username: "admin01"},
        Level: 10,
    }

    fmt.Println(a.Username) // promoted field from embedded User
}
```

### Methods on Structs

You can attach behavior to structs using methods.

```go
package main

import "fmt"

type User struct {
    Username string
    Email    string
    IsActive bool
}

func (u User) IsActiveUser() bool {
    return u.IsActive
}

func main() {
    u := User{Username: "gopher123", Email: "go@example.com", IsActive: true}
    fmt.Println(u.IsActiveUser())
}
```

#### Mutable Methods on Structs

```go
func (u *User) Deactivate() {
    u.IsActive = false
}

func main() {
    u := User{Username: "gopher123", Email: "go@example.com", IsActive: true}
    u.Deactivate()
    fmt.Println(u.IsActive) // false
}
```

#### Contructor Function for Structs

```go
func NewUser(username, email string) (*User, error) { // used pointer return type to avoid copying

    if username == "" || email == "" {
        return nil, fmt.Errorf("invalid user input")
    }

    return &User{
        Username: username,
        Email: email,
        IsActive: true
    }, nil
}

var appUser *User
appUser, err := NewUser("gopher123", "go@example.com")
if err != nil {
    fmt.Println("Error creating user:", err)
    return
}
fmt.Println(appUser.Username) // gopher123
````

#### Struct Embedding vs Inheritance
Go does not support traditional class inheritance. Instead, it uses struct embedding to achieve similar functionality. When you embed a struct, the fields and methods of the embedded struct are promoted to the outer struct, allowing you to access them directly.

```go
type User struct {
	Username string
	Email    string
}

func (u User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(newEmail string) {
	u.Email = newEmail
}

type Admin struct {
	User  // embedded struct
	Level int
}

func main() {
	a := Admin{
		User:  User{Username: "admin01", Email: "admin@example.com"},
		Level: 10,
	}

	fmt.Println(a.Username)   // promoted field from embedded User
	fmt.Println(a.Email)      // promoted field from embedded User
	fmt.Println(a.GetEmail()) // promoted method from embedded User
	a.SetEmail("newemail@example.com")
	fmt.Println(a.Email) // promoted field from embedded User
	fmt.Println(a.Level) // field from Admin struct
}
```


## 3. Interfaces
Interfaces define a set of method signatures. Any type that implements those methods satisfies the interface, allowing for polymorphism. It doesn’t contain data; it only defines behavior.

Benefits of interfaces:
- decoupling: code can work with any type that satisfies the interface
- flexibility: you can define multiple implementations of the same interface
- easier testing: you can create mock implementations for testing
- better code organization: interfaces help define clear contracts for behavior


```go
package main

import "fmt"

// 1. Define the interface (Contract)
type Shape interface {
    Area() float64
}
// Any struct that implements 'Shape' interface must have an Area() method that returns a float64 value.

// 2. Implement with a Struct (Circle)
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// 3. Implement with another Struct (Square)
type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

func main() {
    // Both Circle and Square "are" Shapes
    shapes := []Shape{
        Circle{Radius: 5},
        Square{Side: 10},
    }

    for _, s := range shapes {
        fmt.Printf("Area: %0.2f\n", s.Area())
    }
}
```

#### Empty Interface (any)
The empty interface `interface{}` can hold values of any type. It’s often used for functions that need to accept any type of data.

```go
func PrintValue(v interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)

func PrintValue2(v any) { // 'any' is an alias for 'interface{}'
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}
```

#### Type Assertions and Type Switches
You can use type assertions to extract the underlying value from an interface variable, or type switches to handle multiple types.

```go
func main() {
    var v interface{} = "Hello, Go!"
    // Type assertion
    str, ok := v.(string)
    if ok {
        fmt.Println("String value:", str)
    } else {
        fmt.Println("Not a string")
    }   


    // Type switch
    switch val := v.(type) {
    case string:
        fmt.Println("String value:", val)
    case int:
        fmt.Println("Integer value:", val)
    default:
        fmt.Println("Unknown type")
    }
}
```

#### Interface Embedding
Go allows you to embed interfaces within other interfaces, which promotes the methods of the embedded interface to the outer interface.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}
```

#### Interface Limitations and Dynamic Types
While interfaces provide powerful abstraction, they have some limitations:
- no support for fields (interfaces only define behavior, not data)
- no support for constructors (you need to create instances of concrete types that implement the interface)

```go
func add(a, b interface{}) {
    return a + b // This will cause a compile-time error because the compiler doesn't know how to add two empty interfaces
}

func add(a, b interface{}) interface{} {
    aInt, aIsInt := a.(int)
    bInt, bIsInt := b.(int)
    if aIsInt && bIsInt {
        return aInt + bInt
    }
    
    aFloat, aIsFloat := a.(float64)
    bFloat, bIsFloat := b.(float64)
    if aIsFloat && bIsFloat {
        return aFloat + bFloat
    }
    return nil
}
```

#### Generics Concept
This is a feature introduced in Go 1.18 that allows you to write code that can work with multiple types. Generics enable you to create functions and types that are parameterized by type, providing type safety at compile time.

```go
func add[T comparable](a, b T) T {
    return a + b
}

func add[T any](a, b T) T {
    return a + b
}

func add[T int | float64 | string](a, b T) T {
    return a + b
}
```