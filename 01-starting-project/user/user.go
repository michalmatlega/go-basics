package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     //anonymous embedding, you can do user: User but then you need to invoke methods by user.Foo()
}

func NewAdmin(email string, password string) (*Admin, error) {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthday:  "---",
			createdAt: time.Now(),
		},
	}, nil
}

// V-----RECEIVER ARGUMENT (or just RECEIVER)
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name and last name and birth date must not be empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthDate,
		createdAt: time.Now(),
	}, nil
}
