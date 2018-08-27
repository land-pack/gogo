package models

import (
    "log"
    "db"
    "fmt"
)





///api/v1/register
func  Register(uid int) {
    fmt.Println(uid)
    sql := `INSERT INTO 
            t_account(f_uid, f_balance_freeze, f_coin_type, f_coin_name)
            VALUES(?, ?, ?, ?)
            `
    stmt, err := db.SqlDB.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        log.Fatalln(err)
    }

    ret, err := stmt.Exec(12, 0, 1, "XX")
    if err != nil {
        log.Fatalln(err)
    }
    _id, err := ret.LastInsertId()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Printf("Insert successful %d\n", _id)
}






type Account struct {
    Uid int `json:"uid"`
    CoinType int `json:"coin_type"`
    CoinName string `json:"coin_name"`
}

func FetchAccount() Account{
    //db.Ping()
    sql := `SELECT f_uid as uid, f_coin_type as coin_type, f_coin_name as coin_name
            FROM t_account
            WHERE f_uid=?`
    var account  Account

    stmt, err := db.SqlDB.Prepare(sql)
    defer stmt.Close()

    if err != nil {
        fmt.Println(err)
        log.Fatalln(err)
    }
    err = stmt.QueryRow(88888).Scan(&account.Uid, &account.CoinType, &account.CoinName)
    if err != nil {
        //log.Fatalln(err)
        fmt.Println(err)
     }

    return account
}
