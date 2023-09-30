# DEPRECATED

IEX doesn't have a free plan anymore: https://iexcloud.io/pricing

---

# Stock graph drawer

## Requirements

- Go 1.11

## Usage

Get stock graphs, single graph as an image: http://localhost:8080/stock/KMI

or show multiple images on a html page:
http://localhost:8080/

and get it with a size of your liking
http://localhost:8080/graph/800/400/JNJ

Uses https://github.com/doneland/yquotes to get the stock price history and
https://github.com/wcharczuk/go-chart to draw the graphs.

![screenshot](https://raw.githubusercontent.com/jelinden/stock-graph/master/screenshot.png)

## Build

`go get github.com/jelinden/stock-graph`

`cd $GOPATH/src/github.com/jelinden/stock-graph`

`go build -mod=vendor && ./stock-graph`
