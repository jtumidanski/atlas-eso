package rest

import (
	"atlas-eso/equipment"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	l  *logrus.Logger
	hs *http.Server
}

func NewServer(l *logrus.Logger, db *gorm.DB) *Server {
	router := mux.NewRouter().PathPrefix("/ms/eso").Subrouter()
	router.Use(commonHeader)

	eRouter := router.PathPrefix("/equipment").Subrouter()
	eRouter.HandleFunc("", equipment.HandleCreateEquipment(l, db)).Methods(http.MethodPost)
	eRouter.HandleFunc("/", equipment.HandleCreateEquipment(l, db)).Methods(http.MethodPost)
	eRouter.HandleFunc("/{equipmentId}", equipment.HandleGetEquipmentById(l, db)).Methods(http.MethodGet)

	w := l.Writer()
	defer w.Close()

	hs := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     log.New(w, "", 0), // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	return &Server{l, &hs}
}

func (s *Server) Run() {
	s.l.Infof("Starting server on port 8080")
	err := s.hs.ListenAndServe()
	if err != nil {
		s.l.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func commonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
