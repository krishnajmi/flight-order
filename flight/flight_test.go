package flight

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortService(t *testing.T) {
	tests := []struct {
		name        string
		graph       map[string]string
		expected    []string
		expectError bool
	}{
		{
			name:     "should sort simple linear path A->B->C->D",
			graph:    map[string]string{"A": "B", "B": "C", "C": "D"},
			expected: []string{"A", "B", "C", "D"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := newSortService()
			result, err := service.Sort(context.Background(), tt.graph)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFlightService_GetJourneyOrder(t *testing.T) {
	service := NewFlightJourneyService()

	tests := []struct {
		name        string
		journeys    [][]string
		expected    []string
		expectError bool
	}{
		{
			name:     "valid flight paths",
			journeys: [][]string{{"LAX", "DXB"}, {"JFK", "LAX"}, {"DXB", "SFO"}},
			expected: []string{"JFK", "LAX", "DXB", "SFO"},
		},
		{
			name:        "invalid journey format",
			journeys:    [][]string{{"A"}},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "cyclic path",
			journeys:    [][]string{{"LAX", "DXB"}, {"DXB", "JFK"}, {"JFK", "LAX"}},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "disconnected paths",
			journeys:    [][]string{{"LAX", "DXB"}, {"SFO", "JFK"}},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "duplicate paths",
			journeys:    [][]string{{"LAX", "DXB"}, {"LAX", "DXB"}},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetJourneyOrder(context.Background(), tt.journeys)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestNewFlightController(t *testing.T) {
	service := NewFlightJourneyService()
	controller := NewFlightController(service)

	assert.NotNil(t, controller)
	assert.Equal(t, service, controller.FlightJourneyService)
}
