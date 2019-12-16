package user

import (
	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/handler"
	"go-gin-backend-admin/model"
	"go-gin-backend-admin/pkg/errno"
)

// @Tags 用户模块
// @Summary 获取当前登录用户的用户信息
// @Produce  json
// @Param Authorization header string true "token"
// @Success 200 {object} model.UserInfo "{"code":20000,"message":"OK","data": UserInfo }"
// @Failure 20003 {object} handler.Response "{"code":20003,"message":"Token has timed out.","data": null }"
// @Router /user [get]
func Get(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	jwt := c.MustGet("jwt")
	ctx := jwt.(uint64)

	id := ctx

	// Get the user by the `id` from the database.
	u, err := model.GetUserById(id)
	if err != nil {
		handlerG.Response(errno.ERROR_User_Not_Found, nil)
		return
	}

	user := model.UserInfo{u.ID, u.Username, u.Email, u.Phone, u.CreatedOn, u.ModifiedOn}

	handlerG.Response(nil, user)
}
