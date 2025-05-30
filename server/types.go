package server

import "github.com/labstack/echo/v4"

// Specific router/API group details
type RouterGroup struct {
	Prefix         string
	Routes         []Route
	Middlewares    []echo.MiddlewareFunc
	HealthEndPoint echo.HandlerFunc
}

// each rest end point with associated permission key
type Route struct {
	Method        string
	Path          string
	Handler       echo.HandlerFunc
	PermissionKey string
}

// Router options to assign routes
type RouterOpts func(*echo.Group)

// Server options to
type ServerOpts func(*echo.Echo)
