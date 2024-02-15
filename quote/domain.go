package quote

type Market struct {
    Region   string `json:"region"`
    Status   string `json:"status"`
    TimeZone string `json:"time_zone"`
}

type Quote struct {
    Symbol    string `json:"symbol"`
    Code      string `json:"code"`
    Exchange  string `json:"exchange"`
    Name      string `json:"name"`
    Status    string `json:"status"`
    Current   string `json:"current"`
    Currency  string `json:"currency"`
    Timestamp int64  `json:"timestamp"`
    Open      string `json:"open"`
    LastClose string `json:"last_close"`
    High      string `json:"high"`
    Low       string `json:"low"`
    AvgPrice  string `json:"avg_price"`
    Volume    string `json:"volume"`
    Amount    string `json:"amount"`
    Amplitude string `json:"amplitude"`
}

type Item struct {
    Market Market `json:"market"`
    Quote  Quote  `json:"quote"`
}

type Data struct {
    Items     []*Item `json:"items"`
    ItemsSize int     `json:"items_size"`
}

type QuoteResp struct {
    Data             Data   `json:"data"`
    ErrorCode        int    `json:"error_code"`
    ErrorDescription string `json:"error_description"`
}

type QuoteReq struct {
    Symbols  []string `json:"symbols"`
    Exchange string   `json:"exchange"`
}
