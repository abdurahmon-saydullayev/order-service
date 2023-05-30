package grpc

import (
	"Projects/store/order-service/config"
	"Projects/store/order-service/genproto/order_service"
	"Projects/store/order-service/grpc/client"
	"Projects/store/order-service/grpc/service"
	"Projects/store/order-service/pkg/logger"
	"Projects/store/order-service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	order_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
