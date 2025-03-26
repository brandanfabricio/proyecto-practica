package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var jwtSecret = []byte("secreto")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(username string) (string, error) {
	claims := &JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // Expira en 1 hora
		},
	}

	// Crear el token con los claims y firmarlo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func main() {

	server := echo.New()

	testMidlle := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token no proporcionado")
			}

			tokenString := authHeader[len("Bearer "):] // Extraer el token del header

			token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token inválido")
			}

			// Guardar los claims en el contexto
			claims := token.Claims.(*JWTClaims)
			c.Set("username", claims.Username)

			return next(c)
		}
	}

	server.POST("/login", func(c echo.Context) error {
		fmt.Println("aki")
		new_user := new(User)
		fmt.Println(new_user)

		if err := c.Bind(new_user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
		}
		fmt.Println(new_user)

		token, err := generateToken(new_user.Name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generando token"})
		}
		return c.JSON(http.StatusOK, map[string]string{"token": token})
	})

	server.POST("/user", func(c echo.Context) error {
		user_name := c.Get("username").(string)
		fmt.Println(user_name)

		return c.JSON(http.StatusCreated, user_name)
	}, testMidlle)

	server.Use(middleware.CORS())

	server.Logger.Fatal(server.Start(":5000"))
}

// cors
