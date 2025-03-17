package service

import (
	"testing"

	"learn-back/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockAirportsRepository имитирует поведение реального репозитория
type MockAirportsRepository struct{}

// GetAllAirports возвращает фиксированные данные для тестирования
func (m *MockAirportsRepository) GetAllAirports() ([]*models.Airport, error) {
	return []*models.Airport{
		{Code: "MJZ", Name: "Мирный", City: "Мирный", Longtitude: 62.534689, Latitude: 114.038928, Timezone: "Asia\\Yakutsk"},
		{Code: "NBC", Name: "Бегишево", City: "Бегишево", Longtitude: 55.34, Latitude: 52.06, Timezone: "Europe\\Moscow"},
	}, nil
}

func TestGetAllAirports(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockAirportsRepository{}
		service := (mockRepo)

		users, err := service.GetAllAirports()
		require.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, users[0].Code, "MJZ")
		assert.Equal(t, users[0].Name, "Мирный")
		assert.Equal(t, users[0].City, "Мирный")
		assert.Equal(t, users[0].Longtitude, float32(62.534689))
		assert.Equal(t, users[0].Latitude, float32(114.038928))
		assert.Equal(t, users[0].Timezone, "Asia\\Yakutsk")
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo := &MockAirportsRepository{}
		service := NewAirportrService(mockRepo)

		_, err := service.GetAllAirports()
		require.NoError(t, err)
	})
}
