package controller

import (
	// "Login-Backend/src/config"
	"github.com/gin-gonic/gin"
)

type RoutingRegisterFunc func(router *gin.RouterGroup)

var routingFunctions []RoutingRegisterFunc

func RegisterApiRoute(functions ...RoutingRegisterFunc) {
	routingFunctions = append(routingFunctions, functions...)
}

func SetupRouting(router *gin.Engine) {
	router.Use(PrepareInfo())
	apiRouter := router.Group("localhost:8080")

	for _, routingFunction := range routingFunctions {
		routingFunction(apiRouter)
	}
}
