package api

import (
    "net/http"
    //"gopkg.in/gin-gonic/gin.v1"
    "github.com/gin-gonic/gin"
    "models"
)

func IndexHandler(c *gin.Context) {
    models.FetchVersion()
    c.String(http.StatusOK, "It works")
}
