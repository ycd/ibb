package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// TicketPrices Contains the price information of the
// public transportation tickets used in Istanbul.
type TicketPrices struct {
	ticketPriceRecords `json:"result"`
}

type ticketPriceRecords struct {
	Records []*struct {
		Tip   string `json:"Bilet Tipi"`
		Ad    string `json:"Bilet Adi"`
		Fiyat int    `json:"Bilet Fiyati"`
	} `json:"records"`
}

// IETTTicketPrices Contains the price information of the
// public transportation tickets used in Istanbul.
//
// Available years: 2019, 2020
func (c *Client) IETTTicketPrices(ctx context.Context, year int) (*TicketPrices, error) {
	var tp TicketPrices
	var url string

	switch year {
	case 2019:
		url = resources.BaseURL + resources.IETTTicketPrices2019
	case 2020:
		url = resources.BaseURL + resources.IETTTicketPrices2020
	default:
		url = resources.BaseURL + resources.IETTTicketPrices2020
	}

	resp, err := c.get(ctx, url)
	if err != nil {
		return &TicketPrices{}, err
	}

	json.Unmarshal(resp, &tp)

	return &tp, nil
}
