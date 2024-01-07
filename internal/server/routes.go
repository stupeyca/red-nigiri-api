package server

import (
	"net/http"
	"red-nigiri-api/internal/controller"
	"runtime"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()
	router.LoadHTMLGlob("web/template/*")
	router.GET("/", s.IndexHandler)

	apiRoutes := router.Group("/api")

	publicRoutes := apiRoutes.Group("/auth")
	publicRoutes.POST("/sign-up", controller.SignUp)
	publicRoutes.POST("/sign-in", controller.SignIn)

	return router
}

func (s *Server) IndexHandler(context *gin.Context) {
	runtimeVersion := runtime.Version()
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"runtimeVersion": runtimeVersion,
	})
}
