package api

import (
	pb "github.com/mihajlo-ra92/XML/common/proto/notifications_service"
	"github.com/mihajlo-ra92/XML/notifications_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapNewNotification(request *pb.CreateNotificationRequest) *domain.Notification {
	booking := &domain.Notification{
		Id:                    primitive.NewObjectID(),
		UserId:                request.Notification.UserId,
		HostRequest:           request.Notification.HostRequest,
		HostCancelation:       request.Notification.HostCancelation,
		HostRate:              request.Notification.HostRate,
		HostAccommodationRate: request.Notification.HostAccommodationRate,
		HostOutstanding:       request.Notification.HostOutstanding,
		GuestRequest:          request.Notification.GuestRequest,
	}
	return booking
}

func mapNotification(rating *domain.Notification) *pb.Notification {
	ratingPb := &pb.Notification{
		Id:                    rating.Id.Hex(),
		UserId:                rating.UserId,
		HostRequest:           rating.HostRequest,
		HostCancelation:       rating.HostCancelation,
		HostRate:              rating.HostRate,
		HostAccommodationRate: rating.HostAccommodationRate,
		HostOutstanding:       rating.HostOutstanding,
		GuestRequest:          rating.GuestRequest,
	}
	return ratingPb
}
