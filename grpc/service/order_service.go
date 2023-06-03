package service

import (
	"Projects/store/order-service/config"
	"Projects/store/order-service/genproto/order_service"
	"Projects/store/order-service/genproto/user_service"
	"Projects/store/order-service/grpc/client"
	"Projects/store/order-service/pkg/logger"
	"Projects/store/order-service/storage"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *OrderService {
	return &OrderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *OrderService) Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.Order, err error) {
	s.log.Info("CreateOrder", logger.Any("request", req))

	orderPKey, err := s.strg.Order().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateOrder", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	createdOrder, err := s.strg.Order().GetById(ctx, orderPKey)
	if err != nil {
		s.log.Error("GetOrder", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return createdOrder, nil
}
func (s *OrderService) Delete(context.Context, *order_service.OrderPrimaryKey) (res *emptypb.Empty, err error) {
	return res, nil
}

func (s *OrderService) GetAll(context.Context, *order_service.GetOrderListRequest) (res *order_service.GetOrderListResponse, err error) {
	return res, nil
}

func (s *OrderService) GetById(ctx context.Context, req *order_service.OrderPrimaryKey) (*order_service.Order, error) {
	s.log.Info("GetByID order", logger.Any("request", req))

	order, err := s.strg.Order().GetById(ctx, req)
	if err != nil {
		s.log.Error("GetByID", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.services.UserService().GetByID(ctx, &user_service.UserPrimaryKey{Id: order.UserId})
	if err != nil {
		return nil, nil
	}

	order.Id = user.Id
	order.UserFirstName = user.FirstName
	order.UserLastName = user.LastName
	order.UserPhoneNumber = user.PhoneNumber

	return order, nil
}
