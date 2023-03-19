package routes

import (
	"github.com/agusbasari29/xjx-biller-backend/helper"
	"github.com/gin-gonic/gin"
)

func DefineAuthApiRoutes(e *gin.Engine) {
	handlers := []helper.Handler{
		// AuthRoutes{},
	}
	var routes []helper.Route
	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}
	api := e.Group("/auth")
	for _, route := range routes {
		api.POST(route.Path, route.Handler...)
	}
}
