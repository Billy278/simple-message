package service

import (
	"context"
	"simple-message/model/domain"
	"simple-message/model/web"
)

type ServiceUser interface {
	Register(ctx context.Context, request web.RequestUser)
	FindByUser(ctx context.Context, username string) (domain.User, error)
}
