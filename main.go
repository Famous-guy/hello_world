package main

import (
	"fmt"
	"strings"
)
 func main() {	

for {

  	 bookinglist := []string{}
	var user string = "david"
  fmt.Println(user)
var firstName string
var lastName string
var email string
numberofTickets := 50
var bookedTickets int
 var numberofTicketsRemaining int
 fmt.Println(concat("Hello, ", "World!"))
 fmt.Println(concat("Hello, ", "Famous!"))
 fmt.Println(add(5, 5))
 fmt.Println(sub(10, 5))
 fmt.Println(mul(5, 5))
 fmt.Println(divide(20, 5))
	 fmt.Println()
  fmt.Println("Enter your first name: ")
 fmt.Scan(&firstName)
  fmt.Println()
  fmt.Println("Enter your last name: ")
 fmt.Scan(&lastName)
  fmt.Println()
  fmt.Println("Enter your email address: ")
 fmt.Scan(&email)
  fmt.Println()
  fmt.Println("Enter number of Tickets to book: ")
 fmt.Scan(&bookedTickets)
  numberofTicketsRemaining =  numberofTickets - bookedTickets
  bookinglist = append(bookinglist, firstName + " " + lastName)
  fmt.Println()
  fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, bookedTickets, email)
  fmt.Printf("You have %v tickets remaining \n", numberofTicketsRemaining)

  firstNames := []string{}
  for _, booking := range bookinglist {
    var names = strings.Fields(booking)
    firstNames = append(firstNames, names[0])
  }

  fmt.Printf("this are all the bookings %v \n", firstNames)
  if numberofTicketsRemaining <= 0 {
    break
  } else {
    fmt.Printf("this is it")
  }
}
}
func concat(a string, b string) string{
return a+b
}

func add(a int, b int) int{
	return a+b
	
}

func sub(a int, b int) int{
	return a-b
	
}

func mul(a int, b int) int{
	return a*b
	
}
func divide(a int, b int) int{
	return a/b
	
}