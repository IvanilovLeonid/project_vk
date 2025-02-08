package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PingRequest struct {
	IPAddress string `json:"ip_address"`
	Success   bool   `json:"success"`
}

type Container struct {
	ID                 uint      `json:"id"`
	IPAddress          string    `json:"ip_address"`
	LastPingTime       time.Time `json:"last_ping_time"`
	LastSuccessfulPing time.Time `json:"last_successful_ping"`
}

// Получение всех контейнеров
func GetContainers(c *gin.Context) {
	resp, err := http.Get("http://database-service:8080/containers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных из базы"})
		return
	}
	defer resp.Body.Close()

	var containers []Container
	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка декодирования данных"})
		return
	}

	c.JSON(http.StatusOK, containers)
}

// Пинг контейнера
func PingContainer(c *gin.Context) {
	var req PingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}

	// Формируем запрос для микросервиса базы данных
	body, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сериализации данных"})
		return
	}

	resp, err := http.Post("http://database-service:8080/ping", "application/json", bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки данных в базу"})
		return
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка декодирования ответа"})
		return
	}

	c.JSON(http.StatusOK, response)
}
