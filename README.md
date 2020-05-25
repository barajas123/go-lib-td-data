#go-lib-td-data
Allows the use of the TD Ameritrade API to get delayed/real-time data(candles, quotes) with your account. 

## Getting Price History
	hp := NewHistoryProvider(apiKey,"SPXL")
	data := *hp.GetData()
	fmt.Println(data.Candles)
    
    // OUTPUT is and []Candle (OHLC)
    [{36.32 36.34 36.31 36.34 5530 1590145200000} {36.35 36.35 36.34 36.34 1100 1590145260000}...]
    
## Getting Quotes
	qp := data.NewQuoteProvider(apiKey,"SPXL")
	data := *qp.GetData()
	fmt.Println(data.Symbol.AskPrice)
	
	// OUTPUT
	{ETF EQUITY 25459W862 ETF SPXL Direxion Daily S&P 500 Bull 3X Shares 37.29 5200 P 37.4 1100 P 37.19 25200 P 36.88 37.27 36.38   37.19 0 1.2755553e+07 1590192000034 1590192000003 37.19 p PACIFIC true true 0.0295 2 76.33 16.51 0 0 0.56 1.5 2019-12-23 00:00:00.000 Closed 37.19 252 0 1590192000003 0 0 0 0 true}