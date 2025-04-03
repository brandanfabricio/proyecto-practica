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
	Username string `json:"user"`
	jwt.RegisteredClaims
}

var listUsers []User

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
			isAuth := false

			for _, user := range listUsers {
				fmt.Println(user)
				if user.Name == claims.Username {
					isAuth = true
					break
				}

			}

			if isAuth {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "User invalido")
			}
		}
	}

	api := server.Group("/api")

	api.POST("/register", func(c echo.Context) error {

		var user User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
		}
		listUsers = append(listUsers, user)

		fmt.Println(user)

		return c.JSON(http.StatusCreated, echo.Map{"msj": "User creado"})
	})

	api.POST("/login", func(c echo.Context) error {
		new_user := new(User)

		if err := c.Bind(new_user); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
		}

		token, err := generateToken(new_user.Name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generando token"})
		}
		return c.JSON(http.StatusOK, map[string]string{"token": token})
	})

	api.POST("/user", func(c echo.Context) error {
		user_name := c.Get("username").(string)

		return c.JSON(http.StatusCreated, user_name)
	}, testMidlle)

	server.Use(middleware.CORS())

	server.Logger.Fatal(server.Start(":5000"))
}

// cors
