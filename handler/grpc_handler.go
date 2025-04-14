package handler

import (
	"context"
	"log"

	"ADV3/proto/orderpb"
	"ADV3/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCOrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	Usecase *usecase.OrderUsecase
}

func NewGRPCOrderServer(uc *usecase.OrderUsecase) *GRPCOrderServer {
	return &GRPCOrderServer{Usecase: uc}
}

func (s *GRPCOrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	log.Printf("[gRPC] CreateOrder called with user_id=%s", req.Order.UserId)
	order := ProtoToModel(req.GetOrder())
	id, err := s.Usecase.CreateOrder(order)
	if err != nil {
		log.Printf("[gRPC][ERROR] Failed to create order: %v", err)
		return nil, status.Errorf(codes.Internal, "CreateOrder error: %v", err)
	}
	log.Printf("[gRPC] Order created with id=%s", id)
	return &orderpb.CreateOrderResponse{OrderId: id}, nil
}

func (s *GRPCOrderServer) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	log.Printf("[gRPC] GetOrder called with id=%s", req.GetId())
	order, err := s.Usecase.GetOrderByID(req.GetId())
	if err != nil {
		log.Printf("[gRPC][ERROR] Order not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "Order not found: %v", err)
	}
	return &orderpb.GetOrderResponse{Order: ModelToProto(order)}, nil
}

func (s *GRPCOrderServer) GetAllOrders(ctx context.Context, _ *orderpb.GetAllOrdersRequest) (*orderpb.GetAllOrdersResponse, error) {
	log.Printf("[gRPC] GetAllOrders called")
	orders, err := s.Usecase.GetAllOrders()
	if err != nil {
		log.Printf("[gRPC][ERROR] Failed to get all orders: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to get orders: %v", err)
	}

	var protoOrders []*orderpb.Order
	for _, o := range orders {
		cp := o
		protoOrders = append(protoOrders, ModelToProto(&cp))
	}
	return &orderpb.GetAllOrdersResponse{Orders: protoOrders}, nil
}

func (s *GRPCOrderServer) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*emptypb.Empty, error) {
	log.Printf("[gRPC] UpdateOrderStatus called with id=%s status=%s", req.GetId(), req.GetStatus())
	if err := s.Usecase.UpdateStatus(req.GetId(), req.GetStatus()); err != nil {
		log.Printf("[gRPC][ERROR] Failed to update order status: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update status: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *GRPCOrderServer) DeleteOrder(ctx context.Context, req *orderpb.DeleteOrderRequest) (*emptypb.Empty, error) {
	log.Printf("[gRPC] DeleteOrder called with id=%s", req.GetId())
	if err := s.Usecase.DeleteOrder(req.GetId()); err != nil {
		log.Printf("[gRPC][ERROR] Failed to delete order: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete order: %v", err)
	}
	return &emptypb.Empty{}, nil
}
