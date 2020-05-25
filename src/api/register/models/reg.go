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

func AddUser(u *Users) (int64, error) {
	stms, err := models.DB.Prepare("INSERT INTO users(name,avatar,email,password,created_at,updated_at) values(?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	result, err := stms.Exec(u.Name, u.Avatar, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}
