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

func RegisterHandler(c *gin.Context){
    models.Register(19)
    lib.Render(c,"222", "WW", "xxx")
}
