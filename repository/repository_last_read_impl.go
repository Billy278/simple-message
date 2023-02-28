package repository

import (
	"context"
	"database/sql"
	"simple-message/helper"
	"simple-message/model/domain"
)

type RepositoryLastReadImpl struct {
}

func NewRepositoryLastReadImpl() RepositoryLastRead {
	return &RepositoryLastReadImpl{}
}

func (repository_last *RepositoryLastReadImpl) Create(ctx context.Context, DB *sql.DB, last domain.LastRead) {
	sql := " insert into last_read(message_id,sender_id) values (?,?) On duplicate key update message_id=?"
	_, err := DB.ExecContext(ctx, sql, last.Message_id, last.Sender_id, last.Message_id)
	helper.PanicIfError(err)
}

func (repository_last *RepositoryLastReadImpl) FindMaxByid(ctx context.Context, DB *sql.DB) int {
	sql := "Select max(message_id) from last_read "
	rows, _ := DB.QueryContext(ctx, sql)
	var maxId int
	if rows.Next() {
		rows.Scan(&maxId)
	}
	return maxId

}
