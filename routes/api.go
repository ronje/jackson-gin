package routes

import (
	"github.com/gin-gonic/gin"
	"jackson-gin/app/controllers/app"
	"jackson-gin/app/controllers/common"
	"jackson-gin/app/middleware"
	"jackson-gin/app/services"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
}
