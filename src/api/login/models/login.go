package models

import (
	"database/sql"
	"grpc_blog/src/models"
	"time"
)

type Users struct {
	ID              int64
	Name            string
	Avatar          string
	Email           string
	EmailVerifiedAt sql.NullInt64
	Password        string
	RememberToken   sql.NullString
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func GetUsersByEmail(email string) (*Users, error) {
	u := Users{}
	stms, err := models.DB.Prepare("select * from users where email=?")
	if err != nil {
		return nil, err
	}
	defer stms.Close()
	row := stms.QueryRow(email)
	err = row.Scan(&u.ID, &u.Name, &u.Avatar, &u.Email, &u.EmailVerifiedAt, &u.Password, &u.RememberToken, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, err
}
