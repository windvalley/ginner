package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"ginner/db/rdb"
)

// User user table
type User struct {
	Model
	Username string `json:"username" gorm:"column:username;not null"`
	Password string `json:"password" gorm:"column:password;not null"`
}

// Create insert a record
func (u *User) Create() error {
	return rdb.DBs.MySQL.Create(&u).Error
}

// GetUser get user by username
func GetUser(username string) (*User, error) {
	u := &User{}
	d := rdb.DBs.MySQL.Where("username = ?", username).First(&u)
	return u, d.Error
}

// Update update an user record
func (u *User) Update() error {
	return rdb.DBs.MySQL.Save(u).Error
}

// DeleteUser delete an user by user id
func DeleteUser(id uint) error {
	user := User{}
	user.ID = id
	return rdb.DBs.MySQL.Delete(&user).Error
}

// ListUser list users by username, offset, limit
func ListUser(username string, offset, limit int) ([]*User, uint, error) {
	users := make([]*User, 0)
	var count uint
	where := fmt.Sprintf("username like '%s%%'", username)
	if err := rdb.DBs.MySQL.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := rdb.DBs.MySQL.Where(where).Offset(offset).Limit(limit).Order(
		"id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// EncryptPassword encrypt user's password before write in database
func (u *User) EncryptPassword() (err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	u.Password = string(hashedBytes)
	return err
}

// CheckPassword check validation of user's password
func (u *User) CheckPassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}
