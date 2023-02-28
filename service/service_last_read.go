package service

import (
	"context"
)

type ServiceLastRead interface {
	Create(ctx context.Context, id_last int, sender_id int)
	FindMaxByid(ctx context.Context) int
}
