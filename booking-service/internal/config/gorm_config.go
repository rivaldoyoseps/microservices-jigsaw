package config

import (
	"log"
	"time"

	"github.com/rivaldoyoseps/booking-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormConfig() *gorm.DB {
	dsn := "host=localhost user=postgres password=halakhitado00 dbname=booking_service_db port=5432 sslmode=disable"

	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,   // Query > 1 detik dianggap 
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database booking: %v", err)
	}

	// AutoMigrate tabel Booking
	if err := db.AutoMigrate(&domain.Booking{}); err != nil {
		log.Fatalf("Gagal migrasi tabel Slot: %v", err)
	}

	log.Println("Berhasil terhubung ke database dan migrasi tabel Slot")
	return db
}
