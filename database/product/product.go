package product

import (
	"github.com/garoque/comprovate/model"
	"gorm.io/gorm"
)

type product struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) Database {
	return &product{db}
}

func (p *product) FindAll() ([]model.Product, error) {
	var products []model.Product

	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
