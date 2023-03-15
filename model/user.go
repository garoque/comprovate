package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
