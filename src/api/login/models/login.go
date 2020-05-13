package models

import (
	"grpc_blog/src/models"
	"time"
)

type Users struct {
	ID              int
	Name            string
	Avatar          string
	Email           string
	EmailVerifiedAt string
	Password        string
	RememberToken   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func GetUsersByEmail(email string) (*Users, error) {
	u := Users{}
	err := models.DB.Where("email = ?", email).First(&u).Error
	return &u, err
}
