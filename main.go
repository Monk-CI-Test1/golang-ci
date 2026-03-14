package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Wrap r.Run in a log.Fatal to catch and report startup errors
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/niteshjangid29/practice/controller"
// )

// type UserData struct {
// 	Name            string
// 	Email           string
// 	NumberOfTickets uint
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	var confName = "Go Conference"
// 	const confTickets = 50
// 	var remTickets uint = 50

// 	greetUsers("John")
// 	// helper()
// 	controller.MyController()

// 	fmt.Printf("Welcome to %v booking application\n", confName)
// 	fmt.Println("Total tickets available:", confTickets)
// 	fmt.Printf("Types are %T %T %T\n", confName, confTickets, remTickets)
// 	fmt.Println("Remaining tickets:", remTickets)
// 	fmt.Println("Get your tickets here")

// 	// var bookings = [50]string{"A1", "A2", "A3"}
// 	// var bookings [50]string
// 	var bookings []string

// 	var userData = map[string]string{
// 		"name":  "John",
// 		"email": "john@example.com",
// 	}

// 	fmt.Println("User Data:", userData)

// 	// bookings[0] = "A1"
// 	// bookings = append(bookings, "A1slice")

// 	for remTickets > 40 {
// 		var userName string
// 		var userTickets uint

// 		fmt.Println("Enter your name:")
// 		fmt.Scan(&userName)
// 		fmt.Println("Enter the number of tickets you want to book:")
// 		fmt.Scan(&userTickets)

// 		remTickets = remTickets - userTickets
// 		bookings = append(bookings, userName)

// 		names := []string{}

// 		for index, booking := range bookings {
// 			fmt.Printf("Index: %v, Value: %v\n", index, booking)
// 			names = append(names, booking)
// 		}

// 		fmt.Printf("User: %v booked %v tickets\n", userName, userTickets)
// 		fmt.Printf("Remaining tickets: %v\n", remTickets)

// 		wg.Add(1)
// 		go sendTicket(userName, userTickets)

// 		if remTickets <= 0 {
// 			fmt.Println("Sorry, all tickets are booked!")
// 			break
// 		}
// 		wg.Wait()
// 	}

// 	var city string = "New York"

// 	switch city {
// 	case "New York":
// 		fmt.Println("The city that never sleeps")
// 	case "Paris":
// 		fmt.Println("The city of love")
// 	case "Tokyo":
// 		fmt.Println("The city of technology")
// 	default:
// 		fmt.Println("Unknown city")
// 	}

// 	// fmt.Printf("Bookings = %v\n", bookings)
// }

// func greetUsers(name string) {
// 	fmt.Printf("Welcome to our conference booking application, %v!\n", name)
// }

// func sendTicket(userName string, userTickets uint) {
// 	time.Sleep(10 * time.Second)
// 	fmt.Println("############################")
// 	fmt.Printf("Sending %v tickets to %v\n", userTickets, userName)
// 	fmt.Println("############################")

// 	wg.Done()
// }

// package main

// import (
// 	"time"
// )

// type User struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }

// func main() {

// var revenue float64
// var expenses float64
// var taxRate float64

// fmt.Println("Enter revenue: ")
// fmt.Scan(&revenue)
// fmt.Println("Enter expenses: ")
// fmt.Scan(&expenses)
// fmt.Println("Enter tax rate: ")
// fmt.Scan(&taxRate)

// ebt := revenue - expenses
// profit := ebt * (1 - taxRate/100)
// ratio := ebt / profit

// fmt.Printf("Earnings Before Tax: %v\n", ebt)
// fmt.Printf("Profit: %v\n", profit)
// fmt.Printf("Profit Ratio: %.1f%%\n", ratio*100)
// }

// func getUserInput(prompt string) (string, error) {
// 	fmt.Println(prompt)

// 	var input string
// 	fmt.Scan(&input)

// 	if input == "" {
// 		return "", fmt.Errorf("input cannot be empty")
// 	}

// 	return input, nil
// }
