package user

import (
	"context"

	"github.com/garoque/comprovate/model"
)

type Database interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
