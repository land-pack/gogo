package api

import (
    "net/http"
    //"gopkg.in/gin-gonic/gin.v1"
    "github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
    c.String(http.StatusOK, "It works")
}
