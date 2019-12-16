package user

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/handler"
	"go-gin-backend-admin/model"
	"go-gin-backend-admin/pkg/errno"
	"go-gin-backend-admin/pkg/util"
)

// @Tags 用户模块
// @Summary 登陆授权
// @Produce  json
// @Param user body user.LoginRequest true "登陆授权"
// @Success 200 {object} handler.Response "{"code":20000,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	var u LoginRequest
	if err := c.Bind(&u); err != nil {
		handlerG.Response(errno.ErrBind, nil)
		return
	}

	// Validate the fields
	valid := validation.Validation{}
	if ok, err := valid.Valid(&u); err != nil || !ok {
		handlerG.Response(errno.INVALID_PARAMS, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		handlerG.Response(errno.ERROR_User_Not_Found, nil)
		return
	}

	// Compare the login password with the user password.
	if err := d.Compare(u.Password); err != nil {
		handlerG.Response(errno.ERROR_Password_Incorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := util.GenerateToken(util.TokenContext{ID: d.BaseModel.ID, Username: d.Username, Password: d.Password}, "")
	if err != nil {
		handlerG.Response(errno.ERROR_AUTH_TOKEN, nil)
		return
	}

	handlerG.Response(nil, model.Token{Token: t})
}
