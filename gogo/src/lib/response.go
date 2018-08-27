package lib

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Code string `json:"status"`
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
