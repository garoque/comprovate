package product

import "github.com/garoque/comprovate/model"

type Database interface {
	FindAll() ([]model.Product, error)
}
