package model

import "github.com/aymerick/raymond"

type User struct {
	ID           string
	Name         string
	ProfileImage raymond.SafeString `gorm:"column:profile_image" handlebars:"profile_image"`
	Bio          string
}

type Users []User

// Set User's table name to be `profiles`
func (User) TableName() string {
	return "users"
}

func (user User) String() string {
	return user.Name
}

func (users Users) String() string {
	s := ""
	for i, u := range users {
		if i != 0 {
			s += ", "
		}
		s += u.Name
	}
	return s
}
