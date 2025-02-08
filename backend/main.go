package main

import (
	"fmt"
	"project_vk/backend/routes"
)

func main() {
	
	r := routes.SetupRouter()

	port := "8080"
	fmt.Println("Сервер запущен на порту " + port)
	r.Run(":" + port)
}
