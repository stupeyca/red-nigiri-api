package server

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := gin.Default()
	router.LoadHTMLGlob("web/template/*")
	router.GET("/", s.IndexHandler)
	router.GET("/health", s.healthHandler)

	return router
}

func (s *Server) IndexHandler(c *gin.Context) {
	runtimeVersion := runtime.Version()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"runtimeVersion": runtimeVersion,
	})
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
