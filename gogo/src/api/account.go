package api

import (
    "github.com/gin-gonic/gin"
    "models"
    "lib"
    "define"
)

type Request struct {
    Uid int `json:"uid"`
}

func AccountsHandler(c *gin.Context) {
    var req Request
    c.BindJSON(&req)
    var accounts = models.FetchAccounts(req.Uid)

    lib.Render(c, "200", "OK", accounts)
}


func AccountHandler(c *gin.Context) {
    var account = models.FetchAccount(12,22)
    lib.Render(c, "200", "OK", account)
}

func RegisterHandler(c *gin.Context){
    var req Request
    c.BindJSON(&req)
    balance_freeze := 0.0
    var newAccount = []define.CoinType{}

    ret := models.FetchAccount(req.Uid, define.SEELE.CoinCode)
    if (models.Account{}) == ret {
        models.Register(req.Uid, balance_freeze, define.SEELE)
        newAccount = append(newAccount, define.SEELE)
    }

    ret = models.FetchAccount(req.Uid, define.BTC.CoinCode)
    if (models.Account{}) == ret {
        models.Register(req.Uid, balance_freeze, define.BTC)
        newAccount = append(newAccount, define.BTC)
    }

    lib.Render(c,"222", "WW", newAccount)
}
