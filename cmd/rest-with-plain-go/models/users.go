package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Mobile    string `json:"mobile"`
	Country   string `json:"Country"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Users []*User

var ErrUserNotFound = fmt.Errorf("User not found")

var userList = []*User{
	{
		Id:        1,
		FirstName: "Debi",
		LastName:  "Kenlin",
		Email:     "dkenlin0@nationalgeographic.com",
		Gender:    "Female",
		Mobile:    "1728790050",
		Country:   "Democratic Republic of the Congo",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		Id:        2,
		FirstName: "Thadeus",
		LastName:  "Theuff",
		Email:     "ttheuffa@wiley.com",
		Gender:    "Male",
		Mobile:    "6537048360",
		Country:   "Ivory Coast",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		Id:        3,
		FirstName: "Rosalie",
		LastName:  "Le Brom",
		Email:     "rlebrom4@joomla.org",
		Gender:    "Male",
		Mobile:    "3058511121",
		Country:   "Japan",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func findNextId() int {
	lastUser := userList[len(userList)-1]
	return lastUser.Id + 1
}

func findUserWithId(id int) (*User, int, error) {
	for pos, user := range userList {
		if user.Id == id {
			return user, pos, nil
		}
	}

	return nil, -1, ErrUserNotFound
}

func GetUsers() Users {
	return userList
}

func AddUser(user *User) {
	user.Id = findNextId()
	userList = append(userList, user)
}

func UpdateUser(id int, user *User) error {
	_, pos, err := findUserWithId(id)
	if err != nil {
		return err
	}
	//@todo Need to support partial updates
	userList[pos] = user
	return nil
}
