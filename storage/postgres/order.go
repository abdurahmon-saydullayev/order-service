package postgres

import (
	"Projects/store/order-service/genproto/order_service"
	"Projects/store/order-service/storage"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.OrderRepoI {
	return &orderRepo{
		db: db,
	}
}

func (b *orderRepo) Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.OrderPrimaryKey, err error) {
	id := uuid.New().String()

	query := `insert into orders 
				(id, 
				product_id, 
				user_id,
				user_first_name,
				user_last_name,
				user_phone_number,
				status
				) VALUES (
					$1, 
					$2, 
					$3,
					$4,
					$5,
					$6,
					$7
				)`

	_, err = b.db.Exec(ctx, query,
		id,
		req.ProductId,
		req.UserData.Id,
		req.UserData.FirstName,
		req.UserData.LastName,
		req.UserData.PhoneNumber,
		req.Status,
	)

	if err != nil {
		return resp, err
	}

	resp = &order_service.OrderPrimaryKey{
		Id: req.Id,
	}

	return resp, nil
}

func (b *orderRepo) GetById(ctx context.Context, req *order_service.OrderPrimaryKey) (*order_service.Order, error) {
	var (
		query string
		order order_service.Order
	)

	query = `
		SELECT * FROM orders
		WHERE id = $1
	`

	err := b.db.QueryRow(ctx, query, req.Id).Scan(
		&order.Id,
		&order.ProductId,
		&order.UserData.Id,
		&order.UserData.FirstName,
		&order.UserData.LastName,
		&order.UserData.PhoneNumber,
		&order.Status,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
