package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	Rank               string `json:"rank"`
	Price_usd          string `json:"price_usd"`
	Price_btc          string `json:"price_btc"`
	Day_volume_usd     string `json:"24h_volume_usd"`
	Market_cap_usd     string `json:"market_cap_usd"`
	Available_supply   string `json:"available_supply"`
	Total_supply       string `json:"total_supply"`
	Max_supply         string `json:"max_supply"`
	Percent_change_1h  string `json:"percent_change_1h"`
	Percent_change_24h string `json:"percent_change_24h"`
	Percent_change_7d  string `json:"percent_change_7d"`
	Last_updated       string `json:"last_updated"`
}

func main() {
	url := "https://api.coinmarketcap.com/v1/ticker?limit=10"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var s []response
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println(separator())
	fmt.Println(header())
	fmt.Println(separator())
	for _, c := range s {
		fmt.Println(line(c))
		fmt.Println(separator())
	}
	fmt.Println()

}

func header() string {
	return fmt.Sprintf("| %-6s | %-12s | %-4s | %-11s | %-6s | %-6s | %-6s |", "Symbol", "Name", "Rank", "Price (USD)", "% 1h", "% 24h", "% 7d")
}

func separator() string {
	return fmt.Sprintf("|--------|--------------|------|-------------|--------|--------|--------|")

}

func line(c response) string {
	one := colorize(c.Percent_change_1h)
	twofour := colorize(c.Percent_change_24h)
	seven := colorize(c.Percent_change_7d)
	return fmt.Sprintf("| %-6s | %-12s | %4s | %11s | %s | %s | %s |", c.Symbol, c.Name, c.Rank, c.Price_usd, one, twofour, seven)
}

func colorize(v string) string {
	red := 31
	green := 32

	var color int
	if v[0] == '-' {
		color = red
	} else {
		color = green
	}

	return fmt.Sprintf("\x1b[%dm%6s\x1b[0m", color, v)
}
