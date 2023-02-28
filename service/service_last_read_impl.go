package service

import (
	"context"
	"database/sql"
	"simple-message/model/domain"
	"simple-message/repository"
)

type ServiceLastReadImpl struct {
	DB                 *sql.DB
	Repositorylastread repository.RepositoryLastRead
}

func NewServiceLastReadImpl(db *sql.DB, repository repository.RepositoryLastRead) ServiceLastRead {
	return &ServiceLastReadImpl{
		DB:                 db,
		Repositorylastread: repository,
	}
}

func (service_last *ServiceLastReadImpl) Create(ctx context.Context, id_last int, sender_id int) {
	db := service_last.DB
	last := domain.LastRead{
		Message_id: id_last,
		Sender_id:  sender_id,
	}
	service_last.Repositorylastread.Create(ctx, db, last)
}

func (service_last *ServiceLastReadImpl) FindMaxByid(ctx context.Context) int {
	db := service_last.DB
	id := service_last.Repositorylastread.FindMaxByid(ctx, db)
	return id
}
