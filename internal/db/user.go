package db

import (
	"fmt"
	"github/binweee/picobin/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	RoleID   uint   `gorm:"column:role_id;not null"`
}

func AuthenticateUser(username, password string) (string, string) {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return "用户不存在", ""
	} else if result.Error != nil {
		return "错误", ""
	}
	if user.Password != password {
		return "密码错误", ""
	}
	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		fmt.Println(err)
	}

	return "登陆成功", token
}

func CreateUser(user User) {
	DB.Create(&user)
}
