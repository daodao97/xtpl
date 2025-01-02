package api

import (
	"net/http"

	"github.com/daodao97/xgo/xapp"
	"github.com/gin-gonic/gin"
)

func SetupRouter(e *gin.Engine) {
	e.GET("/ping", Ping)

	g := e.Group("/v1").Use(func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	})

	g.POST("/test", xapp.RegisterAPI(test))
}

type TestApi struct {
	Mode string `json:"mode" form:"mode" binding:"required"`
}

type TestApiResponse struct {
	Message string `json:"message"`
}

func test(ctx *gin.Context, mode TestApi) (*TestApiResponse, error) {
	return &TestApiResponse{Message: "pong"}, nil
}
