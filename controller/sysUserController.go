package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sysUserController struct{}

var SysUserController = new(sysUserController)

func (s *sysUserController) GetUserInfo(ctx *gin.Context) {
	users, err := sysUserRepository.GetAllUsers()
	if err != nil {
		result.FailWithMsg(err.Error(), ctx)
	}
	result.OkWithData(users, ctx)
}

func (s *sysUserController) helloHandler(ctx *gin.Context) {
	result := make(map[string]string, 10)
	name := ctx.Param("name")
	result["message"] = "hello, " + name
	ctx.JSON(http.StatusOK, result)
}
