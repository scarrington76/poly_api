package model

// GroupdDailyStocksALl is the response body for the Grouped Daily
// call to the Polygon API. The response gets the daily open, low
// and close (OHLC) for the entire stocks/equities markets
type GroupedDailyStocksAll struct {

	//Whether or not this response was adjusted for splits.
	Adjusted string `sql:"adjusted"`

	//The number of aggregates (minute or day) used to generate the response.
	QueryCount int64 `sql:"queryCount"`

	//A request id assigned by the server.
	Request_id string `sql:"request_id"`

	//The total number of results for this request.
	ResultsCount int64 `sql:"resultsCount"`

	//The status of this request's response.
	Status string `sql:"status"`

	//The array of results
	Results []GroupedDailyResults `sql:"sql:foreign:day_id"`
}

type GroupedDailyResults struct {

	//The exchange symbol that this item is traded under.
	Ticker string `sql:"ticker"`

	//The close price for the symbol in the given time period.
	C int64 `sql:"close_price"`

	//The highest price for the symbol in the given time period.
	H int64 `sql:"high_price"`

	//The lowest price for the symbol in the given time period.
	L int64 `sql:"low_price"`

	//The number of transactions in the aggregate window.
	N int64 `sql:"vol_transactions"`

	//The open price for the symbol in the given time period.
	O int64 `sql:"open_price"`

	//The Unix Msec timestamp for the start of the aggregate window.
	T int64 `sql:"start_date"`

	//The trading volume of the symbol in the given time period.
	V int64 `sql:"volume"`

	//The volume weighted average price.
	Vw int64 `sql:"aveage_price"`
}
