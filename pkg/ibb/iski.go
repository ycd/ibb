package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// DamOccupancy contains information on the daily and annual
// changes of the occupancy rates of dams in Istanbul.
// https://data.ibb.gov.tr/en/dataset/istanbul-baraj-doluluk-oranlari-verisi
type DamOccupancy struct {
	damOccupancyRecords `json:"result"`
}

type damOccupancyRecords struct {
	Records []*struct {
		Date          string  `json:"DATE"`
		Rate          float32 `json:"GENERAL_DAM_OCCUPANCY_RATE"`
		ReservedWater int     `json:"GENERAL_DAM_RESERVED_WATER"`
	} `json:"records"`
}

// DamOccupancyRates contains information on the daily and annual
// changes of the occupancy rates of dams in Istanbul.
// https://data.ibb.gov.tr/en/dataset/istanbul-baraj-doluluk-oranlari-verisi
func (c *Client) DamOccupancyRates(ctx context.Context) (*DamOccupancy, error) {
	var do DamOccupancy

	url := resources.BaseURL + resources.DamOccupancyRates + resources.Limit100000

	resp, err := c.get(ctx, url)
	if err != nil {
		return &DamOccupancy{}, err
	}

	json.Unmarshal(resp, &do)

	return &do, nil
}
