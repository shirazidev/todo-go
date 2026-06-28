package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	shirazidevLogo "todo-app/shirazidev-logo"
)

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

type Task struct {
	ID          uint
	Title       string
	Description string
	DueDate     string
	IsCompleted bool
	CategoryId  int
	UserId      int
}

type TaskCategory struct {
	ID     int
	Title  string
	Color  string
	UserId int
}

var (
	taskStorage       []Task
	userStorage       []User
	categoryStorage   []TaskCategory
	AuthenticatedUser *User
	serializationMode string
)

const (
	uPath                   string = "users.txt"
	ManualSerializationMode string = "manual"
	JsonSerializationMode   string = "json"
)

func main() {
	clearScreen()
	fmt.Println("Welcome to TODO App!")

	serializeMode := flag.String("serializeMode", "json", "Serialization mode to write data to file")

	command := flag.String("command", "", "command to execute")
	flag.Parse()

	switch *serializeMode {
	case "manual":
		serializationMode = ManualSerializationMode
	default:
		serializationMode = JsonSerializationMode
	}
	loadUserStorageFromFile()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		runCommand(*command)
		fmt.Println("\nPlease enter a command: ")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	if command != "register-user" && command != "login" && command != "exit" && AuthenticatedUser == nil {
		login()

		if AuthenticatedUser == nil {
			return
		}

	}
	switch command {
	case "create-task":
		clearScreen()
		createTask()
	case "list-tasks":
		clearScreen()
		listTasks()
	case "create-category":
		clearScreen()
		createCategory()
	case "register-user":
		clearScreen()
		registerUser()
	case "login":
		clearScreen()
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
	fmt.Println("Authenticated user: ", AuthenticatedUser.Email)
	scanner := bufio.NewScanner(os.Stdin)

	var name, duedate, description string
	//var categoryId uint

	fmt.Println("Enter the name of the task: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Enter the description of the task: ")
	scanner.Scan()
	description = scanner.Text()

	fmt.Println("Enter the Category Id of the task: ")
	scanner.Scan()
	cid := scanner.Text()

	categoryID, err := strconv.Atoi(cid)
	if err != nil {
		fmt.Println("Invalid Category Id", err)
		return
	}
	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserId == AuthenticatedUser.Id {
			isFound = true
			break
		}
		if !isFound {
			fmt.Println("Invalid Category Id", err)
			return
		}
	}

	fmt.Println("Enter the duedate of the task: ")
	scanner.Scan()
	duedate = scanner.Text()

	t := Task{
		ID:          uint(len(taskStorage) + 1),
		Title:       name,
		Description: description,
		DueDate:     duedate,
		CategoryId:  categoryID,
		UserId:      AuthenticatedUser.Id,
		IsCompleted: false,
	}
	taskStorage = append(taskStorage, t)

	fmt.Println("title: ", t.Title, " due date: ", t.DueDate, " id: ", t.ID, " userId: ", t.UserId)
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

	c := TaskCategory{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserId: AuthenticatedUser.Id,
	}
	categoryStorage = append(categoryStorage, c)
	fmt.Println("Category created successfully!")
}

func registerUser() {
	fmt.Println("===== Registering user =====")

	scanner := bufio.NewScanner(os.Stdin)

	var email, password, name string

	fmt.Println("Enter your name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Enter the email of the user: ")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("Enter the password of the user: ")
	scanner.Scan()
	password = scanner.Text()
	defer fmt.Println("email: ", email, "password: ", password)

	u := User{Id: len(userStorage) + 1, Email: email, Password: password, Name: name}
	userStorage = append(userStorage, u)
	saveUser(u)
}

func login() {
	// get email and password from user to login
	fmt.Println("===== Login =====")
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
		if user.Email == email && user.Password == password {
			notFound = false
			AuthenticatedUser = &user
			clearScreen()
			break
		}
	}
	if notFound {
		fmt.Println("Invalid credentials!")
	}

}
func listTasks() {
	for _, t := range taskStorage {
		if t.UserId == AuthenticatedUser.Id {
			fmt.Println("\nId: ", t.ID, "\nTitle: ", t.Title, "\nDescription: ", t.Description)
			fmt.Println("========")
		} else {
			fmt.Println("Not found: ", t.UserId)
			break
		}
	}
}

func (u User) print() {
	fmt.Println("Id: ", u.Id, "Email: ", u.Email)
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	shirazidevLogo.Plogo()
	if AuthenticatedUser != nil {
		fmt.Printf(``)
		fmt.Println("=========\nLogged in!")
		fmt.Printf("User ID: %d, Email: %v\n========\n", AuthenticatedUser.Id, AuthenticatedUser.Email)
	}
	if err != nil {
		return
	}
}

func loadUserStorageFromFile() {
	file, err := os.Open(uPath)
	if err != nil {
		fmt.Println("Cannot access user records: ", err)
	}
	defer file.Close()
	var data = make([]byte, 10240)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("Cannot get user records: ", oErr)
	}

	var dataStr string = string(data)

	userSlice := strings.Split(dataStr, "\n")
	var userStruct *User
	for _, u := range userSlice {
		switch serializationMode {
		case JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("Cannot parse user records: ", uErr)
			}
		case ManualSerializationMode:
			var dErr error
			userStruct, dErr = deserializeFromManual(u)
			if dErr != nil {
				fmt.Println("Error deserializing user records: ", dErr)
				return
			}
		default:
			fmt.Println("Invalid serialization mode")
			return
		}
		if userStruct != nil {
			userStorage = append(userStorage, *userStruct)
			fmt.Println("User records loaded successfully!")
		} else {
			fmt.Println("User records could not be loaded successfully!")
		}
	}

}

func deserializeFromManual(userStr string) (*User, error) {
	u := strings.TrimSpace(userStr)
	if u == "" {
		return &User{}, errors.New("Empty user")
	}
	var user User
	userFields := strings.Split(u, ", ")

	for _, field := range userFields {
		values := strings.SplitN(field, ": ", 2)
		if len(values) != 2 {
			fmt.Println("field is not valid, skipping... ", field)
			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]
		switch fieldName {
		case "ID":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				return &User{}, errors.New("Strconv Err")
			}
			user.Id = id
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		case "name":
			user.Name = fieldValue
		}
	}
	return &user, nil
}
func saveUser(u User) {
	var file *os.File

	// save user data to storage
	var oErr error
	file, oErr = os.OpenFile(uPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if oErr != nil {
		fmt.Println("I cannot open the file!", oErr)
		return
	}
	defer file.Close()
	var data = make([]byte, 10240)
	if serializationMode == ManualSerializationMode {
		data = []byte(fmt.Sprintf("ID: %d, email: %s, name: %s, password: %s\n", u.Id, u.Email, u.Name, u.Password))
	} else if serializationMode == JsonSerializationMode {
		//json
		var jErr error
		data, jErr = json.Marshal(u)
		data = append(data, '\n')
		if jErr != nil {
			fmt.Println("Cannot marshal user struct to json: ", jErr)
		}

	} else {
		fmt.Println("Invalid serialization mode")
		return
	}

	_, wErr := file.Write(data)
	if wErr != nil {
		fmt.Println("Write error: ", wErr)
		return
	}
}
