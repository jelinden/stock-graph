package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"strconv"

	"github.com/doneland/yquotes"
	"github.com/julienschmidt/httprouter"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func main() {
	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/stock/:symbol", drawChart)
	router.GET("/graph/:width/:height/:symbol", drawChart)
	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := template.ParseFiles("./tmpl/index.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "index", nil)
}

func drawChart(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	yv, xv := getStock(ps.ByName("symbol"))
	width, height := 460, 240
	if ps.ByName("width") != "" && ps.ByName("height") != "" {
		width, _ = strconv.Atoi(ps.ByName("width"))
		height, _ = strconv.Atoi(ps.ByName("height"))
	}
	priceSeries := chart.TimeSeries{
		Name: "SPY",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xv,
		YValues: yv,
	}

	smaSeries := chart.SMASeries{
		Name: "SPY - SMA",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{2.0, 2.0},
		},
		InnerSeries: priceSeries,
	}

	bbSeries := &chart.BollingerBandsSeries{
		Name: "SPY - Bol. Bands",
		Style: chart.Style{
			Show:        true,
			StrokeColor: drawing.ColorFromHex("e0e0e0"),
			FillColor:   drawing.ColorFromHex("e0e0e0").WithAlpha(64),
		},
		InnerSeries: priceSeries,
	}

	graph := chart.Chart{
		Width:  width,
		Height: height,
		XAxis: chart.XAxis{
			Style:        chart.Style{Show: true},
			TickPosition: chart.TickPositionBetweenTicks,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{Show: true},
			Range: &chart.ContinuousRange{
				Max: maxIntSlice(yv) + 3,
				Min: minIntSlice(yv),
			},
		},
		Series: []chart.Series{
			bbSeries,
			priceSeries,
			smaSeries,
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func getStock(symbol string) ([]float64, []time.Time) {
	prices, err := yquotes.HistoryForYears(symbol, 2, yquotes.Daily)
	if err != nil {
		log.Println(err)
	}
	var closePrices, dates = []float64{}, []time.Time{}
	for _, price := range prices {
		closePrices = append(closePrices, price.AdjClose)
		dates = append(dates, price.Date)
	}
	return closePrices, dates
}

func minIntSlice(v []float64) float64 {
	var m float64
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
		}
	}
	if m > 3 {
		m -= 3
	}
	return m
}

func maxIntSlice(v []float64) float64 {
	var m float64
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
		}
	}
	return m
}
