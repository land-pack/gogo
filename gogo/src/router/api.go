package router

import (
    //"gopkg.in/gin-gonic/gin.v1"
    "github.com/gin-gonic/gin"
    "api"
)

func InitRouter() *gin.Engine{
    router := gin.Default()

    router.GET("/", api.IndexHandler)
    return router
}
