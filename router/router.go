package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "go-gin-backend-admin/docs"
	"go-gin-backend-admin/handler/mock"
	"go-gin-backend-admin/handler/user"
	"go-gin-backend-admin/middleware/jwt"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New((func() cors.Config {
		config := cors.DefaultConfig()
		config.AddAllowHeaders("Authorization")
		config.AllowAllOrigins = true
		// config.AllowCredentials = true
		return config
	})()))

	// 404 Handler.
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	u := r.Group("/user")
	{
		u.POST("/login", user.Login)
		u.POST("/register", user.Create)
	}

	v1 := r.Group("/")
	v1.Use(jwt.JWT())
	{
		v1.GET("/user", user.Get)
	}

	mock1 := r.Group("mock")
	{
		mock1.GET("sys/log/count", mock.SysLogCount)
	}

	return r
}
