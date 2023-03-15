package migration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/garoque/comprovate/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertMockedUser(db *gorm.DB) error {
	user := model.User{ID: "1", Name: "Gabriel", Email: "g@g.com", Password: "123456"}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("migration.InsertMockedUser.bcrypt.GenerateFromPassword err: ", err.Error())
		return err
	}

	user.Password = string(hash)

	return db.Create(user).Error
}

func GetMockedProducts() []model.Product {
	file, err := os.ReadFile("./mockProducts.json")
	if err != nil {
		log.Fatal("migration.GetMockedProducts.os.ReadFile err: ", err.Error())
	}

	var products []model.Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		log.Fatal("migration.GetMockedProducts.json.Unmarshal err: ", err.Error())
	}

	return products
}

func InsertMockedProducts(db *gorm.DB, products []model.Product) error {
	return db.CreateInBatches(products, len(products)).Error
}
