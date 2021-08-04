package api

import (
	"context"
	"study_car_sharing/service"
	"study_car_sharing/store"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServerRunner struct {
	DB *mongo.Database
}

func (s *ServerRunner) Route() *mux.Router {
	ctx := context.Background()
	reservationService := service.ReservationService{
		DB: s.DB,
		ReservationStore: &store.ReservationStore{
			Ctx: &ctx,
			DB:  s.DB,
		},
		Ctx: &ctx,
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/reservations/{id}", reservationService.FindReservationByID).Methods("GET")
	router.HandleFunc("/api/reservations", reservationService.FindReservationByID).Methods("GET").
		Queries("limit", "{limit}", "offset", "{offset}")
	router.HandleFunc("/api/reservations/{id}", reservationService.FindReservationByID).Methods("POST")
	router.HandleFunc("/api/reservations/{id}", reservationService.FindReservationByID).Methods("UPDATE")
	router.HandleFunc("/api/reservations/{id}", reservationService.FindReservationByID).Methods("DELETE")
	return router
}
