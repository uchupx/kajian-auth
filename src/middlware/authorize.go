package middlware

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/uchupx/kajian-auth/pkg/enums"
)

func (m *Middleware) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		reg := regexp.MustCompile(`Bearer[\s]`)

		auth = reg.ReplaceAllString(auth, "")
		if strings.TrimSpace(auth) != "" {
			if err := m.Redis.Get(c.Request().Context(), fmt.Sprintf(enums.RedisKeyAuthorizationToken, auth)).Err(); err == redis.Nil {
				return echo.NewHTTPError(401, "Unauthorized")
			}

			next(c)
		} else {
			return echo.NewHTTPError(401, "Unauthorized")
		}
		return nil
	}
}
