package handler

import (
	"ADV3/proto/orderpb"
	"ADV3/repository"
)

func ProtoToModel(o *orderpb.Order) *repository.Order {
	items := make([]repository.OrderItem, len(o.GetItems()))
	for i, it := range o.GetItems() {
		items[i] = repository.OrderItem{
			ProductID:   it.GetProductId(),
			ProductName: it.GetProductName(),
			Quantity:    int(it.GetQuantity()),
			Price:       int(it.GetPrice()),
			ImageURL:    it.GetImageUrl(),
		}
	}
	return &repository.Order{
		UserID: o.GetUserId(),
		Status: o.GetStatus(),
		Items:  items,
		// CreatedAt will be set in grpc_handler, ignore here
	}
}

func ModelToProto(o *repository.Order) *orderpb.Order {
	items := make([]*orderpb.OrderItem, len(o.Items))
	for i, it := range o.Items {
		items[i] = &orderpb.OrderItem{
			ProductId:   it.ProductID,
			ProductName: it.ProductName,
			Quantity:    int32(it.Quantity),
			Price:       int32(it.Price),
			ImageUrl:    it.ImageURL,
		}
	}
	return &orderpb.Order{
		Id:        o.ID.Hex(),
		UserId:    o.UserID,
		Status:    o.Status,
		Items:     items,
		CreatedAt: o.CreatedAt,
	}
}
