package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	repository "learn-back/internal/repository/postgresql"
	"learn-back/internal/service"
	pkgHttp "learn-back/pkg/http"
)

const (
	defaultPort = ":8080"
)

func main() {
	// Загрузка переменных окружения из .env файла
	godotenv.Load(".env")

	// Создание нового маршрутизатора
	router := mux.NewRouter().StrictSlash(true)

	// Подключение к базе данных
	db, err := repository.ConnectToPostgres(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	// Создание сервисов
	airportService := service.NewAirportrService(repository.NewAirportRepository(db))
	aircraftService := service.NewAircraftService(repository.NewAicraftRepository(db))

	// Регистрация маршрутов
	pkgHttp.InitRoutes(router, airportService, aircraftService)

	// Настройка сервера
	srv := &http.Server{
		Addr:         defaultPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Запуск сервера
	go func() {
		log.Println("Starting server on port", defaultPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Обработка сигналов для корректного завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Остановка сервера с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}
	log.Println("Server exited properly")
}
