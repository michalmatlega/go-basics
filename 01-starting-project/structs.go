package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, err := user.NewAdmin("test@text.com", "pass")

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	//mozna nie podawac keysow ale uwazaj na kolejnosc!
	//nie musisz wszystkich podawac, ustawi sie null value danego typu

	//appUser = user{} //Struct Literal or Composite Literal

	// ... do something awesome with that gathered data!

	//fmt.Println(firstName, lastName, birthdate)
	//outputUserDetails(appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

//func outputUserDetailsByPointer(u *user) {
//	fmt.Println(u.firstName, u.lastName, (*u).birthday) //uzywajac pointera mozna uzywac shortcuta u. lub zrobic konwersje na wartosc (*u). (sugar syntax)
//}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	//fmt.Scan(&value)	//nie wpisujac nic i wciskajac enter Scan wciaz czeka na input
	fmt.Scanln(&value) //Scanln czeka na nowa linie czyli enter
	return value
}
