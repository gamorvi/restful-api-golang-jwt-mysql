package models

import (
	"fmt"
	"strings"

	u "github.com/gamorvi/restapi2/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"type:varchar(100);unique_index"`
	Password string
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if user.Name == "" {
		return u.Message(false, "Name should be on the payload"), false
	}

	if user.Username == "" {
		return u.Message(false, "Email should be on the payload"), false
	}

	if !strings.Contains(user.Username, "@") {
		return u.Message(false, "Email address is required"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		u.Message(false, "There was an internal error")
		return nil
	}
	user.Password = string(hash)

	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}

func GetUser(id int) *User {
	user := &User{}
	err := GetDB().First(&user, id).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUsers() []*User {
	users := make([]*User, 0)
	err := GetDB().Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return users
}

func UpdateUser(user *User) (err error) {

	err = GetDB().Save(user).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func DeleteUser(user *User) (err error) {
	err = GetDB().Delete(user).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

func GetUserForUpdateOrDelete(id int, user *User) (err error) {
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsername(username string) *User {
	user := &User{}
	if err := GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return user
}
