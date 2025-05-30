package server

import (
	"net/http"

	"github.com/kp/flight-order/flight"
	"github.com/labstack/echo/v4"
)

func FlightRouterGroup(servicePrefix string, journeyService flight.FlightJourneyService, middlewares ...echo.MiddlewareFunc) RouterGroup {
	return RouterGroup{
		Prefix:      servicePrefix,
		Routes:      flightRoutes(journeyService),
		Middlewares: middlewares,
	}
}

func flightRoutes(journeyService flight.FlightJourneyService) []Route {
	f := flight.NewFlightController(journeyService)
	return []Route{
		{
			Method:  http.MethodPost,
			Path:    "/journey/order/",
			Handler: f.GetFlightOrderView,
		},
	}
}
