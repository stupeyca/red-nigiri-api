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
	router.GET("/sign-up", controller.SignUp)

	return router
}

func (s *Server) IndexHandler(context *gin.Context) {
	runtimeVersion := runtime.Version()
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"runtimeVersion": runtimeVersion,
	})
}
