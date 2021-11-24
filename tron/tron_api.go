package tron

type TRC20 struct {
	TokenID          string  `json:"tokenId"`
	Balance          string  `json:"balance"`
	TokenName        string  `json:"tokenName"`
	TokenAbbr        string  `json:"tokenAbbr"`
	TokenDecimal     int     `json:"tokenDecimal"`
	TokenType        string  `json:"tokenType"`
	TokenPriceInTrx  float64 `json:"tokenPriceInTrx"`
}

type Balance struct {
	TokenPriceInTrx float64    `json:"tokenPriceInTrx"`
	TokenID         string `json:"tokenId"`
	Balance         string `json:"balance"`
	TokenName       string `json:"tokenName"`
	TokenDecimal    int    `json:"tokenDecimal"`
	TokenAbbr       string `json:"tokenAbbr"`
	TokenType       string `json:"tokenType"`
}

type Account struct {
	TRC20 []TRC20 `json:"trc20token_balances"`
	TRC10 []Balance `json:"tokenBalances"`
	Balances []Balance `json:"balances"`
}
