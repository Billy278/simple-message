package repository

import (
	"context"
	"database/sql"
	"simple-message/model/domain"
)

type RepositoryMessage interface {
	Create(ctx context.Context, DB *sql.DB, message domain.Massage) domain.Massage
	SelectPartSender(ctx context.Context, DB *sql.DB, sender string, receiver string, offside int, limit int) []domain.Massage
	SelectAllSenderWithLastMessage(ctx context.Context, DB *sql.DB) []domain.Massage
}
