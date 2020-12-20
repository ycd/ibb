package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// IGDASSubscribers contains IGDAS (gas company) subscription numbers
// in the 39 districts of Istanbul.
// https://data.ibb.gov.tr/en/dataset/ilcelere-gore-abone-sayilari
type IGDASSubscribers struct {
	igdasSubscriberRecords `json:"result"`
}

type igdasSubscriberRecords struct {
	Records []*struct {
		Ilce        string `json:"ILCE"`
		AboneSayisi int    `json:"ABONE SAYI"`
	} `json:"records"`
}

// IGDASSubscribers contains IGDAS (gas company) subscription numbers
// in the 39 districts of Istanbul.
// https://data.ibb.gov.tr/en/dataset/ilcelere-gore-abone-sayilari
func (c *Client) IGDASSubscribers(ctx context.Context) (*IGDASSubscribers, error) {
	var is IGDASSubscribers

	url := resources.BaseURL + resources.IGDASSubscribers

	resp, err := c.get(ctx, url)
	if err != nil {
		return &IGDASSubscribers{}, err
	}

	json.Unmarshal(resp, &is)

	return &is, nil
}
