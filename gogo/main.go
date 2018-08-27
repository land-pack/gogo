package main

import "db"
import "router"

func main(){
    db.Ping()
    router := router.InitRouter()
    router.Run(":8000")
}
