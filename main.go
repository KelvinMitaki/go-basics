package main

import (
	"fmt"
	"slices"
	"time"
)

// STRUCTS
type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	birthDate   time.Time
}

// METHOD STRUCTS
func (u *User) printFirstName() string {
	return u.firstName
}

func main() {

	useSlices()

	useVariables()

	usePointers()

	useStruct()
}

func useVariables() {
	// VARIABLES
	fmt.Println("hello world")

	var whatToSay string = "goodbye"
	var i int = 1

	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, whatWasSaid2 := saySomething()

	fmt.Println(whatWasSaid, whatWasSaid2)
}

func saySomething() (string, string) {
	return "say something", "string 2"
}

func usePointers() {
	// POINTERS
	oldValue := "Green"

	fmt.Println("the old value is", oldValue)

	changeUsingPointer(&oldValue)

	fmt.Println("the new value is", oldValue)
}

func changeUsingPointer(s *string) {
	newValue := "Red"

	*s = newValue
}

func useStruct() {
	user := User{
		firstName: "Kelvin",
		lastName:  "Mitaki",
	}

	var user2 User

	user2.firstName = "John"
	user2.lastName = "Doe"

	fmt.Println(user.firstName, user.lastName)

	fmt.Println(user2.firstName, user2.lastName)

	fmt.Println("METHOD STRUCT", user.printFirstName())
}

func useMaps() {
	myMap := make(map[string]string)

	myMap["dog"] = "German Shephard"
	myMap["other-dog"] = "Rottweiler"

	fmt.Println(myMap["dog"])
	fmt.Println(myMap["other-dog"])

	myNewMap := make(map[string]User)

	me := User{
		firstName: "Kelvin", lastName: "Mitaki",
	}

	myNewMap["me"] = me

	fmt.Println(myNewMap["me"].firstName)
}

func useSlices() {
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	slices.Sort(mySlice)

	fmt.Println(mySlice)

	mySecondSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(mySecondSlice)
	fmt.Println(mySecondSlice[1:4])

	names := []string{"first", "second", "third"}

	fmt.Println(names)
}

func useDecisions() {
	num := 1

	if num > 0 {
		fmt.Println("number is greater than 0")
	}

	pet := "cat"

	switch pet {
	case "cat":
		fmt.Println("this is a cat")
	case "dog":
		fmt.Println("this is a dog")
	default:
		fmt.Println("this is neither a cat or a dog")
	}
}

func useLoops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numbers := []string{"one", "two", "three", "four", "five"}

	for j := 0; j < len(numbers); j++ {
		fmt.Println(numbers[j])
	}

	for k, number := range numbers {
		fmt.Println(k, number)
	}

	animals := make(map[string]string)

	animals["dog"] = "this is a dog"
	animals["cat"] = "this is a cat"
	animals["cow"] = "this is a cow"

	for _, animal := range animals {
		fmt.Println(animal)
	}

	name := "kelvin mitaki"

	for _, letter := range name {
		// RETURNS A STRING IN BYTES FORMAT
		fmt.Println(letter)
	}

	var users []User
	users = append(users, User{firstName: "john", lastName: "doe"})
	users = append(users, User{firstName: "mary", lastName: "jane"})
	users = append(users, User{firstName: "mark", lastName: "alex"})

	for _, user := range users {
		fmt.Println(user.firstName, user.lastName)
	}
}
