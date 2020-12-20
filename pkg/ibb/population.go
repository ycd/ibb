package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// MunicipalityPopulation2019 contains the data about the population
// based on municipality's on Istanbul in 2019
// https://data.ibb.gov.tr/en/dataset/belediye-nufuslari-veri-seti
type MunicipalityPopulation2019 struct {
	municipalityPopulation2019Records `json:"result"`
}

type municipalityPopulation2019Records struct {
	Records []*struct {
		Belediye string `json:"Belediyeler"`
		Nufus    int    `json:"2019 yili nufuslari"`
	} `json:"records"`
}

// MunicipalityPopulation2019 contains the data about the population
// based on municipality's on Istanbul in 2019
// https://data.ibb.gov.tr/en/dataset/belediye-nufuslari-veri-seti
func (c *Client) MunicipalityPopulation2019(ctx context.Context) (*MunicipalityPopulation2019, error) {
	var mp MunicipalityPopulation2019

	url := resources.BaseURL + resources.MunicipalityPopulation2019

	resp, err := c.get(ctx, url)
	if err != nil {
		return &MunicipalityPopulation2019{}, err
	}

	json.Unmarshal(resp, &mp)

	return &mp, nil
}
