package storage

import (
	"Projects/store/order-service/genproto/order_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Order() OrderRepoI
}

type OrderRepoI interface {
	Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.OrderPrimaryKey, err error)
	GetById(ctx context.Context, req *order_service.OrderPrimaryKey) (*order_service.Order, error)
}
