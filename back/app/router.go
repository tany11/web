package app

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/index")
	router.GET("/cast")
	return router
}
