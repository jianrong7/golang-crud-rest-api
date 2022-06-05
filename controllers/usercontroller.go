package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	var hashedPassword, err = HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	database.Instance.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var users []entities.User
	database.Instance.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func checkIfUserExists(userId string) bool {
	var user entities.User
	database.Instance.First(&user, userId)
	return user.ID != 0
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var user entities.User
	database.Instance.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}