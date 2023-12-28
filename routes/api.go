package routes

import "go-gin/controller"

var (
	sysUserController = controller.SysUserController
)

type RouterGroup struct {
	SysUserApi
}

var RouterGroupApp = new(RouterGroup)
