// Everything we are using in application must belong to package
package main

// We need to explicitely import the packages whose properties we are using
import (
	"fmt"
	"sync"
	"time"
	// "strconv"
	// Not a built-in package thus need to explicitely mention module and package name
	"bookingApp/helper"
)

// Use documentation to know what belongs to which package

// Package level variables
var confName = "Go Conference"
const confTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// Create empty list of map with initial size 1 -> var bookings = make([]map[string]string, 1)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Creating waitgroup
var wg = sync.WaitGroup{}

func main() {

	// Function call
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Syntactic sugar in programming is only applicable to local variable and not consts
		// Also, if you want to define datatype explicitely, then too := cannot be used
		firstNames := getFirstNames()
		fmt.Printf("The First Names of Bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Printf("%v is booked out. Come back next year!\n", confName)
		}

	} else {

		if !isValidName {
			fmt.Println("The first or Last name you enteresd is too short.")
		}
		if !isValidEmail {
			fmt.Println("Entered email address is incorrect.")
		}
		if !isValidTicketNumber {
			fmt.Println("No. of Ticket entered is invalid.")
		}

	}
	wg.Wait() // blocks until wg counter is 0
}

func greetUsers() {
	// Printf is used for printing formatted data
	// Using placeholder, v stands for variable values
	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Printf("We have total of %v ickets and %v are available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// %T tells the type of variable
	// fmt.Printf("confTickets is %T type, remainingTickets is %T type and confName is %T type\n", confTickets, remainingTickets, confName)

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Asking for user's name by reading input value through Scan
	// Passing the memmory address to input data
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you need:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// Create map for users - In map, we cannot mix datatypes
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// FormatUint converts uint values and formats it into string type, 10 represents base 10- decimal no.s
	// userData["Tickets booked"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, confName)
}

func getFirstNames() []string {
	firstNames := []string{}

	// For arrays and slices range provide the index for ele
	// _ is used to ignore the unused variables thus replaced index with _

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Blocking execution of current thread
	time.Sleep(10 * time.Second)
	// Sprint method prints the output in console and saves result in string type variable
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("********************")
	fmt.Printf("Sending tickets:\n%v\nto email address: %v\n", ticket, email)
	fmt.Println("********************")
	wg.Done() // Removes the thread that we added - decrements the counter
}

/** In a single thread application, during the 10 sec delay the user will not get the information of next step,
  To optimize the program, thus, we'll make it multi-threaded so that blocking is avoided
  Golang for concurrency:-
  In order to do this, write "go" keyword (go routine) in front of the function call causing delay
  Main thread does not wait for any additional thread to complete
  Fix: creating a wait group. Now, main thread has to wait for a definite no. wait group before termination
*/

/** Go Uses "Green Thread", which is an abstraction of an actual thread (OS thread) -> Go Roautine
    We can create thousands of thread instead of worrying about low level OS thread
    Advantages: Cheaper, light weight, less memory space without affecting application performance
    Channels in Golang - Built-in functionality for Go routines to communicate with each other
    Java uses OS thread that needs more memory space
*/
