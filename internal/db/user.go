package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	RoleID   uint   `gorm:"column:role_id;not null"`
}

func AuthenticateUser(username, password string) string {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return "User not found"
	} else if result.Error != nil {
		return "error"
	}
	if user.Password != password {
		return "Incorrect password"
	}
	return "success"
}

func CreateUser(user User) {
	DB.Create(&user)
}
