package usecase

import "ADV3/repository"

type OrderUsecase struct {
	repo *repository.OrderRepository
}

func NewOrderUsecase(r *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: r}
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

func (u *OrderUsecase) CreateOrder(order *repository.Order) (string, error) {
	return u.repo.Create(order)
}
