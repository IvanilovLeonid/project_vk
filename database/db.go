package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Модель контейнера
type Container struct {
	ID                 uint      `gorm:"primaryKey"`
	IPAddress          string    `gorm:"unique;not null"`
	LastPingTime       time.Time `gorm:"not null"`
	LastSuccessfulPing time.Time
}

// Структура для подключения к базе данных
var DB *gorm.DB

// Инициализация базы данных
func initDB() {
	// Строка подключения для PostgreSQL
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=container_monitoring port=5432 sslmode=disable"
	var err error

	// Подключаемся к базе данных
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Выполняем миграцию (создаем таблицы)
	if err := DB.AutoMigrate(&Container{}); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}
	log.Println("База данных и миграция успешно инициализированы")
}

// Обработчик для получения всех контейнеров
func GetContainers(c *gin.Context) {
	var containers []Container
	if err := DB.Find(&containers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}
	c.JSON(http.StatusOK, containers)
}

// Обработчик для создания или обновления контейнера
func PingContainer(c *gin.Context) {
	var req struct {
		IPAddress string `json:"ip_address"`
		Success   bool   `json:"success"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}

	var container Container
	if err := DB.Where("ip_address = ?", req.IPAddress).First(&container).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			container = Container{
				IPAddress:    req.IPAddress,
				LastPingTime: time.Now(),
			}
			if req.Success {
				container.LastSuccessfulPing = time.Now()
			}
			if err := DB.Create(&container).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания нового контейнера"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения контейнера"})
			return
		}
	} else {
		container.LastPingTime = time.Now()
		if req.Success {
			container.LastSuccessfulPing = time.Now()
		}
		if err := DB.Save(&container).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления контейнера"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ping received"})
}

func main() {
	// Инициализируем базу данных
	initDB()

	// Инициализируем роутер
	r := gin.Default()

	// Определяем маршруты
	r.GET("/containers", GetContainers)
	r.POST("/ping", PingContainer)

	// Запуск сервиса
	r.Run(":8081") // HTTP сервер на порту 8081 для Database Service
}
