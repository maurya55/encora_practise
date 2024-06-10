package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n ", conferenceTicket, remainingTickets, conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaialbe \n", conferenceTicket, remainingTickets)

	fmt.Println("Get your tickets to attend ")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	// userName = "Tom"
	// userTickets = 2
	// fmt.Println(&userName)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remeaining for %v\n ", remainingTickets, conferenceName)

}
