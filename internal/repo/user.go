package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uchupx/kajian-api/pkg/db"
	"github.com/uchupx/kajian-auth/internal/repo/model"
)

const (
	findUserByUsernameEmailQuery = "SELECT * FROM users WHERE username = ? OR email = ?"
	insertUserQuery              = "INSERT INTO users(id, username, password, email, created_at) VALUES (?, ?, ?, ?, ?)"
)

type UserRepo struct {
	BaseRepo
	db *db.DB
}

func (r *UserRepo) FindUserByUsernameEmail(ctx context.Context, val string) (*model.User, error) {
	var user model.User

	stmt, err := r.db.FPreparexContext(ctx, findUserByUsernameEmailQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.FQueryRowxContext(ctx, val, val)
	err = row.StructScan(&user)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &user, nil
}

func (r *UserRepo) Insert(ctx context.Context, data model.User) (*string, error) {
	stmt, err := r.db.FPreparexContext(ctx, insertUserQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	id := r.ID()

	defer stmt.Close()

	_, err = stmt.FExecContext(ctx,
		id,
		data.Username.String,
		data.Password.String,
		data.Email.String,
		data.CreatedAt.Time,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return id, nil
}

func NewUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{db: db}
}
