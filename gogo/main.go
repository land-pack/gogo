package main

import "db"
import "router"


func main(){
    db.InitDB()
    db.Ping()
    //models.FetchVersion()
    //sql_cnx := db.SqlDB
    //defer sql_cnx.Close()
    router := router.InitRouter()
    router.Run(":8000")
    db.SqlDB.Close()
}
