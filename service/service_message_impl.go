package service

import (
	"context"
	"database/sql"
	"simple-message/helper"
	"simple-message/model/domain"
	"simple-message/model/web"
	"simple-message/repository"

	"github.com/go-playground/validator/v10"
)

type ServiceMessageImpl struct {
	DB                *sql.DB
	Validate          *validator.Validate
	RepositoryMessage repository.RepositoryMessage
}

func NewServiceMessageImpl(db *sql.DB, validate *validator.Validate, repository repository.RepositoryMessage) ServiceMessage {
	return &ServiceMessageImpl{
		DB:                db,
		Validate:          validate,
		RepositoryMessage: repository,
	}
}
func (service_message *ServiceMessageImpl) Create(ctx context.Context, request web.RequestMessage) web.ResponseMessage {
	err := service_message.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service_message.DB
	message := domain.Massage{
		Sender:    request.Sender,
		Date:      request.Date,
		Message:   request.Message,
		Receiver:  request.Receiver,
		Sender_id: request.Sender_id,
	}
	message = service_message.RepositoryMessage.Create(ctx, db, message)

	return helper.ResponseMessageSucces(message)

}
func (service_message *ServiceMessageImpl) CreateTest(ctx context.Context, request web.RequestMessage) web.ResponseMessage {
	err := service_message.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service_message.DB
	message := domain.Massage{
		Sender:    request.Sender,
		Date:      request.Date,
		Message:   request.Message,
		Receiver:  request.Receiver,
		Sender_id: request.Sender_id,
	}
	message = service_message.RepositoryMessage.CreateTest(ctx, db, message)

	return helper.ResponseMessageSucces(message)
}
func (service_message *ServiceMessageImpl) SelectPartSender(ctx context.Context, sender string, receiver string, offside int, limit int) []web.ResponseSenderMessage {
	db := service_message.DB
	message := service_message.RepositoryMessage.SelectPartSender(ctx, db, sender, receiver, offside, limit)
	return helper.ResponsesSenderMessage(message)
}
func (service_message *ServiceMessageImpl) SelectAllSenderWithLastMessage(ctx context.Context, receiver string) []web.ResponseMessage {
	db := service_message.DB
	message := service_message.RepositoryMessage.SelectAllSenderWithLastMessage(ctx, db, receiver)
	return helper.ResponsesMessage(message)

}
