package routes

import "github.com/gin-gonic/gin"

type SysUserApi struct{}

func (s *SysUserApi) InitSysUserRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	router := Router.Group("user")
	{
		router.GET("info", sysUserController.GetUserInfo)
	}
}
