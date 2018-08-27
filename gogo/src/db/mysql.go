package db

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


func Ping(){
    fmt.Println("Ping ....")
}


var SqlDB * sql.DB

func InitDB(){
    var err error
    SqlDB, err = sql.Open("mysql", "frank:openmysql@tcp(192.168.1.203:3306)/gotest?parseTime=true")

    if err != nil {
        log.Fatalln(err)
    }
    err = SqlDB.Ping()
    if err != nil {
        log.Fatalln(err.Error())
    }
}
