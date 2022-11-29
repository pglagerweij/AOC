package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Python Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(remainingTickets, userTickets, firstName, lastName, email)
		// Make it concurrent
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of all our bookings: %v.\n", firstNames)

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

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets now")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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

	// Create struct for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thanks you %v %v for booking %v ticket(s). You will receive an email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining out of %v for %v.\n", remainingTickets, conferenceTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############################")
	wg.Done()
}
