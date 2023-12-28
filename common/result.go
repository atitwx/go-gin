package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	ERROR   = 201
	SUCCESS = 200
)

var Result = &result{}

func res(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, result{
		code,
		data,
		msg,
	})
}

func (r *result) Ok(ctx *gin.Context) {
	res(SUCCESS, map[string]interface{}{}, "操作成功", ctx)
}

func (r *result) OkWithData(data interface{}, ctx *gin.Context) {
	res(SUCCESS, data, "操作成功", ctx)
}

func (r *result) OkWithMsg(msg string, ctx *gin.Context) {
	res(SUCCESS, map[string]interface{}{}, msg, ctx)
}

func (r *result) Fail(ctx *gin.Context) {
	res(ERROR, map[string]interface{}{}, "操作失败", ctx)
}

func (r *result) FailWithMsg(msg string, ctx *gin.Context) {
	res(ERROR, map[string]interface{}{}, msg, ctx)
}
