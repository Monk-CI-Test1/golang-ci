# Go Basics 3

A structured reference guide to core Go concepts with examples.

## Table of Contents

- [1. Why Go?](#1-why-go)
- [2. Pointers](#2-pointers)
  - [Pointer Example](#pointer-example)
  - [Pointer Operators](#pointer-operators)
  - [Pass by Value vs Pointer](#pass-by-value-vs-pointer)
- [3. Structs](#3-structs)
  - [Basic Struct Example](#basic-struct-example)
  - [Struct Embedding (Composition)](#struct-embedding-composition)
  - [Methods on Structs](#methods-on-structs)
  - [Mutable Methods on Structs](#mutable-methods-on-structs)
  - [Constructor Function for Structs](#contructor-function-for-structs)
  - [Struct Embedding vs Inheritance](#struct-embedding-vs-inheritance)
- [4. Interfaces](#4-interfaces)
  - [Empty Interface (any)](#empty-interface-any)
  - [Type Assertions and Type Switches](#type-assertions-and-type-switches)
  - [Interface Embedding](#interface-embedding)
  - [Interface Limitations and Dynamic Types](#interface-limitations-and-dynamic-types)
  - [Generics Concept](#generics-concept)
- [5. Sync and Async in Go](#5-sync-and-async-in-go)
  - [Synchronous Execution](#synchronous-execution)
  - [Asynchronous Execution with Goroutines](#asynchronous-execution-with-goroutines)
  - [Coordination with sync.WaitGroup](#coordination-with-syncwaitgroup)
- [6. Goroutines and Channels](#6-goroutines-and-channels)
- [7. Defer, Panic, and Recover](#7-defer-panic-and-recover)
- [8. Java Thread vs. Go Goroutine](#8-java-thread-vs-go-goroutine)
- [9. sync.Mutex (Mutual Exclusion)](#9-syncmutex-mutual-exclusion)

---

## 1. Why Go?
Go is a statically typed, compiled language designed for simplicity, performance, and concurrency. It’s great for building scalable web servers, microservices, and command-line tools.
- Fast compilation: Go compiles quickly, making development faster.
- Concurrency: Go has built-in support for concurrent programming with goroutines and channels. Helps in handling multiple tasks simultaneously.
- Static Typing: Go’s static type system helps catch errors at compile time, improving code reliability.
- Simplicity: Go has a clean syntax and a small standard library, making it easy to learn and use.
- Deployment: Go compiles to a single binary, simplifying deployment without worrying about dependencies.
- Concurrency in Go is Cheap and Easy.

## 2. Pointers

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

## 3. Structs

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


## 4. Interfaces
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

## 5. Sync and Async in Go
Go handles "sync" (sequential, blocking execution) and "async" (concurrent via goroutines/channels).
Goroutines provide lightweight concurrency while `sync` package coordinates them safely.

### Synchronous Execution
Operations run one-by-one; each blocks until complete. Simple but slow for I/O or parallelizable work.

### Asynchronous Execution with Goroutines
Goroutines are functions that run concurrently. They are lightweight and managed by the Go runtime. You can start a goroutine with the `go` keyword. They run concurrently without blocking main.

```go
go func() {
    fmt.Println("This runs in a goroutine")
}()
```

- Goroutines overlap execution; main needs time.Sleep or `sync` primitives to wait.

### Coordination with sync.WaitGroup
- Main thread does not wait for goroutines to finish by default. Use `sync.WaitGroup` to coordinate.
WaitGroup tracks completion; Add() increments counter, Done() decrements, Wait() blocks until zero.
Add(4): set the no. of goroutines to wait for.
Wait(): blocks until all goroutines call Done(). OR until waitgroup counter is zero.
Done(): decrements the waitgroup counter by 1, called by goroutine when it finishes its work.

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // signal completion
    fmt.Printf("Worker %d starting\n", id)
    // Simulate work
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    numWorkers := 3

    wg.Add(numWorkers) // set the number of goroutines to wait for

    for i := 1; i <= numWorkers; i++ {
        go worker(i, &wg) // start worker goroutine
    }

    wg.Wait() // wait for all workers to finish
    fmt.Println("All workers completed")
}
```

## 6. Goroutines and Channels
Goroutines are lightweight threads managed by the Go runtime. They allow you to run functions concurrently without blocking the main thread. 
Channels are used to communicate between goroutines safely, allowing you to send and receive data without explicit locks.

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, ch chan string) {
    time.Sleep(time.Second) // Simulate work
    ch <- fmt.Sprintf("Worker %d done", id) // Send result to channel
}

func main() {
    ch := make(chan string) // Create a channel

    for i := 1; i <= 3; i++ {
        go worker(i, ch) // Start worker goroutines
    }

    for i := 1; i <= 3; i++ {
        msg := <-ch // Receive messages from channel
        fmt.Println(msg)
    }
}
```

## 7. Defer, Panic, and Recover
- `defer`: schedules a function to run after the current function completes, useful for cleanup. E.g., closing files, releasing locks, etc.

- `panic`: stops normal execution and begins panicking, used for unrecoverable errors. E.g., when a critical error occurs that cannot be handled gracefully.
- `recover`: allows you to regain control of a panicking goroutine, used to handle panics gracefully. E.g., to prevent a program from crashing and to log the error instead.

```go
package main

import "fmt"
func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Something went wrong!")
}
func main() {
    riskyFunction()
    fmt.Println("Program continues after recovery")
}
```

## 8. Java Thread vs. Go Goroutine

| Feature    | Java Thread (OS Thread)              | Go Goroutine (Green Thread)          |
|------------|--------------------------------------|--------------------------------------|
| Memory     | Starts at ~1MB stack size.           | Starts at ~2KB stack size.           |
| Creation   | Expensive (System call to OS).       | Very cheap (Managed by Go runtime).  |
| Limit      | A few thousand per CPU.              | Millions per CPU.                    |
| Switching  | Managed by OS (Context switch is slow). | Managed by Go Scheduler (Fast).   |


## 9. sync.Mutex (Mutual Exclusion)
Used to prevent Race Conditions (When multiple goroutines access the same resource concurrently).
A `sync.Mutex` is a mutual exclusion lock that protects shared resources from concurrent access. It ensures that only one goroutine can access a critical section of code at a time.

- NOTE: ISSUE IN THE EXAMPLE BELOW: When multiple goroutines increment the `views` field of the `post` struct concurrently, it can lead to a race condition, resulting in an incorrect final count of views. This is because the increment operation (`p.views += 1`) is not atomic, and multiple goroutines can read and write to `views` simultaneously, causing some increments to be lost.
```go
package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views += 1
}

// Request will come concurrently

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.incrementViews(&wg)
	}
	wg.Wait()                 // Wait for all increments to finish
	fmt.Println(myPost.views) // Output: random
}
```

- FIXED CODE
```go
package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

// func (p *post) incrementViews(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	p.mu.Lock()
// 	p.views += 1
// 	defer p.mu.Unlock()
// }
func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.incrementViews(&wg)
	}
	wg.Wait()                 // Wait for all increments to finish
	fmt.Println(myPost.views) // Output: 1000
}
```

## 10. sync.RWMutex (Read-Write Mutex)
A `sync.RWMutex` allows multiple readers or one writer at a time. It provides better performance for read-heavy workloads by allowing concurrent reads while still ensuring exclusive access for writes.

```go
package main

import (
    "fmt"
    "sync"
)

type post struct {
    views int
    mu    sync.RWMutex
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
    defer func() {
        wg.Done()
        p.mu.Unlock()
    }()
    p.mu.Lock()
    p.views += 1
}

func (p *post) getViews() int {
    p.mu.RLock()
    defer p.mu.RUnlock()
    return p.views
}

func main() {
    var wg sync.WaitGroup
    myPost := post{views: 0}

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go myPost.incrementViews(&wg)
    }
    wg.Wait() // Wait for all increments to finish

    fmt.Println(myPost.getViews()) // Output: 1000
}
```
