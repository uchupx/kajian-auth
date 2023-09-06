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
	insertUserQuery              = "INSERT INTO users(username, password, email) VALUES (?, ?, ?)"
)

type UserRepo struct {
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

func (r *UserRepo) Insert(ctx context.Context, data model.User) (*int64, error) {
	stmt, err := r.db.FPreparexContext(ctx, insertUserQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row, err := stmt.FExecContext(ctx,
		data.Username,
		data.Password,
		data.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	lastId, err := row.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return &lastId, nil
}

func NewUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{db: db}
}
