package usecase

import (
	"context"
	"fmt"

	"ADV3/proto/gen"
	"ADV3/repository"
)

type OrderUsecase struct {
	repo           *repository.OrderRepository
	productService gen.ProductServiceClient
}

func NewOrderUsecase(repo *repository.OrderRepository, productClient gen.ProductServiceClient) *OrderUsecase {
	return &OrderUsecase{
		repo:           repo,
		productService: productClient,
	}
}

func (u *OrderUsecase) CreateOrder(order *repository.Order) (string, error) {
	ctx := context.Background()

	// Validate each product
	for _, item := range order.Items {
		resp, err := u.productService.GetProduct(ctx, &gen.GetProductRequest{Id: item.ProductID})
		if err != nil {
			return "", fmt.Errorf("product %s not found: %v", item.ProductID, err)
		}
		if resp.Stock < int32(item.Quantity) {
			return "", fmt.Errorf("product %s out of stock (have %d, need %d)", item.ProductID, resp.Stock, item.Quantity)
		}
	}

	return u.repo.Create(order)
}

func (u *OrderUsecase) GetAllOrders() ([]repository.Order, error) {
	return u.repo.GetAll()
}

func (u *OrderUsecase) GetOrderByID(id string) (*repository.Order, error) {
	return u.repo.GetByID(id)
}

func (u *OrderUsecase) UpdateStatus(id string, status string) error {
	return u.repo.UpdateStatus(id, status)
}

func (u *OrderUsecase) DeleteOrder(id string) error {
	return u.repo.Delete(id)
}
