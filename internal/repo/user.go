package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/kajian-auth/internal/repo/model"
)

const (
	findUserByUsernameEmailQuery = "SELECT * FROM users WHERE username = ? OR email = ?"
)

type UserRepo struct {
	db *sqlx.DB
}

func (r *UserRepo) FindUserByUsernameEmail(ctx context.Context, val string) (*model.User, error) {
	var user model.User

	stmt, err := r.db.PreparexContext(ctx, findUserByUsernameEmailQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	row := stmt.QueryRowxContext(ctx, val, val)
	err = row.StructScan(&user)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &user, nil
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}
