package rdb

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Username string `json:"username" gorm:"column:username;not null"`
	Password string `json:"password" gorm:"column:password;not null"`
}

func (u *User) Create() error {
	return DBs.MySQL.Create(&u).Error
}

func GetUser(username string) (*User, error) {
	u := &User{}
	d := DBs.MySQL.Where("username = ?", username).First(&u)
	return u, d.Error
}

func (u *User) Update() error {
	return DBs.MySQL.Save(u).Error
}

func DeleteUser(id uint) error {
	user := User{}
	user.ID = id
	return DBs.MySQL.Delete(&user).Error
}

func ListUser(username string, offset, limit int) ([]*User, uint, error) {
	users := make([]*User, 0)
	var count uint
	where := fmt.Sprintf("username like '%s%%'", username)
	if err := DBs.MySQL.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DBs.MySQL.Where(where).Offset(offset).Limit(limit).Order(
		"id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (u *User) EncryptPassword() (err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	u.Password = string(hashedBytes)
	return err
}

func (u *User) CheckPassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}
