package auth

import (
	Cls "zan/class"

	"github.com/labstack/echo/middleware"
)

// Config middleware with the custom claims type
var Config = middleware.JWTConfig{
	Claims:     &Cls.JwtCustomClaims{},
	SigningKey: []byte("secret"),
}
