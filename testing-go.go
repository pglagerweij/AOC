package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Python Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of all our bookings: %v.\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Your First name or last Name is too short, try again.\n")
			}
			if !isValidEmail {
				fmt.Printf("Your Email address is incorect, doesn't contain an @ sign, try again.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining so you can't book %v tickets.\n", remainingTickets, userTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets now")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their input
	fmt.Println("Enter your user name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adddress: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thanks you %v %v for booking %v ticket(s). You will receive an email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining out of %v for %v.\n", remainingTickets, conferenceTickets, conferenceName)
}
