package application

import (
	"github.com/mihajlo-ra92/XML/rating_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingService struct {
	store                      domain.RatingStore
	accommodationClientAddress string
	bookingClientAddress       string
	userClientAddress          string
}

func NewRatingService(store domain.RatingStore, accommodationClientAddress string, bookingClientAddress string, userClientAddress string) *RatingService {
	return &RatingService{
		store:                      store,
		accommodationClientAddress: accommodationClientAddress,
		bookingClientAddress:       bookingClientAddress,
		userClientAddress:          userClientAddress,
	}
}

func (service *RatingService) Get(id primitive.ObjectID) (*domain.Rating, error) {
	return service.store.Get(id)
}

func (service *RatingService) GetAll() ([]*domain.Rating, error) {
	return service.store.GetAll()
}

func (service *RatingService) Create(user *domain.Rating) error {

	//TODO: Optional
	err := service.store.Insert(user)
	if err != nil {
		return err
	}
	return nil
}
