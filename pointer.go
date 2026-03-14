// Pointer is simply a variable that stores the address of another variable. For which this gives you the power to manipulate the memory directly, access the data stored and create references to the variable.

// When a function receives a value type as an argument, it creates a copy of that data. This means if the original data is large, we now have two versions (imagine if one is about 200GB, that’d be a lot to work with) of it in memory leading to increased memory usage and potential performance issues.

package main

import (
	"fmt"
)

// func main() {
// 	var x int = 10
// 	var p *int = &x // p is a pointer to an integer, and it holds the address of x

// 	fmt.Printf("Value of x: %d\n", x)          // Output: Value of x: 10
// 	fmt.Printf("Address of x: %p\n", &x)       // Output: Address of x: 0xc0000140a8 (example)
// 	fmt.Printf("Value of p: %p\n", p)          // Output: Value of p: 0xc0000140a8 (same as address of x)
// 	fmt.Printf("Value at address p: %d\n", *p) // Output: Value at address p: 10

// 	*p = 20

// 	fmt.Printf("New value of x: %d\n", x)    // Output: New value of x: 20
// 	fmt.Printf("New address of x: %p\n", &x) // Output: New address of x: 0xc0000140a8 (same as address of p)
// }

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
