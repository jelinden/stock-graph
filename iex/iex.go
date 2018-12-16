package iex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Graph5Y(symbol string) ([]float64, []time.Time, error) {
	res, err := http.Get("https://api.iextrading.com/1.0/stock/" + symbol + "/chart/5y")
	if err != nil {
		log.Println(err)
		return nil, nil, fmt.Errorf("getting 5y graph failed")
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, nil, fmt.Errorf("getting 5y graph failed")
	}
	prices := []closePrice{}
	err = json.Unmarshal(body, &prices)
	if err != nil {
		log.Println("err", err)
	}
	var closePrices, dates = []float64{}, []time.Time{}
	for _, price := range prices {
		closePrices = append(closePrices, price.Close)
		dates = append(dates, price.Date.TimeUTC())
	}
	return closePrices, dates, nil
}

type closePrice struct {
	Date          JSONTime `json:"date"`
	Close         float64  `json:"close"`
	ChangePercent float64  `json:"changePercent"`
}

type JSONTime struct {
	time.Time
}

func (t *JSONTime) TimeUTC() time.Time {
	return t.UTC()
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse("2006-01-02", strings.Trim(string(data), `"`))
	if err != nil {
		log.Println(err)
		return err
	}
	t.Time = parsedTime
	return nil
}
