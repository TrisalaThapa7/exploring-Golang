package main

import (
	"fmt"
	"strings"
)

func main() {
	var conference_name string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("conference_name is %T, conferenceTickets is %T, remainingTickets is %T. \n", conference_name, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our %v booking application.\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v conferenceTickets.\n", remainingTickets, conferenceTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {

				fmt.Println("Our conference is fully booked, please come next time!")
				break
			}

		} else {

			fmt.Println("Your input is invalid.")

		}

	}

}
