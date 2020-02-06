package middlewares

import (
	"github.com/astaxie/beego/context"
	"tng/common/logger"
)

// BaseMiddleware is basic for all middleware.
type baseMiddleware struct{}

func (m *baseMiddleware) response(c *context.Context, code int, data interface{}) {
	c.Output.SetStatus(code)
	err := c.Output.JSON(data, false, false)
	if err != nil {
		logger.Errorf(c.Request.Context(), "Decoding JSON in middleware: %v", err)
	}
}
