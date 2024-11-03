package main

import "fmt"

type User struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" validate:"required,email" gorm:"unique;"`
	Role     string
}

func main() {
	var user *User
	fmt.Println(user == nil)
}
