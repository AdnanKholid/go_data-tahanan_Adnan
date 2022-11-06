package middlewares

import (
	"mini_project/app/tokens"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}))
}

func CheckTokenMiddlewareUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := tokens.ExtractTokenUser(c)

		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"messege": "invalid create token",
			})
		}
		return next(c)
	}
}

func CheckTokenMiddlewareStore(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		storeID := tokens.ExtractTokenStore(c)

		if storeID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"messege": "invalid create token",
			})
		}
		return next(c)
	}
}

func CheckTokenMiddlewareAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		adminID := tokens.ExtractTokenAdmin(c)

		if adminID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"messege": "invalid create token",
			})
		}
		return next(c)
	}
}
