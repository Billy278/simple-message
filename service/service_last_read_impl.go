package service

import (
	"context"
	"database/sql"
	"errors"
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

// func (service_last *ServiceLastReadImpl) CreateNew(ctx context.Context, message_id int, sender_id int) {
// 	db := service_last.DB
// 	last := domain.LastRead{
// 		Message_id: message_id,
// 		Sender_id:  sender_id,
// 	}
// 	service_last.Repositorylastread.CreateNew(ctx, db, last)
// }

//	func (service_last *ServiceLastReadImpl) FindMaxByid(ctx context.Context) int {
//		db := service_last.DB
//		id := service_last.Repositorylastread.FindMaxByid(ctx, db)
//		return id
//	}
func (service_last *ServiceLastReadImpl) FindByid(ctx context.Context, id int) (int, error) {
	db := service_last.DB
	id, err := service_last.Repositorylastread.FindByid(ctx, db, id)
	if err != nil {
		return id, errors.New("NOT FOUND")
	} else {
		return id, nil
	}
}
