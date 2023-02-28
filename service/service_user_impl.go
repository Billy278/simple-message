package service

import (
	"context"
	"database/sql"
	"errors"
	"simple-message/helper"
	"simple-message/model/domain"
	"simple-message/model/web"
	"simple-message/repository"

	"github.com/go-playground/validator/v10"
)

type ServiceUserImpl struct {
	DB             *sql.DB
	Validate       *validator.Validate
	Repositoryuser repository.RepositoryUser
}

func NewServiceUserImpl(db *sql.DB, validate *validator.Validate, repository repository.RepositoryUser) ServiceUser {
	return &ServiceUserImpl{
		DB:             db,
		Validate:       validate,
		Repositoryuser: repository,
	}
}

func (service_user *ServiceUserImpl) Register(ctx context.Context, request web.RequestUser) {
	err := service_user.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service_user.DB
	user := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}
	service_user.Repositoryuser.Register(ctx, db, user)

}
func (service_user *ServiceUserImpl) FindByUser(ctx context.Context, username string) (domain.User, error) {
	db := service_user.DB
	user, err := service_user.Repositoryuser.FindByUser(ctx, db, username)
	if err != nil {
		return user, errors.New("NOT FOUND")
	} else {
		return user, nil
	}
}
