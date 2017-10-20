package controller

import (
	"net/http"
	"time"
	Cls "zan/class"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login is 登陆
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "zan" && password == "zan" {

		// Set custom claims
		claims := &Cls.JwtCustomClaims{
			Name:  "Julian zan",
			Admin: true,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// Accessible is
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Restricted is
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Cls.JwtCustomClaims)
	return c.JSON(http.StatusOK, echo.Map{
		"user": claims,
	})
	// name := claims.Name
	// return c.String(http.StatusOK, "Welcome "+name+"!")
}
