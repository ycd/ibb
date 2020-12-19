package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// WasteAmounts contains information on the amount of waste by district, year and type of waste
// (solid waste collection activities, Medical waste collection activities,
// mechanical sweeping activities to prevent pollution in the main arteries and squares)
// https://data.ibb.gov.tr/en/dataset/ilce-yil-ve-atik-turu-bazinda-atik-miktari
func (c *Client) WasteAmounts(ctx context.Context) (*WasteAmount, error) {
	var wa WasteAmount

	url := resources.BaseURL + resources.WasteAmounts

	resp, err := c.get(ctx, url)
	if err != nil {
		return &WasteAmount{}, err
	}

	json.Unmarshal(resp, &wa)

	return &wa, nil
}
