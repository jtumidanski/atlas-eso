package rest

import (
	"atlas-eso/equipment"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, db *gorm.DB, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes(db))
}

func ProduceRoutes(db *gorm.DB) func(l logrus.FieldLogger) http.Handler {
	return func(l logrus.FieldLogger) http.Handler {
		router := mux.NewRouter().PathPrefix("/ms/eso").Subrouter()
		router.Use(CommonHeader)

		eRouter := router.PathPrefix("/equipment").Subrouter()
		eRouter.HandleFunc("", equipment.HandleCreateRandomEquipment(l, db)).Queries("random", "{random}").Methods(http.MethodPost)
		eRouter.HandleFunc("", equipment.HandleCreateEquipment(l, db)).Methods(http.MethodPost)
		eRouter.HandleFunc("/{equipmentId}", equipment.HandleGetEquipmentById(l, db)).Methods(http.MethodGet)
		eRouter.HandleFunc("/{equipmentId}", equipment.HandleDeleteEquipment(l, db)).Methods(http.MethodDelete)

		return router
	}
}