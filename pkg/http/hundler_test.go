package http

import (
	"io/ioutil"
	"learn-back/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockAirportService имитирует поведение реального сервиса
type MockAirportService struct{}

// GetAllAirports возвращает фиксированные данные для тестирования
func (m *MockAirportService) GetAllAirports() ([]*models.Airport, error) {
	return []*models.Airport{
		{Code: "MJZ", Name: "Мирный", City: "Мирный", Longtitude: 62.534689, Latitude: 114.038928, Timezone: "Asia\\Yakutsk"},
		{Code: "NBC", Name: "Бегишево", City: "Бегишево", Longtitude: 55.34, Latitude: 52.06, Timezone: "Europe\\Moscow"},
	}, nil
}

func TestRegisterHandlers(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	mockService := &MockAirportService{}
	GetAllAirportsHundlers(mockService.GetAllAirports)

	req, err := http.NewRequest("GET", "/airports", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	resp := recorder.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	expectedBody := `[{"Code":"MJZ", "City":"Мирный", "Latitude":114.038925, "Longtitude":62.53469, "Name":"Мирный", "Timezone":"Asia\\Yakutsk"}, {"City":"Бегишево", "Code":"NBC", "Latitude":52.06, "Longtitude":55.34, "Name":"Бегишево", "Timezone":"Europe\\Moscow"}]`
	assert.JSONEq(t, expectedBody, strings.TrimSpace(string(body)))
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
