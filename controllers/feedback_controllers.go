package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"strings"

	"go-feedback-app/database"
	"go-feedback-app/middleware"
	"go-feedback-app/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return middleware.JwtKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "user" {
		http.Error(w, "Only users can submit feedback", http.StatusForbidden)
		return
	}

	userID := uint(claims["user_id"].(float64))
	var feedback models.Feedback
	json.NewDecoder(r.Body).Decode(&feedback)
	feedback.UserID = userID

	database.DB.Create(&feedback)
	json.NewEncoder(w).Encode(map[string]string{"message": "Feedback submitted"})
}

func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var feedback models.Feedback
	if err := database.DB.First(&feedback, id).Error; err != nil {
		http.Error(w, "Feedback not found", http.StatusNotFound)
		return
	}

	if user.Role != "admin" && feedback.UserID != user.ID {
		http.Error(w, "Unauthorized to delete this feedback", http.StatusForbidden)
		return
	}

	database.DB.Delete(&feedback)
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
}

func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	var feedbacks []models.Feedback
	database.DB.Find(&feedbacks)
	json.NewEncoder(w).Encode(feedbacks)
}
