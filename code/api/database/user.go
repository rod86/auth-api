package database

import (
	"errors"
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string     `json:",omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func GetUsers() []User {
	var users []User
	db.Select("id, username, email, created_at, updated_at, deleted_at").Order("created_at DESC").Find(&users)
	return users
}

func GetUserById(id int) (bool, User) {
	var user User
	db.Where("id = ?", id).Select("id, username, email, created_at, updated_at, deleted_at").First(&user)
	if user.ID == 0 {
		return false, user
	}
	return true, user
}

func FindOneByEmailOrUsername(username, email string, excludeId int) (bool, User) {
	var user User

	if excludeId > 0 {
		db.Where("(username = ? OR email = ?) AND id != ?", username, email).First(&user)
	} else {
		db.Where("username = ? OR email = ?", username, email).First(&user)
	}

	if user.ID == 0 {
		return false, user
	}
	return true, user
}

func CreateUser(username, email, password string) error {
	var user = User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(id int, data map[string]interface{}) error {
	ok, user := GetUserById(id)
	if !ok {
		return errors.New("User to update not found")
	}
	err := db.Model(&user).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	err := db.Where("id = ?", id).Delete(User{}).Error
	if err != nil {
		return err
	}
	return nil
}
