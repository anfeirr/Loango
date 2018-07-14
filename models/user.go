package models

import (
	. "Loango/database"
	"errors"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name     string `gorm:"size:255" json:"Name"`
	Password string `gorm:"size:255" json:"Password"`
}

func (u *User) Creat() error {
	if u.Name != "" {
		db := InitDB()
		defer db.Close()

		return db.Create(&u).Error
	}
	return errors.New("User name is nil")

}

func GetAllUsers() (*[]User, error) {
	db := InitDB()
	defer db.Close()

	var users []User
	err := db.Find(&users).Error
	return &users, err
}

func GetUserByName(name string) (*User ,error){
	db := InitDB()
	defer db.Close()

	var user User
	err := db.Find(&user).Where("name=",name).Error
	return &user,err
}

func DeleteUserById(id string) error {
	if id != "2" {
		db := InitDB()
		defer db.Close()
		return db.Where("ID = ?", id).Delete(&User{}).Error
	} else {
		err := errors.New("admin can't be deleted ")
		return err
	}

}


