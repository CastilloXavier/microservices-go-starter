package domain

import (
	"context"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID        primitive.ObjectID
	UserId    string
	Status    string
	RiderFare *RiderFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, RiderFare *RiderFareModel) (*TripModel, error)
	GetRoute(ct context.Context, pickup, destination *types.Coordinate, useORMApi bool) (*types.OsrmApiResponse, error)
}
