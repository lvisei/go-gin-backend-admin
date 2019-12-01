package router

import (
	// _ "go-gin-backend-admin/docs"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-backend-admin/middleware/jwt"
	"go-gin-backend-admin/router/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongs",
		})
	})

	r.POST("/auth", user.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	user := r.Group("/user")
	user.Use(jwt.JWT())

	user.GET("hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code":    200,
			"message": "This works",
			"data":    nil,
		})
	})

	return r
}
