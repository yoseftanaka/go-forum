package middlewares

import (
	"fmt"
	"forum/main/config"
	"forum/main/constants"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get token from the Authorization header
		tokenString := c.Request().Header.Get("Authorization")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		userIDString := fmt.Sprintf("%v", claims[constants.USER_ID])
		_, err = config.RedisClient.Get(config.RedisContext, fmt.Sprintf("%s-%s", constants.USER, userIDString)).Result()

		if err == redis.Nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token Expired")
		} else if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Something wrong with redis")
		}

		// Token is valid, proceed
		return next(c)
	}
}
