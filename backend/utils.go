package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("Request: %s %s took %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}

func deleteContainer(c *gin.Context) {
	ipAddress := c.Param("ip_address")
	var container Container
	err := DB.Where("ip_address = ?", ipAddress).First(&container).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Контейнер не найден"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения контейнера"})
		return
	}

	if err := DB.Delete(&container).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления контейнера"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Контейнер удален"})
}

func healthCheck(c *gin.Context) {
	containersCount := int64(0)
	DB.Model(&Container{}).Count(&containersCount)

	c.JSON(http.StatusOK, gin.H{
		"status":         "ok",
		"containers":     containersCount,
		"last_ping_time": time.Now().String(),
	})
}
