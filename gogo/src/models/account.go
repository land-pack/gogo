package models

import (
    "log"
    "db"
    "fmt"
    "define"
)

//api/v1/register
func  Register(uid int, balance_freeze float64, coinT define.CoinType) {
    fmt.Println(uid)
    coin_type := coinT.CoinCode
    coin_name := coinT.CoinName
    fmt.Printf("coin_type=%d | coin_name=%s\n", coin_type, coin_name)
    sql := `INSERT INTO 
            t_account(f_uid, f_balance_freeze, f_coin_type, f_coin_name)
            VALUES(?, ?, ?, ?)
            `
    stmt, err := db.SqlDB.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        log.Fatalln(err)
    }

    ret, err := stmt.Exec(uid, balance_freeze, coin_type, coin_name)
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
    Balance string `json:"balance"`
    BalanceFreeze string `json:"balance_freeze"`
}

type AccountData struct {
    Accounts []Account `json:"accounts"`
}

func FetchAccount(uid int, coin_type int) Account{
    sql := `SELECT f_uid as uid, f_coin_type as coin_type, f_coin_name as coin_name, 
                f_balance as balance, f_balance_freeze as balance_freeze
            FROM t_account
            WHERE f_uid=? AND f_coin_type=?`
    var account  Account

    stmt, err := db.SqlDB.Prepare(sql)
    defer stmt.Close()

    if err != nil {
        fmt.Println(err)
        log.Fatalln(err)
    }
    err = stmt.QueryRow(uid, coin_type).Scan(&account.Uid, &account.CoinType,
    &account.CoinName, &account.Balance, &account.BalanceFreeze)
    if err != nil {
        //log.Fatalln(err)
        fmt.Println(err)
     }
    return account
}

func FetchAccounts(uid int) AccountData{
    sql := `SELECT f_uid as uid, f_coin_type as coin_type, f_coin_name as coin_name,
                f_balance as balance, f_balance_freeze as balance_freeze
            FROM t_account
            WHERE f_uid=?`
    var account  Account
    var accounts = make([]Account, 0)

    stmt, err := db.SqlDB.Prepare(sql)
    defer stmt.Close()

    if err != nil {
        fmt.Println(err)
        log.Fatalln(err)
    }
    rows, err := stmt.Query(uid)
    if err != nil {
        //log.Fatalln(err)
        fmt.Println(err)
     }
    for rows.Next() {
        rows.Scan(&account.Uid, &account.CoinType, &account.CoinName,
        &account.Balance, &account.BalanceFreeze)
        accounts = append(accounts, account)
    }

    return AccountData{Accounts:accounts}
}

func FetchBalance(uid int, coin_type int) Account{
    sql := `SELECT f_uid as uid, f_coin_type as coin_type, f_coin_name as coin_name, 
                f_balance as balance, f_balance_freeze as balance_freeze
            FROM t_account
            WHERE f_uid=? AND f_coin_type=?`
    var account  Account

    stmt, err := db.SqlDB.Prepare(sql)
    defer stmt.Close()

    if err != nil {
        fmt.Println(err)
        log.Fatalln(err)
    }
    err = stmt.QueryRow(uid, coin_type).Scan(&account.Uid, &account.CoinType,
    &account.CoinName, &account.Balance, &account.BalanceFreeze)
    if err != nil {
        //log.Fatalln(err)
        fmt.Println(err)
     }
    return account
}

