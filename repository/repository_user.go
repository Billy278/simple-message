package repository

import (
	"context"
	"database/sql"
	"simple-message/model/domain"
)

type RepositoryUser interface {
	Register(ctx context.Context, DB *sql.DB, user domain.User)
	FindByUser(ctx context.Context, DB *sql.DB, username string) (domain.User, error)
}
