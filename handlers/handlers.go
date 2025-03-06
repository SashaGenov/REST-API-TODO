package handlers

import (
        "context"
        "strconv"
        "time"

        "github.com/gofiber/fiber/v2"
        "todo-api/database"
        "todo-api/models"
)

// CreateTask - обработчик для создания задачи.
func CreateTask(c *fiber.Ctx) error {
        task := new(models.Task)
        // Парсим JSON в структуру Task.
        if err := c.BodyParser(task); err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        //Выполняем SQL-запрос для создания задачи.
        _, err := database.Conn.Exec(context.Background(),
                "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)",
                task.Title, task.Description, task.Status)
        if err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.SendStatus(fiber.StatusCreated)
}

// GetTasks - обработчик для получения списка задач.
func GetTasks(c *fiber.Ctx) error {
        //Выполняем SQL-запрос для получения списка задач.
        rows, err := database.Conn.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
        if err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        defer rows.Close()

        tasks := []models.Task{}
        // Итерация по строкам результата запроса.
        for rows.Next() {
                task := models.Task{}
                err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
                if err != nil {
                        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
                }
                tasks = append(tasks, task)
        }

        return c.JSON(tasks)
}

// UpdateTask - обработчик для обновления задачи.
func UpdateTask(c *fiber.Ctx) error {
        // Получаем ID задачи из параметров запроса.
        id, err := strconv.Atoi(c.Params("id"))
        if err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
        }

        task := new(models.Task)
        if err := c.BodyParser(task); err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        //Выполняем SQL-запрос для обновления задачи.
        _, err = database.Conn.Exec(context.Background(),
                "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5",
                task.Title, task.Description, task.Status, time.Now(), id)
        if err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.SendStatus(fiber.StatusOK)
}

// DeleteTask - обработчик для удаления задачи.
func DeleteTask(c *fiber.Ctx) error {
        id, err := strconv.Atoi(c.Params("id"))
        if err != nil {
                return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
        }

        //Выполняем SQL-запрос для удаления задачи.
        _, err = database.Conn.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
        if err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.SendStatus(fiber.StatusOK)
}