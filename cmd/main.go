package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pebruwantoro/technical_test_dealls/config"
	"github.com/pebruwantoro/technical_test_dealls/handler"
	"github.com/pebruwantoro/technical_test_dealls/helper"
	"github.com/pebruwantoro/technical_test_dealls/repository"
	"github.com/pebruwantoro/technical_test_dealls/usecase"
)

func main() {
	e := echo.New()

	server := newServer()

	users := e.Group("users")
	users.POST("/sign-up", server.SignUp)
	users.POST("/login", server.Login)
	users.POST("/swipe", server.Swipe, JWTMiddleware(config.JWT_SECRET_KEY))
	users.POST("/purchase", server.PurchasePremiumPackage, JWTMiddleware(config.JWT_SECRET_KEY))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization},
		AllowCredentials: true,
		MaxAge:           600,
	}))

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(config.PORT))
}

func newServer() *handler.Server {
	dbDsn := config.DATABASE

	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	var usecase usecase.UsecaseInterface = usecase.NewUsecase(repo)

	opts := handler.NewUsecaseOptions{
		Usecase: usecase,
	}

	return handler.NewServer(opts)
}

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			tokenStr := parts[1]

			claims := &helper.Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return []byte(secretKey), nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token: "+err.Error())
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			c.Set("claims", claims)

			return next(c)
		}
	}
}
