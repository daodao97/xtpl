package api

import (
	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xlog"
	"github.com/gin-gonic/gin"
)

func SetupRouter(e *gin.Engine) {
	e.GET("/ping", Ping)

	g := e.Group("/v1").Use(func(ctx *gin.Context) {
		// auth := ctx.GetHeader("Authorization")
		// if auth == "" {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()
	})

	g.POST("/test", xapp.RegisterAPI(test))
	g.GET("/page", xapp.RegisterAPI(Page))
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

func Page(ctx *gin.Context, page xapp.Page) (*xapp.PageResult[any], error) {
	xlog.Info("page: %v", page)
	return &xapp.PageResult[any]{}, nil
}
