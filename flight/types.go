package flight

import "context"

// FlightController handles HTTP requests related to flights
type FlightController struct {
	FlightJourneyService
}

// NewFlightController creates a new instance of FlightController
func NewFlightController(journeyService FlightJourneyService) *FlightController {
	return &FlightController{FlightJourneyService: journeyService}
}

func NewFlightJourneyService() FlightJourneyService {
	return &flightJourney{}
}

type FlightOrderRequest struct {
	Journies [][]string `json:"journies"` // Each journey is an array of 2 strings [source, destination]
}

type TravelJourney struct {
	Source      string `json:"source"`      // First element represents source
	Destination string `json:"destination"` // Second element represents destination
}

type flightJourney struct {
	Journies []TravelJourney `json:"journies"`
}

type FlightJourneyService interface {
	GetJourneyOrder(ctx context.Context, journies [][]string) ([]string, error)
}

func newSortService() sortService {
	return &topologicalSorter{}
}

type sortService interface {
	Sort(ctx context.Context, graph map[string]string) ([]string, error)
}

type topologicalSorter struct {
}
