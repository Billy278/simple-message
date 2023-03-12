package service

import (
	"context"
)

type ServiceLastRead interface {
	Create(ctx context.Context, id_last int, sender_id int)
	//CreateNew(ctx context.Context, message_id int, sender_id int)
	//FindMaxByid(ctx context.Context) int
	FindByid(ctx context.Context, id int) (int, error)
}
