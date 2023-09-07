package dto

import (
	"time"

	"github.com/uchupx/kajian-auth/internal/repo/model"
)

type User struct {
	ID       int64     `json:"id"`
	Password string    `json:"-"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func (u *User) Model(p *model.User) {
	u.ID = p.ID.Int64
	u.Username = p.Username.String
	u.Email = p.Email.String
	u.Created = p.CreatedAt.Time
	u.Updated = p.UpdatedAt.Time
}

func (u *User) ToModel() model.User {
	var m model.User

	m.ID.Int64 = u.ID
	m.Username.String = u.Username
	m.Password.String = u.Password
	m.Email.String = u.Email
	m.CreatedAt.Time = u.Created
	m.UpdatedAt.Time = u.Updated

	return m
}
