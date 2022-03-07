package stocks

import (
	"api/core/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func MarketStatus() {

	url := "https://api.polygon.io/v1/marketstatus/now?apiKey=" + common.Key

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
