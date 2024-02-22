package controller

// 做个日志
import (
	"ops_client/internal/middleware"
	_ "ops_client/swagger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.Cors())
	// ---------API版本区分----------
	v1 := r.Group("/api/v1")
	// ------------验证相关------------
	v1.Use(middleware.AuthMiddleware()).Use(middleware.RecordLog())
	{
		// -------------接口测试--------------
		v1.GET("ping1", Test1)
		// ------------通用相关---------------
		generalRoute := v1.Group("general")
		generalRoute.GET("server-info", GetServerInfo)
		generalRoute.POST("exec-command", ExecCommand)
		// ------------更新相关----------------
		updateRouter := v1.Group("update")
		updateRouter.POST("server", UpdateServer)
		updateRouter.POST("exec", UpdateExec)
	}
	return r
}
