package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/pkg/consts"
	"go-gin-backend-admin/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = consts.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = consts.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = consts.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = consts.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != consts.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  consts.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
