package middlware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/uchupx/kajian-api/pkg/logger"
)

func (m *Middleware) Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := time.Now()

		next(c)

		latency := time.Since(t)
		meta := logMeta{
			Methods: c.Request().Method,
			Path:    c.Request().URL.Path,
			Latency: fmt.Sprintf("%d ms", latency.Milliseconds()), // Log latency 2-10 ms tolerance
			Status:  c.Response().Status,                          // Log response http status
			// Response: c.Response(),
		}

		logger.Logger.WithFields(logrus.Fields{
			"meta": meta, // Log response body
		}).Info(http.StatusText(c.Response().Status))
		return nil
	}
}
