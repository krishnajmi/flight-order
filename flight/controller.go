package flight

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
Example request for Postman:
{
    "journies": [
        ["LAX", "DXB"],
        ["JFK", "LAX"],
        ["SFO", "SJC"],
        ["DXB", "SFO"]
    ]
}
*/

// GetFlights handles GET request to retrieve flights
func (fc *FlightController) GetFlightOrderView(c echo.Context) error {
	var flightOrderRequest FlightOrderRequest
	if err := c.Bind(&flightOrderRequest); err != nil {
		slog.Error("invalid request format", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	journeyOrder, err := fc.FlightJourneyService.GetJourneyOrder(c.Request().Context(), flightOrderRequest.Journies)
	if err != nil {
		slog.Error("error getting journey order", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"journey_order": journeyOrder,
	})
}
