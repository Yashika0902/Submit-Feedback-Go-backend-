package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-feedback-app/database"
	"go-feedback-app/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Role != "admin" && user.Role != "user" {
		http.Error(w, "Role must be 'admin' or 'user'", http.StatusBadRequest)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var dbUser models.User
	json.NewDecoder(r.Body).Decode(&input)

	if err := database.DB.Where("username = ?", input.Username).First(&dbUser).Error; err != nil {
		http.Error(w, "Invalid username", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"role":     dbUser.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
