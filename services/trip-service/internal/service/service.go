package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripService struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *TripService {
	return &TripService{
		repo: repo,
	}
}

func (s *TripService) CreateTrip(ctx context.Context, fare *domain.RiderFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:        primitive.NewObjectID(),
		UserId:    fare.UserID,
		Status:    "pending",
		RiderFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)
}
