package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTicket = 50

	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Println("Remaining Ticket :", remainingTicket, "\nAvailable ticket", conferenceTickets)

	var userName string
	var userTicket int
	// for name

	userName = "Tom"
	userTicket = 2
	fmt.Printf("%v booked %v tickets\n", userName, userTicket)
	fmt.Scan(&userName)
	fmt.Printf("Name %v", userName)
}
