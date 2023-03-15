package product

import (
	"context"

	"github.com/garoque/comprovate/model"
	"gorm.io/gorm"
)

type product struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) Database {
	return &product{db}
}

func (p *product) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	if err := p.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
