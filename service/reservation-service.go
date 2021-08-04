package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"study_car_sharing/store"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationService struct {
	DB               *mongo.Database
	ReservationStore store.IReservationStore
	Ctx              *context.Context
}

func (r *ReservationService) FindAllReservation(w http.ResponseWriter, request *http.Request) {
	param := mux.Vars(request)
	offset, _ := strconv.Atoi(param["offset"])
	limit, _ := strconv.Atoi(param["limit"])
	response, err := r.ReservationStore.FindAllReservation(offset, limit)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, "error")
	}
	respondWithJson(w, http.StatusOK, response)
}

func (r *ReservationService) FindReservationByID(w http.ResponseWriter, request *http.Request) {
	param := mux.Vars(request)
	id := param["id"]
	response, err := r.ReservationStore.FindReservationByID(id)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest, "error")
	}
	respondWithJson(w, http.StatusOK, response)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
