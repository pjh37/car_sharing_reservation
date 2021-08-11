package store

import (
	"context"
	"study_car_sharing/store/datamodel"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationStore struct {
	DB  *mongo.Database
	Ctx *context.Context
}

func (r *ReservationStore) FindReservationByID(rid string) (datamodel.Reservation, error) {

	return datamodel.Reservation{
		UserID: "1234",
	}, nil
}

func (r *ReservationStore) FindAllReservation(offset, limit int) ([]datamodel.Reservation, error) {
	return []datamodel.Reservation{}, nil
}
func (r *ReservationStore) SaveReservation() error {
	return nil
}
