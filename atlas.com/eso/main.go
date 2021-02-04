package main

import (
	"atlas-eso/database/equipment"
	"atlas-eso/rest"
	"atlas-eso/retry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func connectToDatabase(attempt int) (bool, interface{}, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:the@tcp(atlas-db:3306)/atlas-eso?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return true, nil, err
	}
	return false, db, err
}

func main() {
	l := log.New(os.Stdout, "eso ", log.LstdFlags|log.Lmicroseconds)

	r, err := retry.RetryResponse(connectToDatabase, 10)
	if err != nil {
		panic("failed to connect database")
	}
	db := r.(*gorm.DB)

	// Migrate the schema
	equipment.Migration(db)

	createRestService(l, db)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Println("[INFO] shutting down via signal:", sig)
}

func createRestService(l *log.Logger, db *gorm.DB) {
	rs := rest.NewServer(l, db)
	go rs.Run()
}
