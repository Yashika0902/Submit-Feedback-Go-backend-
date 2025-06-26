package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id"`
}
