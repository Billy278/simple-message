package repository

import (
	"context"
	"database/sql"
	"simple-message/model/domain"
)

type RepositoryLastRead interface {
	Create(ctx context.Context, DB *sql.DB, last domain.LastRead)
	//CreateNew(ctx context.Context, DB *sql.DB, last domain.LastRead)
	//FindMaxByid(ctx context.Context, DB *sql.DB) int
	FindByid(ctx context.Context, DB *sql.DB, id int) (int, error)
}
