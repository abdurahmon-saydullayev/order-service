package postgres

import (
	"Projects/store/order-service/genproto/order_service"
	"Projects/store/order-service/storage"
	"context"
	"database/sql"
	"fmt"

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

	query := `insert into "order"
				(id, 
				product_id, 
				user_id,
				status
				) VALUES (
					$1, 
					$2, 
					$3,
					$4
				)`

	_, err = b.db.Exec(ctx, query,
		id,
		req.ProductId,
		req.UserId,
		req.Status,
	)
	fmt.Println(query)
	if err != nil {
		return resp, err
	}

	resp = &order_service.OrderPrimaryKey{
		Id: id,
	}

	return resp, nil
}

func (b *orderRepo) GetById(ctx context.Context, req *order_service.OrderPrimaryKey) (resp *order_service.Order, err error) {
	fmt.Println("laskdmalksdmalskdmasklcaskmdc")

	query := `
		SELECT id,product_id,user_id,status
	     FROM "order"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		product_id sql.NullString
		user_id    sql.NullString
		status     sql.NullString
	)
	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&product_id,
		&user_id,
		&status,
	)
	if err != nil {
		fmt.Println("this aksdjnaksjndcaskncdkasnc")
		return nil, err
	}

	resp = &order_service.Order{
		Id:        id.String,
		ProductId: product_id.String,
		UserId:    user_id.String,
		Status:    status.String,
	}

	return
}
