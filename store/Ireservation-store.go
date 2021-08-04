package store

import "study_car_sharing/store/datamodel"

type IReservationStore interface {
	FindReservationByID(rid string) (datamodel.Reservation, error)
	FindAllReservation(offset, limit int) ([]datamodel.Reservation, error)
	// SaveReservation()
	// UpdateReservation()
	// DeleteReservation()
}
