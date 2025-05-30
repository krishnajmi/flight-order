package cmd

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kp/flight-order/flight"
	"github.com/kp/flight-order/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var (
	flightJourneyService flight.FlightJourneyService
)

func initServer() {
	flightJourneyService = flight.NewFlightJourneyService()
}

var apiCmd = &cobra.Command{
	Use:   "apis",
	Short: "Start the Flight-Order api server",
	Long:  "This is starting point of apis",
	Run: func(cmd *cobra.Command, args []string) {
		servicePrefix := "/flight/api/v0"
		router := server.InitServer(
			nil,
			server.CreateRoutes(
				server.FlightRouterGroup(servicePrefix, flightJourneyService),
			),
		)

		server := http.Server{
			Addr:    ":8000",
			Handler: router,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				slog.Error("server error", "error", err)
				panic(err)
			}
		}()

		slog.Info("server started", "addr", server.Addr)
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("server shutdown error", "error", err)
			os.Exit(1)
		}
		slog.Info("server exiting")
	},
}
