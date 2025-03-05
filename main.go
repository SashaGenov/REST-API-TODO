package main

import (
	"context"
	"fmt"
	"log"

	"todo-api/database"
	"todo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Подключаемся к базе данных.
	database.ConnectDB()
	defer func() {
		if database.Conn != nil {
			err := database.Conn.Close(context.Background())
			if err != nil {
				log.Printf("Error closing database connection: %v", err)
			}
		}
	}()

	// Создание экземпляра Fiber.
	app := fiber.New()

	// Определение маршрутов API.
	app.Post("/tasks", handlers.CreateTask)       // Маршрут для создания задачи.
	app.Get("/tasks", handlers.GetTasks)          // Маршрут для получения списка задач.
	app.Put("/tasks/:id", handlers.UpdateTask)    // Маршрут для обновления задачи.
	app.Delete("/tasks/:id", handlers.DeleteTask) // Маршрут для удаления задачи.

	// Запуск сервера.
	port := 8080
	log.Printf("Server listening on port %d", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
