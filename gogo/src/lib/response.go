package lib

import (
    "github.com/go-sql-driver/mysql"
    "net/http"
)

type Response struct {
    Code string `json:"code"`
    Message string `json:"msg"`
    Data interface{} `json:"data"`
}

func Render(c *gin.Context, code string, message string, data interface{}){
    c.JSON(http.StatusOK, Response{
        Code: code,
        Message: message,
        Data:data,
    })
}
