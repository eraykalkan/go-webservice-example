package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

// AddUser return User and error
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include id.")
	}
	u.ID = nextID
	nextID++
	// users array holding pointers to user objects so we need to pass the reference
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	// write only variable to ignore the key, we only need value u
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	// ErrorF function allows us to return a formatted string as an error object
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			// splice operation, take from index i and append it to the array after the index i+1 - delete
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
