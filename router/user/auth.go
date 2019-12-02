package user

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/pkg/app"
	"go-gin-backend-admin/pkg/consts"
	"go-gin-backend-admin/pkg/util"
	auth_service "go-gin-backend-admin/service/auth"
)

type auth struct {
	Username string `valid:"Required; MinSize(5) MaxSize(50)" example:"admin"`
	Password string `valid:"Required;  MinSize(6) MaxSize(50)" example:"123456"`
}

// @Tags 用户模块
// @Summary 登陆授权
// @Produce  json
// @Param auth body user.auth true "登陆授权"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	var a auth

	c.BindJSON(&a)

	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: a.Username, Password: a.Password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, consts.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
		"token": token,
	})
}
