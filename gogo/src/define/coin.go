package define

type CoinType struct {
    CoinName string `json:"coin_name"`
    CoinCode int `json:"coin_code"`
}


var SEELE CoinType= CoinType{"seele", 3545}
var BTC CoinType= CoinType{"btc", 3546}

