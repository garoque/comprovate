package user

import "github.com/garoque/comprovate/model"

type Database interface {
	FindByEmail(email string) (*model.User, error)
}
