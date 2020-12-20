package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// RailSystemsTimeline contains data on the number of voyages of the
// rail system lines in Istanbul. The data are in daily, monthly and yearly formats.
// https://data.ibb.gov.tr/en/dataset/rayli-sistemler-gunluk-aylik-yillik-hat-bazli-sefer-sayilari
type RailSystemsTimeline struct {
	railSystemsRecords `json:"result"`
}

type railSystemsRecords struct {
	Records []*struct {
		Yil     int    `json:"Yil"`
		Periyot string `json:"Periyot"`
		Hatlar
	} `json:"records"`
}

// Hatlar contains information about the railway routes
type Hatlar struct {
	M1Hatti  int `json:"M1 Hatti"`
	M2Hatti  int `json:"M2 Hatti"`
	M3Hatti  int `json:"M3 Hatti"`
	M4Hatti  int `json:"M4 Hatti"`
	M5Hatti  int `json:"M5 Hatti"`
	M6Hatti  int `json:"M6 Hatti"`
	T1Hatti  int `json:"T1 Hatti"`
	T3Hatti  int `json:"T3 Hatti"`
	T4Hatti  int `json:"T4 Hatti"`
	F1Hatti  int `json:"F1 Hatti"`
	TF1Hatti int `json:"TF1 Hatti"`
	TF2Hatti int `json:"TF2 Hatti"`
}

// RailSystemsTimeline contains data on the number of voyages of the
// rail system lines in Istanbul. The data are in daily, monthly and yearly formats.
// https://data.ibb.gov.tr/en/dataset/rayli-sistemler-gunluk-aylik-yillik-hat-bazli-sefer-sayilari
func (c *Client) RailSystemsTimeline(ctx context.Context) (*RailSystemsTimeline, error) {
	var rs RailSystemsTimeline

	url := resources.BaseURL + resources.RailSystemsTimeline

	resp, err := c.get(ctx, url)
	if err != nil {
		return &RailSystemsTimeline{}, err
	}

	json.Unmarshal(resp, &rs)

	return &rs, nil
}

// RailSystemsDailyMaximumJourneys contains the maximum number of daily trips using rail system lines in Istanbul.
type RailSystemsDailyMaximumJourneys struct {
	railSytemsDMJourneyRecords `json:"result"`
}

type railSytemsDMJourneyRecords struct {
	Records []*struct {
		HatAdi string `json:"Hat Adi"`
		yearsWithDot
		IsletmeTuru string `json:"Isletme Turu"`
	} `json:"records"`
}

// DailyMaximumJourneyByRailSystem contains the maximum number of daily trips using rail system lines in Istanbul.
// MARMARAY maximum line-based daily trips are included.
// https://data.ibb.gov.tr/en/dataset/rayli-sistem-gunluk-maksimum-yolculuk-sayilari
func (c *Client) DailyMaximumJourneyByRailSystem(ctx context.Context) (*RailSystemsDailyMaximumJourneys, error) {
	var rs RailSystemsDailyMaximumJourneys

	url := resources.BaseURL + resources.DailyMaximumJourneyByRailSystem

	resp, err := c.get(ctx, url)
	if err != nil {
		return &RailSystemsDailyMaximumJourneys{}, err
	}

	json.Unmarshal(resp, &rs)

	return &rs, nil
}
