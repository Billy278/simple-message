package service

import (
	"context"
	"simple-message/model/web"
)

type ServiceMessage interface {
	Create(ctx context.Context, request web.RequestMessage) web.ResponseMessage
	SelectPartSender(ctx context.Context, sender string, receiver string, offside int, limit int) []web.ResponseSenderMessage
	SelectAllSenderWithLastMessage(ctx context.Context) []web.ResponseMessage
}
