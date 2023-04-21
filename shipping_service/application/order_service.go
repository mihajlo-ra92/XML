package application

import (
	"github.com/mihajlo-ra92/XML/shipping_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	store domain.OrderStore
}

func NewOrderService(store domain.OrderStore) *OrderService {
	return &OrderService{
		store: store,
	}
}

func (service *OrderService) Get(id primitive.ObjectID) (*domain.Order, error) {
	return service.store.Get(id)
}

func (service *OrderService) GetAll() ([]*domain.Order, error) {
	return service.store.GetAll()
}

func (service *OrderService) Create(order *domain.Order) error {
	order.Status = domain.Scheduled
	return service.store.Insert(order)
}

func (service *OrderService) Cancel(order *domain.Order) error {
	order.Status = domain.Cancelled
	return service.store.UpdateStatus(order)
}
