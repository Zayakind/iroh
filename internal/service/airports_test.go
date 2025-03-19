package service

import (
	"testing"

	"learn-back/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockAirportsRepository struct{}

type MockAircrafsRepository struct{}

func (m *MockAirportsRepository) GetAllAirports() ([]models.Airport, error) {
	return []models.Airport{
		{Code: "MJZ", Name: "Мирный", City: "Мирный", Longtitude: 62.534689, Latitude: 114.038928, Timezone: "Asia\\Yakutsk"},
		{Code: "NBC", Name: "Бегишево", City: "Бегишево", Longtitude: 55.34, Latitude: 52.06, Timezone: "Europe\\Moscow"},
	}, nil
}

func (m *MockAircrafsRepository) GetAllAircrafts() ([]models.Aircraft, error) {
	return []models.Aircraft{
		{Code: "773", Model: "Boeing 777-300", Range: "11100"},
		{Code: "763", Model: "Boeing 767-300", Range: "7900"},
	}, nil
}

func TestGetAllAirports(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockAirportsRepository{}
		service := NewAirportrService(mockRepo)

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
}

func TestGetAllAircraft(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockAircrafsRepository{}
		service := NewAircraftService(mockRepo)

		aircrafts, err := service.GetAllAircrafts()
		require.NoError(t, err)
		assert.Len(t, aircrafts, 2)
		assert.Equal(t, aircrafts[0].Code, "773")
		assert.Equal(t, aircrafts[0].Model, "Boeing 777-300")
		assert.Equal(t, aircrafts[0].Range, "11100")
	})
}
