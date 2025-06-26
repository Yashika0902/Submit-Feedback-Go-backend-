// main.go
package main

import (
	"go-feedback-app/controllers"
	"go-feedback-app/database"
	"go-feedback-app/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Authenticated routes
	r.Handle("/feedback", middleware.AuthMiddleware(http.HandlerFunc(controllers.SubmitFeedback))).Methods("POST")
	r.Handle("/feedback/{id}", middleware.AuthMiddleware(http.HandlerFunc(controllers.DeleteFeedback))).Methods("DELETE")

	// Admin-only
	r.Handle("/feedbacks", middleware.AdminOnly(http.HandlerFunc(controllers.GetFeedbacks))).Methods("GET")

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
