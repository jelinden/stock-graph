#Stock graph drawer

Get stock graphs, single graph as an image: http://localhost:8080/stock/KMI

or show multiple images on a html page:
http://localhost:8080/stock/KMI

Uses https://github.com/doneland/yquotes to get the stock price history and
https://github.com/wcharczuk/go-chart to draw the graphs.

## Build

go get github.com/jelinden/stock-graph

cd $GOPATH/src/github.com/jelinden/stock-graph

go build && ./stock-graph