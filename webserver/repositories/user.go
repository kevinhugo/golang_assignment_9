package repositories

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Users = []User{
	{
		ID:      1,
		Name:    "Hacktiv8",
		Address: "Jakarta",
	},
}

func CreateUser(user *User) error {
	Users = append(Users, *user)
	return nil
}

func GetUsers(id int) ([]User, error) {
	if len(Users) == 0 {
		return nil, fmt.Errorf("No Data !")
	}

	if id != 0 {
		var oneUser []User
		if len(Users) < id {
			return oneUser, fmt.Errorf("Index Beyond My Reach")
		}
		oneUser = append(oneUser, Users[id-1])
		return oneUser, nil
	}

	return Users, nil
}
