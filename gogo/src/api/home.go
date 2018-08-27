package api

import (
    //"net/http"
    //"gopkg.in/gin-gonic/gin.v1"
    "github.com/gin-gonic/gin"
    "models"
    "lib"
)

func IndexHandler(c *gin.Context) {
    var version = models.FetchVersion()

    //c.String(http.StatusOK, "It works")
    lib.Render(c, "200", "OK", version)
}
