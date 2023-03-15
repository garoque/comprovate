package main

import (
	"log"
	"net/http"
	"os"

	apiProduct "github.com/garoque/comprovate/api/product"
	apiUser "github.com/garoque/comprovate/api/user"
	"github.com/garoque/comprovate/database/product"
	"github.com/garoque/comprovate/database/user"

	"github.com/garoque/comprovate/migration"

	"github.com/garoque/comprovate/model"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	os.Remove("database/comprovate.db")
}

func main() {
	db, err := gorm.Open(sqlite.Open("database/comprovate.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Product{}, &model.User{})
	dbUser := user.NewUser(db)
	dbProduct := product.NewProduct(db)

	products := migration.GetMockedProducts()
	err = migration.InsertMockedProducts(db, products)
	if err != nil {
		log.Fatal("main.migration.InsertMockedProducts err: ", err.Error())
	}

	err = migration.InsertMockedUser(db)
	if err != nil {
		log.Fatal("main.migration.InsertMockedUser err: ", err.Error())
	}
	log.Println("Success up migrations")

	handlerUser := apiUser.NewUserHandler(dbUser)
	handlerProduct := apiProduct.NewProductHandler(dbProduct)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", tokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", 10))

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/products", handlerProduct.GetProducts)
	})

	r.Post("/user/auth", handlerUser.Login)

	http.ListenAndServe(":3000", r)
}
