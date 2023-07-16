package models

import "RecordWords/dao"

type User struct {
	ID       string `json:"id"`
	Number   int    `json:"number"`
	Password int    `json:"password"`
}

func CreateUser(user *User) (err error) {
	err = dao.DBUser.Create(&user).Error
	return
}
