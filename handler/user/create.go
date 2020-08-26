package user

import (
	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/handler"
	"go-gin-backend-admin/model"
	"go-gin-backend-admin/pkg/errno"
)

// @Tags 用户模块
// @Summary 用户注册
// @Produce  json
// @Accept json
// @Param user body user.CreateRequest true "用户注册"
// @Success 200 {object} handler.Response "{"code":20000,"data":{"name": "admin"},"msg":"ok"}"
// @Router /user/register [post]
func Create(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	var u model.User

	if err := c.Bind(&u); err != nil {
		handlerG.Response(errno.ErrBind, nil)
		return
	}

	// Validate the fields
	if ok, err := u.Validate(); err != nil || !ok {
		handlerG.Response(errno.INVALID_PARAMS, nil)
		return
	}

	if _, err := model.GetUser(u.Username); err == nil {
		handlerG.Response(errno.ERROR_EXIST_USER, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handlerG.Response(errno.ERROR_Encrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handlerG.Response(errno.ERROR_ADD_User_FAIL, nil)
		return
	}

	handlerG.Response(nil, map[string]interface{}{"name": u.Username})
}
