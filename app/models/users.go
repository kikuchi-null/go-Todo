package models

import (
	"time"
)

type User struct {
	User_ID   int
	UUID      string
	Name      string
	Email     string
	Password  string
	Create_At time.Time
	Update_At time.Time
}

func (u *User) CreateUser() (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Create(&u)
	return result.Error
}

func GetUserByID(user_id int) (user User, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("user_id = ?", user_id).Find(&user)
	return user, result.Error
}

func GetUserByEmail(email string) (user User, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("email = ?", email).Find(&user)

	return user, result.Error
}

func (u *User) UpdateUser(newName, newEmail string) (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Model(&u).Where("user_id = ?", u.User_ID).Update(
		map[string]interface{}{"name": newName, "email": newEmail})

	return result.Error
}

func (u *User) DeleteUser() (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("user_id = ?", u.User_ID).Delete(&u)

	return result.Error
}
