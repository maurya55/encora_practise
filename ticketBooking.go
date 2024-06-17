package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50
	// var bookings [50]string
	bookings := []string{}

	fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n ", conferenceTicket, remainingTickets, conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaialbe \n", conferenceTicket, remainingTickets)

	fmt.Println("Get your tickets to attend ")

	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			// userName = "Tom"
			// userTickets = 2
			// fmt.Println(&userName)
			remainingTickets = remainingTickets - userTickets

			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remeaining for %v\n ", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var name = strings.Fields(booking)
				firstNames = append(firstNames, name[0])

			}

			fmt.Printf("The first names of bookings are %v\n", firstNames)
			// fmt.Printf("These are all our bookings: %v\n", bookings)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name  you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email addres you entered is not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you enterd is invalid ")
			}
			// fmt.Println("Your input data is invalid, try again")
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}

}
