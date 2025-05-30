package server

import (
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
)

var defaultMiddlewares = []echo.MiddlewareFunc{recoveryMiddleware()}

func InitServer(commonMiddlewares []echo.MiddlewareFunc, opts ...ServerOpts) *echo.Echo {
	e := echo.New()

	// Disable debug mode in production environments
	appEnv := os.Getenv("APP_ENV")
	if strings.EqualFold(appEnv, "production") || strings.EqualFold(appEnv, "prod") {
		e.Debug = false
	}

	if commonMiddlewares == nil {
		commonMiddlewares = []echo.MiddlewareFunc{}
	}

	e.GET("/health/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	// Apply middleware
	for _, middleware := range commonMiddlewares {
		e.Use(middleware)
	}

	// Apply server options
	for _, opt := range opts {
		opt(e)
	}

	return e
}

func CreateRoutes(g ...RouterGroup) ServerOpts {
	return func(e *echo.Echo) {
		for _, group := range g {
			if len(group.Middlewares) < 1 {
				group.Middlewares = []echo.MiddlewareFunc{}
			}
			middlewares := append(defaultMiddlewares, group.Middlewares...)
			routerGroup := e.Group(group.Prefix, middlewares...)
			routeGenerator := WithRoutes(group.Routes)
			routeGenerator(routerGroup)
		}
	}
}

func WithRoutes(routes []Route) RouterOpts {
	return func(r *echo.Group) {
		for _, route := range routes {
			r.Add(route.Method, route.Path, route.Handler)
		}
	}
}

func recoveryMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4<<10) // 4 KB
					length := runtime.Stack(stack, false)
					traceMsg := strings.ReplaceAll(string(stack[:length]), "\n", " ")
					slog.Error("server error",
						"request_id", c.Get("request_id"),
						"rsession_id", c.Get("rsession_id"),
						"url", c.Request().RequestURI,
						"response_status", c.Response().Status,
						"method", c.Request().Method,
						"stack_trace", traceMsg,
						"response_code", http.StatusInternalServerError)
					c.NoContent(http.StatusInternalServerError)
				}
			}()
			return next(c)
		}
	}
}
