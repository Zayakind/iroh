package http

import (
	"io/ioutil"
	"learn-back/internal/models"
	"learn-back/internal/service"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockAirportService имитирует поведение реального сервиса
type MockAirportRepository struct{}

type MockAirportService struct{}

// GetAllAirports возвращает фиксированные данные для тестирования
func (m *MockAirportRepository) GetAllAirports() ([]models.Airport, error) {
	return []models.Airport{
		{
			Code:       "MJZ",
			City:       "Мирный",
			Latitude:   114.038925,
			Longtitude: 62.53469,
			Name:       "Мирный",
			Timezone:   "Asia/Yakutsk",
		},
		{
			Code:       "NBC",
			City:       "Бегишево",
			Latitude:   52.06,
			Longtitude: 55.34,
			Name:       "Бегишево",
			Timezone:   "Europe/Moscow",
		},
	}, nil
}

func (mas *MockAirportService) GetAllAirports() ([]models.Airport, error) {
	return []models.Airport{
		{
			Code:       "MJZ",
			City:       "Мирный",
			Latitude:   114.038925,
			Longtitude: 62.53469,
			Name:       "Мирный",
			Timezone:   "Asia/Yakutsk",
		},
		{
			Code:       "NBC",
			City:       "Бегишево",
			Latitude:   52.06,
			Longtitude: 55.34,
			Name:       "Бегишево",
			Timezone:   "Europe/Moscow",
		},
	}, nil
}

type MockAircraftRepository struct{}

type MockAircraftService struct{}

// GetAllAirports возвращает фиксированные данные для тестирования
func (m *MockAircraftRepository) GetAllAircrafts() ([]models.Aircraft, error) {
	return []models.Aircraft{
		{Code: "773", Model: "Boeing 777-300", Range: "11100"},
		{Code: "763", Model: "Boeing 767-300", Range: "7900"},
	}, nil
}

func (mas *MockAircraftService) GetAllAircrafts() ([]models.Aircraft, error) {
	return []models.Aircraft{
		{Code: "773", Model: "Boeing 777-300", Range: "11100"},
		{Code: "763", Model: "Boeing 767-300", Range: "7900"},
	}, nil
}

func TestSuccessfulResponse(t *testing.T) {
	t.Run("Success Airport", func(t *testing.T) {
		// Создаем моки репозитория
		mockRepo := &MockAirportRepository{}

		// Создаем службу, используя моки репозитория
		srv := service.NewAirportrService(mockRepo)

		// Создаем маршрутизатор
		router := mux.NewRouter().StrictSlash(true)

		// Регистрируем маршруты с использованием сервиса
		router.HandleFunc("/airports", GetAllAirportsHandler(srv)).Methods(http.MethodGet)

		// Создаем запрос
		req, err := http.NewRequest("GET", "/airports", nil)
		require.NoError(t, err)

		// Создаем рекордер для записи ответа
		recorder := httptest.NewRecorder()

		// Обрабатываем запрос
		router.ServeHTTP(recorder, req)

		// Получаем результат
		resp := recorder.Result()
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		// Ожидаемый ответ
		expectedBody := `[{"Code":"MJZ", "City":"Мирный", "Latitude":114.038925, "Longtitude":62.53469, "Name":"Мирный", "Timezone":"Asia/Yakutsk"}, {"City":"Бегишево", "Code":"NBC", "Latitude":52.06, "Longtitude":55.34, "Name":"Бегишево", "Timezone":"Europe/Moscow"}]`

		// Сравниваем полученные данные с ожидаемыми
		assert.JSONEq(t, expectedBody, strings.TrimSpace(string(body)))
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("Success Aircraft", func(t *testing.T) {
		// Создаем моки репозитория
		mockRepo := &MockAircraftRepository{}

		// Создаем службу, используя моки репозитория
		srv := service.NewAircraftService(mockRepo)

		// Создаем маршрутизатор
		router := mux.NewRouter().StrictSlash(true)

		// Регистрируем маршруты с использованием сервиса
		router.HandleFunc("/aircrafts", GetAllAircraftsHandler(srv)).Methods(http.MethodGet)

		// Создаем запрос
		req, err := http.NewRequest("GET", "/aircrafts", nil)
		require.NoError(t, err)

		// Создаем рекордер для записи ответа
		recorder := httptest.NewRecorder()

		// Обрабатываем запрос
		router.ServeHTTP(recorder, req)

		// Получаем результат
		resp := recorder.Result()
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		// Ожидаемый ответ
		expectedBody := `[{"aircraft_code":"773", "model":"Boeing 777-300", "range":"11100"}, {"aircraft_code":"763", "model":"Boeing 767-300", "range":"7900"}]`

		// Сравниваем полученные данные с ожидаемыми
		assert.JSONEq(t, expectedBody, strings.TrimSpace(string(body)))
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	})
}
