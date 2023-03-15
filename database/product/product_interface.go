package product

import (
	"context"

	"github.com/garoque/comprovate/model"
)

type Database interface {
	FindAll(ctx context.Context) ([]model.Product, error)
}
