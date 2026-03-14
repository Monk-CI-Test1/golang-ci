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

