package cryptopia

import (
	"encoding/json"
	"fmt"
)

type Ticks []*Tick

type Tick struct {
	Label 			   string  `json:"Label"`
	TradePairId        int     `json:"TradePairId"`
	AskPrice 		   float64 `json:"AskPrice"`
	BidPrice   		   float64 `json:"BidPrice"`
	Low    			   float64 `json:"Low"`
	High          	   float64 `json:"High"`
	Volume             float64 `json:"Volume"`
	LastPrice          float64 `json:"bidPrice"`
	BuyVolume          float64 `json:"BuyVolume"`
	SellVolume         float64 `json:"SellVolume"`
	Change             float64 `json:"Change"`
	Open          	   float64 `json:"Open"`
	Close          	   float64 `json:"Close"`
	BaseVolume         float64 `json:"BaseVolume"`
	BuyBaseVolume      float64 `json:"BuyBaseVolume"`
	SellBaseVolume     float64 `json:"SellBaseVolume"`
}

// cryptopia API implementation of Ticker endpoint.
//
// Endpoint:  /GetMarkets
// Method: GET
//
// Example: https://www.cryptopia.co.nz/api/GetMarkets
//
// Sample Response:
//
/*
[
    {
      "TradePairId": 1261,
      "Label": "$$$/BTC",
      "AskPrice": 8.3e-7,
      "BidPrice": 7.3e-7,
      "Low": 6.8e-7,
      "High": 9.9e-7,
      "Volume": 1604248.36500224,
      "LastPrice": 7.9e-7,
      "BuyVolume": 62311617.668848,
      "SellVolume": 6287620.29130799,
      "Change": -14.13,
      "Open": 9.2e-7,
      "Close": 7.9e-7,
      "BaseVolume": 1.34830227,
      "BuyBaseVolume": 1.72482846,
      "SellBaseVolume": 173448559.91465986
    },
  ]
*/

func (client *Client) GetTickers() (Ticks, error) {

	resp, err := client.do("GetMarkets", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Ticks, 0)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}

func (client *Client) GetTicker(id string) (*Tick, error) {

	resp, err := client.do("GetMarkets"+"/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Ticks, 1)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res[0], nil
}
