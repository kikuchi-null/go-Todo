package models

import (
	"fmt"
	"time"
)

type Session struct {
	Session_id int
	User_ID    int
	UUID       string
	Email      string
	Create_At  time.Time
}

func (u *User) CreateSession() (session Session, err error) {
	DB := gormConnect()
	defer DB.Close()

	session = Session{
		User_ID:   u.User_ID,
		UUID:      u.UUID,
		Email:     u.Email,
		Create_At: time.Now(),
	}

	DB.Create(&session)

	result := DB.Where("user_id = ? AND email = ?", u.User_ID, u.Email).Find(&session)

	return session, result.Error
}

func (sess *Session) IsSession() (valid bool, err error) {
	DB := gormConnect()
	defer DB.Close()

	var session Session
	result := DB.Where("uuid = ?", sess.UUID).Find(&session)

	if result.Error != nil {
		fmt.Println(result.Error, result.Value)
		valid = false
		return
	}

	if session.User_ID != 0 {
		valid = true
	}

	return valid, result.Error
}

func (sess *Session) GetUserBySession() (user User, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("uuid = ?", sess.UUID).Find(&user)

	return user, result.Error
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("uuid = ?", sess.UUID).Delete(&sess)

	return result.Error
}
