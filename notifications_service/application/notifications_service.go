package application

import (
	"github.com/mihajlo-ra92/XML/notifications_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationsService struct {
	store domain.NotificationsStore
}

func NewNotificationsService(store domain.NotificationsStore) *NotificationsService {
	return &NotificationsService{
		store: store,
	}
}

func (service *NotificationsService) Get(id primitive.ObjectID) (*domain.Notification, error) {
	return service.store.Get(id)
}

func (service *NotificationsService) GetByUser(userId string) (*domain.Notification, error) {
	return service.store.GetByUser(userId)
}

func (service *NotificationsService) GetAll() ([]*domain.Notification, error) {
	return service.store.GetAll()
}

func (service *NotificationsService) Delete(notificationId string) error {
	id, _ := primitive.ObjectIDFromHex(notificationId)
	rating := domain.Notification{Id: id}
	return service.store.Delete(&rating)
}

func (service *NotificationsService) Create(notification *domain.Notification) error {
	return service.store.Insert(notification)
}
