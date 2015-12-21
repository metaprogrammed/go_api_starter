package resources

import (
	"fmt"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

var currentId = 0
var AllUsers Users

// Make some example data
func init() {
	CreateUser(User{Username: "user1", Name: "John", Email: "john@example.com"})
	CreateUser(User{Username: "user2", Name: "Bill", Email: "bill@example.com"})
}

func FindUser(id int) User {
	for _, user := range AllUsers {
		if user.Id == id {
			return user
		}
	}
	// return empty User if not found
	return User{}
}

func CreateUser(user User) User {
	currentId += 1
	user.Id = currentId
	AllUsers = append(AllUsers, user)
	return user
}

func DestroyUser(id int) error {
	for i, user := range AllUsers {
		if user.Id == id {
			AllUsers = append(AllUsers[:i], AllUsers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d.", id)
}
