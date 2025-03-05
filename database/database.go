package database

import (
        "context"
        "fmt"
        "log"
        "os"

        "github.com/jackc/pgx/v5"
        "github.com/joho/godotenv"
)

var Conn *pgx.Conn
// ConnectDB - функция для подключения к базе данных.

func ConnectDB() {
        err := godotenv.Load()
        if err != nil {
                log.Fatal("Error loading .env file")
        }

         // Формирование строки подключения к базе данных.
        connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
                os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

        // Установка соединения с базой данных.
        Conn, err = pgx.Connect(context.Background(), connStr)
        if err != nil {
                log.Fatalf("Unable to connect to database: %v\n", err)
        }

        fmt.Println("Connected to database")
}