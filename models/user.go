package models

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username   string `json:"username" binding:"required" gorm:"type:varchar(100);unique_index"`
	Password   string `json:"password" binding:"required"`
	SessionKey string
}

func (u *User) Index() {
	users := []User{}
	db.Find(&users)
	fmt.Print(users)
}

func CheckToken(token string) (User, error) {
	var user User

	// query user
	if db.Where("session_key = ?", token).First(&user).RecordNotFound() {
		return user, errors.New("Wrong credentials")
	}

	return user, nil
}

func (u *User) Register() error {
	// db.Create(value)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Failed to register user")
	}

	u.Password = string(hashedPassword)
	db.Create(u)

	if db.NewRecord(&u) {
		return errors.New("Failed to register user. Username taken.")
	}

	return nil
}

func (u *User) Login() (User, error) {
	var user User

	// query user
	if db.Where("username = ?", u.Username).First(&user).RecordNotFound() {
		return user, errors.New("Wrong credentials")
	}

	// validate password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return user, errors.New("Wrong credentials")
	}

	// TODO: make utils function for this

	// Generate session key, this is supposed to be derived from secret_key config
	// That must be different each instance
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// make 32-char session-key
	runes := make([]rune, 32)
	for i := range runes {
		runes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	sessionKey := string(runes)

	user.SessionKey = sessionKey
	db.Save(&user)

	return user, nil
}
