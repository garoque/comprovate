package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/garoque/comprovate/database/user"

	"github.com/garoque/comprovate/model"
	"github.com/go-chi/jwtauth"
)

type UserHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	db user.Database
}

func NewUserHandler(db user.Database) UserHandlerInterface {
	return &handler{db}
}

func (u *handler) Login(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var input model.UserLogin

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user, err := u.db.FindByEmail(r.Context(), input.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if !user.IsValidPassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := make(map[string]interface{})
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix()

	_, token, err := jwt.Encode(claims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	accessToken := model.JWTResponse{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}
