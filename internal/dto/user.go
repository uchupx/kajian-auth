package dto

import (
	"time"

	"github.com/uchupx/kajian-auth/internal/repo/model"
)

type User struct {
	ID          string    `json:"id"`
	Password    string    `json:"-"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	ClientAppId string    `json:"client_app_id"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

func (u *User) Model(p *model.User) {
	u.ID = p.ID.String
	u.Username = p.Username.String
	u.ClientAppId = p.ClientAppID.String
	u.Email = p.Email.String
	u.Created = p.CreatedAt.Time
	u.Updated = p.UpdatedAt.Time
}

func (u *User) ToModel() model.User {
	var m model.User

	m.ID.String = u.ID
	m.ClientAppID.String = u.ClientAppId
	m.Username.String = u.Username
	m.Password.String = u.Password
	m.Email.String = u.Email
	m.CreatedAt.Time = u.Created
	m.UpdatedAt.Time = u.Updated

	return m
}
