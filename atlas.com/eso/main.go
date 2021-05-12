package main

import (
	"atlas-eso/database"
	"atlas-eso/logger"
	"atlas-eso/rest"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := logger.CreateLogger()

	db := database.ConnectToDatabase(l)

	createRestService(l, db)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
}

func createRestService(l *logrus.Logger, db *gorm.DB) {
	rs := rest.NewServer(l, db)
	go rs.Run()
}
