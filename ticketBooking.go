package main

import (
	// "booking-app/helper"
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTicket = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings [50]string

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

func main() {

	greetUser()

	// fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n ", conferenceTicket, remainingTickets, conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still avaialbe \n", conferenceTicket, remainingTickets)

	// fmt.Println("Get your tickets to attend ")

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// fmt.Printf("These are all our bookings: %v\n", bookings)
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

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

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaialbe \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets to attend ")
}

func getFirstName() []string {

	firstNames := []string{}

	for _, booking := range bookings {
		// var name = strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remeaining for %v\n ", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########")

}
