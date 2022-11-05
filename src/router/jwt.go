package router

import (
	"net/http"
	"os"
	"time"
	"truphone-api/src/models"
	"truphone-api/src/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var CorsConfig = middleware.CORSConfig{
	AllowOrigins: []string{"https://localhost:1323", "https://localhost:1323"},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
}

// Configure middleware with the custom claims type
var Config = middleware.JWTConfig{
	Claims:     &models.JwtCustomClaims{},
	SigningKey: []byte("secret"),
}

// @Summary      Auth Login
// @Description  POST models.Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Login   body     models.Login  true  "Login"
// @Success      200  {object}	string
// @Failure      400  {object}  errs.Handling
// @Router       /auth [post]
func Auth(c echo.Context) error {
	utils.LoadEnv()
	login := new(models.Login)
	if err := c.Bind(&login); err != nil {
		c.JSON(http.StatusBadRequest, "Error to Bind Login")
	}

	user := os.Getenv("AUTH_USER")
	pass := os.Getenv("AUTH_PASS")

	// Throws unauthorized error
	if login.Username != user || login.Password != pass {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &models.JwtCustomClaims{
		Name:  "root",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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

// @Summary      Health Check Endpoint
// @Description  Check if server is running
// @Tags         Health
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /health_check/ [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server Avalible!")
}
