package api

import (
    "github.com/gin-gonic/gin"
    "models"
    "lib"
)

func AccountHandler(c *gin.Context) {
    var account = models.FetchAccount()
    lib.Render(c, "200", "OK", account)
}
