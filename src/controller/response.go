package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	SigOk                 = 0x00
	SigInternalFailed     = 0xFF
	SigAuthFailed         = 0xF0
	SigAccessDenied       = 0xF1
	SigRobotTestFailed    = 0xF2
	SigAccountNotVerified = 0xF3
)

type Response struct {
	Code 	int 		`json:"code"`
	Data 	interface{}	`json:"data"`
	Message string		`json:"message"`
}

func sendResponse(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx *gin.Context) {
	sendResponse(ctx, SigOk, nil, "succeeded")
}

func OkWithMessage(ctx *gin.Context,msg string) {
	sendResponse(ctx, SigOk, nil, msg)
}

func OkWithData(c *gin.Context, data interface{}) {
	sendResponse(c, SigOk, data, "succeeded")
}

func OkWithDetail(c *gin.Context, data interface{}, message string) {
	sendResponse(c, SigOk, data, message)
}

func InternalFailed(c *gin.Context) {
	sendResponse(c, SigInternalFailed, nil, "failed")
}

func InternalFailedWithMessage(c *gin.Context, message string) {
	sendResponse(c, SigInternalFailed, nil, message)
}

func InternalFailedWithDetail(c *gin.Context, data interface{}, message string) {
	sendResponse(c, SigInternalFailed, data, message)
}

func AuthFailed(c *gin.Context) {
	sendResponse(c, SigAuthFailed, nil, "please login first")
}

func AccessDenied(c *gin.Context) {
	sendResponse(c, SigAccessDenied, nil, "access denied")
}

func RobotTestFailed(c *gin.Context) {
	sendResponse(c, SigRobotTestFailed, nil, "hey robot")
}

func AccountNotVerified(c *gin.Context) {
	sendResponse(c, SigAccountNotVerified, nil, "please verify your account first")
}