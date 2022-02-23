package stocks

import (
	"api/core/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func TickerDetails(date string, ticker string) {

	url := "https://api.polygon.io/v3/reference/tickers/" + ticker + "?date=" + date + "&apiKey=" + common.Key

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
