package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/handler"
	"go-gin-backend-admin/pkg/errno"
	"go-gin-backend-admin/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerG := handler.Gin{C: c}

		token := c.Request.Header.Get("Authorization")

		if len(token) == 0 {
			handlerG.Response(errno.ERROR_Not_TOKEN_EXIST, nil)
			c.Abort()
			return
		}

		ctx, err := util.ParseToken(token, "")
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				err = errno.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				err = errno.ERROR_AUTH_CHECK_TOKEN_FAIL
			}

			handlerG.Response(err, nil)
			c.Abort()
			return
		}

		c.Set("jwt", ctx.ID)

		c.Next()

	}
}
