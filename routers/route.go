package routers

import (
	"blog/controller"
	"blog/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	r := gin.Default()

	v1 := r.Group("/api/base")
	{
		v1.POST("/signup", controller.SignupHandler)
		v1.POST("/login", controller.LoginHandler)
		// 使用中间件的路由
		v1.GET("/index", middlewares.JWTAuthMiddleware(), controller.ResponseDataAboutIndex)
		v1.GET("/essay", middlewares.JWTAuthMiddleware(), controller.ResponseDataAboutEssay)
	}

	v2 := r.Group("/api/manager")
	{
		v2.POST("/addClassify", controller.AddClassifyHandler)
		v2.POST("/addEssay", controller.AddEssayHandler)
		v2.PUT("/updateEssay", controller.UpdateEssayHandler)
	}

	r.Use(middlewares.JWTAuthMiddleware())
	{
		r.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}