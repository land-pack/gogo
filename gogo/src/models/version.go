package models

import (
    "log"
    "db"
    "fmt"
)


type MySQLVersion struct {
    Version string `json:"version"`
}

func FetchVersion(){
    //db.Ping()
    sql := `SELECT VERSION() as version;`
    var version  string

    stmt, err := db.SqlDB.Prepare(sql)

    fmt.Println("xxxx")
    if err != nil {
        log.Fatalln(err)
    }
    stmt.QueryRow().Scan(&version)
    fmt.Println(version)
    //return version
}
