package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Container struct {
	ID                 uint      `gorm:"primaryKey"`
	IPAddress          string    `gorm:"unique;not null"`
	LastPingTime       time.Time `gorm:"not null"`
	LastSuccessfulPing time.Time
}

var DB *gorm.DB

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("Переменная DATABASE_URL не установлена в .env файле")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	if err := DB.AutoMigrate(&Container{}); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}
	log.Println("База данных инициализация завершена")
}

func getContainers(c *gin.Context) {
	var containers []Container
	if err := DB.Find(&containers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}
	c.JSON(http.StatusOK, containers)
}

func pingContainer(c *gin.Context) {
	var req struct {
		IPAddress string `json:"ip_address"`
		Success   bool   `json:"success"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}

	var container Container
	err := DB.Where("ip_address = ?", req.IPAddress).First(&container).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			container = Container{
				IPAddress:          req.IPAddress,
				LastPingTime:       time.Now(),
				LastSuccessfulPing: time.Time{},
			}
			if req.Success {
				container.LastSuccessfulPing = time.Now()
			}
			if err := DB.Create(&container).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания контейнера"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных контейнера"})
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
	initDB()

	r := gin.Default()
	r.GET("/containers", getContainers)
	r.POST("/ping", pingContainer)

	r.Run(":8081")
}
