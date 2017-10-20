package router

import (
	Auth "zan/auth"
	Controller "zan/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Register 注册所有路由
func Register(e *echo.Echo) {
	// Login route
	e.POST("/login", Controller.Login)

	// Unauthenticated route
	e.GET("/", Controller.Accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	r.Use(middleware.JWTWithConfig(Auth.Config))
	r.GET("", Controller.Restricted)

}
