package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Container struct {
	ID                 uint      `gorm:"primaryKey"` // Уникальный ID
	IPAddress          string    `gorm:"unique;not null"`
	LastPingTime       time.Time `gorm:"not null"`
	LastSuccessfulPing time.Time
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Container{}); err != nil {
		log.Println("Ошибка миграции:", err)
		return err
	}
	log.Println("Миграция выполнена успешно.")
	return nil
}
