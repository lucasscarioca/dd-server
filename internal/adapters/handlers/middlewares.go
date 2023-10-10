package handlers

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func customLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} - ${uri} [${method} - ${status}] ${latency_human} - ${error}\n",
	})
}

func customCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("HTTP_ALLOWED_ORIGINS")},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	})
}

// func cacheControl(maxAge time.Duration) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			v := "no-cache, no-store"
// 			if maxAge > 0 {
// 				v = fmt.Sprintf("public, max-age=%.0f", maxAge.Seconds())
// 			}
// 			c.Response().Header().Set("Cache-Control", v)
// 			return next(c)
// 		}
// 	}
// }
