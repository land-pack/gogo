package api

import (
    "github.com/gin-gonic/gin"
    "models"
    "lib"
    "fmt"
)

type Request struct {
    Uid int `json:"uid"`
}

func AccountHandler(c *gin.Context) {
    var account = models.FetchAccount(12,22)
    lib.Render(c, "200", "OK", account)
}

func RegisterHandler(c *gin.Context){
    var req Request
    c.BindJSON(&req)
    fmt.Println("req>>:", req)
    fmt.Println("uid>>", req.Uid)
    coin_type := 3545
    ret := models.FetchAccount(req.Uid, coin_type)
    fmt.Println("ret>", ret)
    //models.Register(19)
    if (models.Account{}) == ret {
        fmt.Println("Create new one")
    }else{
        fmt.Println(ret)

    }
    lib.Render(c,"222", "WW", "xxx")
}
