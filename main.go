package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	Id       int
	Email    string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("Welcome to TODO App!")

	command := flag.String("command", "", "command to execute")
	flag.Parse()

	if *command != "register-user" && *command != "exit" {

	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		runCommand(*command)
		fmt.Println("\nPlease enter a command: ")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		fmt.Printf("User storage: %+v", userStorage)
		os.Exit(0)
	default:
		if command == "" {
			fmt.Println("no command provided!")
		} else {
			fmt.Println("Invalid command: ", command)
		}
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)

	var name, duedate, category string

	fmt.Println("Enter the name of the task: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Enter the category of the task: ")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Enter the duedate of the task: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("title: ", name, " due date: ", duedate, " category: ", category)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string
	fmt.Println("Enter the name of the category: ")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("Enter the color of the category: ")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("category name: ", title, " category color: ", color)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)

	var email, password string

	fmt.Println("Enter the email of the user: ")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("Enter the password of the user: ")
	scanner.Scan()
	password = scanner.Text()
	defer fmt.Println("email: ", email, "password: ", password)

	u := User{Id: len(userStorage) + 1, Email: email, Password: password}
	userStorage = append(userStorage, u)
}

func login() {
	// get email and password from user to login
	fmt.Println("You must login first!")
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your email: ")
	sc.Scan()
	email := sc.Text()
	fmt.Println("Please enter your password: ")
	sc.Scan()
	password := sc.Text()
	// if there is a user record with corresponding data, allow user to continue!
	notFound := true
	for _, user := range userStorage {
		if user.Email == email || user.Password == password {
			notFound = false
		}
	}
	if notFound {
		fmt.Println("There is no user with this credintials!")
		return
	}

}
