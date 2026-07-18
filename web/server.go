package web

import (
	"time"

	"github.com/daodao97/gossr"
	"github.com/gin-gonic/gin"
)

type homePayload struct {
	Message     string
	Path        string
	GeneratedAt string
}

func (payload homePayload) AsMap() map[string]any {
	return map[string]any{
		"message":     payload.Message,
		"path":        payload.Path,
		"generatedAt": payload.GeneratedAt,
	}
}

// SetupRouter mounts the public SSR site after the API and admin routes.
func SetupRouter(router *gin.Engine) error {
	data := gossr.NewDataEngine()
	data.GET("/", gossr.WrapSSR(func(c *gin.Context) (gossr.SSRPayload, error) {
		return homePayload{
			Message:     "xtpl SSR is ready",
			Path:        c.Request.URL.Path,
			GeneratedAt: time.Now().Format(time.RFC3339),
		}, nil
	}))

	return gossr.SsrWithOptions(router, Dist, gossr.Options{
		DataEngine: data,
	})
}
